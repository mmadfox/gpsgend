package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"text/tabwriter"

	"log/slog"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/mmadfox/go-gpsgen"
	config "github.com/mmadfox/gpsgend/configs"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/broker"
	"github.com/mmadfox/gpsgend/internal/generator"
	storagemongo "github.com/mmadfox/gpsgend/internal/storage/mongodb"
	transportgrpc "github.com/mmadfox/gpsgend/internal/transport/grpc"
	transporthttp "github.com/mmadfox/gpsgend/internal/transport/http"
	transportws "github.com/mmadfox/gpsgend/internal/transport/websocket"
	"github.com/oklog/run"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"moul.io/banner"
)

func main() {
	showBanner()

	fs := flag.NewFlagSet("gpsgend", flag.ExitOnError)
	var (
		confFilename = fs.String("config", "", "Path to the configuration file")
	)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	fs.Parse(os.Args[1:])

	conf, err := config.FromFile(*confFilename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	ctx := context.Background()
	logger := setupLogger(conf)

	// storage
	mongoConn, err := setupMongodb(ctx, conf.Storage.Mongodb.URI)
	if err != nil {
		logger.Error("Failed to connect to mongodb", "err", err)
		os.Exit(1)
	}

	db := mongoConn.Database(conf.Storage.Mongodb.DatabaseName)
	col := db.Collection(conf.Storage.Mongodb.CollectionName)
	monogoStorage := storagemongo.New(col)
	mongoBootstraper := storagemongo.NewBootstraper(col)
	mongoQuery := storagemongo.NewQuery(col)
	if err := storagemongo.EnsureIndexes(ctx, col); err != nil {
		logger.Error("Failed to create indexes in mongo", "err", err)
		os.Exit(1)
	}

	// events broker
	eventBroker := broker.New(conf.EventBrokerOpts())
	go func() {
		eventBroker.Run()
	}()

	// processes
	processes := gpsgen.New(conf.GeneratorOpts())
	go func() {
		processes.Run()
	}()

	processes.OnError(func(err error) {
		logger.Error("GPS data generation error", "err", err)
	})
	processes.OnPacket(func(b []byte) {
		eventBroker.PublishTrackerChanged(ctx, b)
	})

	// generator service
	gen := generator.New(
		monogoStorage,
		processes,
		mongoBootstraper,
		mongoQuery,
		eventBroker,
	)

	// bootstrap processes
	if err := gen.Run(ctx); err != nil {
		logger.Error("Failed to bootstrap generator service", "err", err)
		os.Exit(1)
	}

	var g run.Group
	{
		grpcAddr := conf.Transport.GRPC.Listen
		grpcListener, err := net.Listen("tcp", grpcAddr)
		if err != nil {
			logger.Error("Failed to listen to address", "addr", grpcAddr, "err", err)
			os.Exit(1)
		}
		g.Add(func() error {
			logger.Info("gRPC server is running", "addr", grpcAddr)

			opts := []logging.Option{
				logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
			}
			baseServer := grpc.NewServer([]grpc.ServerOption{
				grpc.ChainUnaryInterceptor(
					logging.UnaryServerInterceptor(
						transportgrpc.InterceptorLogger(logger), opts...),
				),
			}...)
			trackerServer := transportgrpc.NewTrackServer(eventBroker)
			generatorServer := transportgrpc.NewGeneratorServer(gen)
			gpsgendproto.RegisterTrackerServiceServer(baseServer, trackerServer)
			gpsgendproto.RegisterGeneratorServiceServer(baseServer, generatorServer)
			return baseServer.Serve(grpcListener)
		}, func(error) {
			logger.Info("gRPC server stopped")
			grpcListener.Close()
		})
	}
	{
		httpAddr := conf.Transport.HTTP.Listen
		httpServer := transporthttp.New(httpAddr, gen, logger)
		g.Add(func() error {
			logger.Info("HTTP server is running", "addr", httpAddr)
			return httpServer.Listen()
		}, func(error) {
			logger.Info("HTTP server stopped")
			httpServer.Close()
		})
	}
	{
		wsAddr := conf.Transport.Websocket.Listen
		wsServer := transportws.New(wsAddr, eventBroker, logger)
		g.Add(func() error {
			logger.Info("Websocket server is running", "addr", wsAddr)
			return wsServer.Listen()
		}, func(error) {
			logger.Info("Websocket server stopped")
			wsServer.Close()
		})
	}
	{
		cancelInterrupt := make(chan struct{})
		g.Add(func() error {
			c := make(chan os.Signal, 1)
			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
			logger.Info("Geodata generator is running")
			select {
			case <-c:
				return nil
			case <-cancelInterrupt:
				return nil
			}
		}, func(error) {
			gen.Close(ctx)
			logger.Info("Bootstrapper stopped")

			processes.Close()
			logger.Info("GPS data generator stopped")

			eventBroker.Close()
			logger.Info("Event broker stopped")

			mongoConn.Disconnect(ctx)
			logger.Info("MongoDB stopped")

			close(cancelInterrupt)
		})
	}

	if err := g.Run(); err != nil {
		logger.Error("Exit", "err", err)
	} else {
		logger.Info("Exit OK")
	}
}

func usageFor(fs *flag.FlagSet, short string) func() {
	return func() {
		fmt.Fprintf(os.Stderr, "USAGE\n")
		fmt.Fprintf(os.Stderr, "  %s\n", short)
		fmt.Fprintf(os.Stderr, "\n")
		fmt.Fprintf(os.Stderr, "FLAGS\n")
		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
		fs.VisitAll(func(f *flag.Flag) {
			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
		})
		w.Flush()
		fmt.Fprintf(os.Stderr, "\n")
	}
}

func setupMongodb(ctx context.Context, uri string) (*mongo.Client, error) {
	opts := options.Client().ApplyURI(uri)
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	return mongo.Connect(ctx, opts)
}

func setupLogger(conf *config.Config) *slog.Logger {
	var handler slog.Handler
	slogLevel := slog.LevelInfo
	format := strings.ToLower(strings.TrimSpace(conf.Logger.Format))
	level := strings.ToLower(strings.TrimSpace(conf.Logger.Format))
	switch level {
	case "info":
		slogLevel = slog.LevelInfo
	case "warn":
		slogLevel = slog.LevelWarn
	case "debug":
		slogLevel = slog.LevelDebug
	case "error":
		slogLevel = slog.LevelError
	}
	opts := slog.HandlerOptions{Level: slogLevel}
	switch format {
	case "json":
		handler = slog.NewJSONHandler(os.Stderr, &opts)
	default:
		handler = slog.NewTextHandler(os.Stderr, &opts)
	}
	logger := slog.New(handler)
	return logger.With("service", conf.Service)
}

func showBanner() {
	fmt.Println("===============================================")
	fmt.Println(banner.Inline("gpsgend") + "\n")
	fmt.Println("===============================================")
}
