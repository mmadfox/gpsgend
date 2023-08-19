package main

// import (
// 	"context"
// 	"flag"
// 	"fmt"
// 	"os"
// 	"os/signal"
// 	"strings"
// 	"syscall"

// 	"text/tabwriter"

// 	"github.com/mmadfox/go-gpsgen"
// 	"github.com/mmadfox/gpsgend/config"
// 	apihttp "github.com/mmadfox/gpsgend/internal/api/http"
// 	"github.com/mmadfox/gpsgend/internal/device"
// 	"github.com/mmadfox/gpsgend/internal/publisher"
// 	"github.com/mmadfox/gpsgend/internal/service"
// 	storagemongo "github.com/mmadfox/gpsgend/internal/storage/mongo"
// 	"github.com/oklog/run"
// 	"go.mongodb.org/mongo-driver/mongo"
// 	"go.mongodb.org/mongo-driver/mongo/options"
// 	"golang.org/x/exp/slog"
// 	"moul.io/banner"
// )

// func main() {
// 	showBanner()

// 	fs := flag.NewFlagSet("gpsgend", flag.ExitOnError)
// 	var (
// 		confFilename = fs.String("config", "", "path to the configuration file")
// 	)
// 	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
// 	fs.Parse(os.Args[1:])

// 	conf, err := config.FromFile(*confFilename)
// 	if err != nil {
// 		fmt.Println(err)
// 		os.Exit(1)
// 	}

// 	ctx := context.Background()
// 	logger := setupLogger(conf)

// 	mongoClient, err := setupMongoClient(ctx, conf)
// 	if err != nil {
// 		logger.Error("setup mongodb client", "err", err)
// 		os.Exit(1)
// 	}
// 	defer mongoClient.Disconnect(ctx)

// 	dataGenerator := gpsgen.New()

// 	deviceStorage := storagemongo.NewStorage(
// 		mongoClient,
// 		conf.Storage.Mongodb.DatabaseName,
// 		conf.Storage.Mongodb.CollectionName,
// 	)
// 	if err := deviceStorage.EnsureIndexes(); err != nil {
// 		logger.Error("ensure mongodb indexes", "err", err)
// 		os.Exit(1)
// 	}

// 	deviceQuery := service.NewDeviceQuery(deviceStorage, logger)
// 	devicePub := publisher.New(logger)

// 	deviceService := service.NewDeviceService(
// 		deviceStorage,
// 		devicePub,
// 		dataGenerator,
// 	)
// 	if err := deviceService.Bootstrap(ctx); err != nil {
// 		logger.Error("bootstrap gpsgend service", "err", err)
// 		os.Exit(1)
// 	}
// 	defer deviceService.Close(ctx)

// 	var deviceUseCase device.UseCase
// 	{
// 		deviceUseCase = service.WithLogging(logger)(deviceService)
// 	}

// 	dataGenerator.Run()

// 	logger.Info("service started")

// 	var g run.Group
// 	{
// 		httpSrv := apihttp.NewServer(deviceUseCase, deviceQuery, devicePub, logger)
// 		g.Add(func() error {
// 			addr := conf.API.HTTP.Addr
// 			logger.Info("HTTP API started", "addr", addr)
// 			return httpSrv.Listen(addr)
// 		}, func(_ error) {
// 			httpSrv.Close()
// 		})
// 	}
// 	{
// 		cancelInterrupt := make(chan struct{})
// 		g.Add(func() error {
// 			c := make(chan os.Signal, 1)
// 			signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
// 			select {
// 			case sig := <-c:
// 				return fmt.Errorf("received signal %s", sig)
// 			case <-cancelInterrupt:
// 				return nil
// 			}
// 		}, func(error) {
// 			close(cancelInterrupt)
// 		})
// 	}

// 	logger.Error("exit", "err", g.Run())
// }

// func usageFor(fs *flag.FlagSet, short string) func() {
// 	return func() {
// 		fmt.Fprintf(os.Stderr, "USAGE\n")
// 		fmt.Fprintf(os.Stderr, "  %s\n", short)
// 		fmt.Fprintf(os.Stderr, "\n")
// 		fmt.Fprintf(os.Stderr, "FLAGS\n")
// 		w := tabwriter.NewWriter(os.Stderr, 0, 2, 2, ' ', 0)
// 		fs.VisitAll(func(f *flag.Flag) {
// 			fmt.Fprintf(w, "\t-%s %s\t%s\n", f.Name, f.DefValue, f.Usage)
// 		})
// 		w.Flush()
// 		fmt.Fprintf(os.Stderr, "\n")
// 	}
// }

// func setupMongoClient(ctx context.Context, conf *config.Config) (*mongo.Client, error) {
// 	opts := options.Client().ApplyURI(conf.Storage.Mongodb.URI)
// 	ctx, cancel := context.WithCancel(ctx)
// 	defer cancel()
// 	return mongo.Connect(ctx, opts)
// }

// func setupLogger(conf *config.Config) *slog.Logger {
// 	var handler slog.Handler
// 	slogLevel := slog.LevelInfo
// 	format := strings.ToLower(strings.TrimSpace(conf.Logger.Format))
// 	level := strings.ToLower(strings.TrimSpace(conf.Logger.Format))
// 	switch level {
// 	case "info":
// 		slogLevel = slog.LevelInfo
// 	case "warn":
// 		slogLevel = slog.LevelWarn
// 	case "debug":
// 		slogLevel = slog.LevelDebug
// 	case "error":
// 		slogLevel = slog.LevelError
// 	}
// 	opts := slog.HandlerOptions{Level: slogLevel}
// 	switch format {
// 	case "json":
// 		handler = slog.NewJSONHandler(os.Stderr, &opts)
// 	default:
// 		handler = slog.NewTextHandler(os.Stderr, &opts)
// 	}
// 	logger := slog.New(handler)
// 	return logger.With("service", conf.Service)
// }

// func showBanner() {
// 	fmt.Println("===============================================")
// 	fmt.Println(banner.Inline("gpsgend") + "\n")
// 	fmt.Println("===============================================")
// }
