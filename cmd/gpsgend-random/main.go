package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"os"
	"text/tabwriter"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/random"
	transportgrpc "github.com/mmadfox/gpsgend/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var routeLevels = []int{
	gpsgen.RouteLevelXS,
	gpsgen.RouteLevelM,
	gpsgen.RouteLevelL,
	gpsgen.RouteLevelXL,
	gpsgen.RouteLevelXXL,
	gpsgen.RouteLevelS,
}

type arguments struct {
	addr             string
	numTrackers      int
	routesPerTracker int
	countryCode      string
	fillMode         bool
}

func (a *arguments) prepare() {
	if a.numTrackers <= 0 {
		a.numTrackers = 10
	}
	if a.routesPerTracker <= 0 {
		a.routesPerTracker = 3
	}
	if a.routesPerTracker > 10 {
		a.routesPerTracker = 10
	}
	if !a.fillMode {
		if len(a.countryCode) == 0 {
			a.countryCode = "us"
		}
	}
}

func main() {
	args := parseArgs()

	var countryName string
	var err error
	if !args.fillMode {
		countryName, err = random.CountryName(args.countryCode)
		if err != nil {
			fmt.Println("[ERROR]", err)
			os.Exit(1)
		}
	}

	fmt.Println("connecting to gpsgend server...")

	conn, err := grpc.Dial(args.addr, []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	if args.fillMode {
		fmt.Printf("starting for all countries...\n")
	} else {
		fmt.Printf("starting for country %s...\n", countryName)
	}

	cli := transportgrpc.New(conn)
	ctx := context.Background()
	defer conn.Close()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
	var total int

	if args.fillMode {
		total = genForEachCountry(ctx, args, cli, rnd)
	} else {
		total = genForCountry(ctx, args, cli, rnd)
	}

	fmt.Printf("successfully added: %d trackers\n", total)
}

func genForCountry(ctx context.Context, args *arguments, cli *transportgrpc.Client, rnd *rand.Rand) (total int) {
	for i := 0; i < args.numTrackers; i++ {
		tracker, err := cli.AddTracker(ctx, newOpts(rnd))
		if err != nil {
			panic(err)
		}

		routes := make([]*gpsgen.Route, 0, args.routesPerTracker)
		for r := 0; r < args.routesPerTracker; r++ {
			if args.fillMode {

			} else {
				latLon, _ := random.LatLonByCountry(args.countryCode)
				route := gpsgen.RandomRoute(latLon.Lon, latLon.Lat, randomNumTracks(rnd), randomRouteLevel(rnd))
				routes = append(routes, route)
			}
		}

		if err := cli.AddRoutes(ctx, tracker.ID, routes...); err != nil {
			panic(err)
		}

		if _, err := cli.AddSensor(
			ctx,
			tracker.ID,
			"sensor-"+random.String(5),
			1, between(rnd, 2, 25), 4, 0,
		); err != nil {
			panic(err)
		}

		if err := cli.StartTracker(ctx, tracker.ID); err != nil {
			panic(err)
		}

		fmt.Printf("added tracker %s for country %s\n", tracker.ID, args.countryCode)
		total++
	}
	return
}

func genForEachCountry(ctx context.Context, args *arguments, cli *transportgrpc.Client, rnd *rand.Rand) (total int) {
	random.EachCountry(func(code, countryName string) {
		if code == "GL" || code == "NZ" || code == "NL" {
			return
		}
		for i := 0; i < args.numTrackers; i++ {
			tracker, err := cli.AddTracker(ctx, newOpts(rnd))
			if err != nil {
				panic(err)
			}

			routes := make([]*gpsgen.Route, 0, args.routesPerTracker)
			for r := 0; r < args.routesPerTracker; r++ {
				latLon, _ := random.LatLonByCountry(code)
				route := gpsgen.RandomRoute(latLon.Lon, latLon.Lat, randomNumTracks(rnd), randomRouteLevel(rnd))
				routes = append(routes, route)
			}

			if err := cli.AddRoutes(ctx, tracker.ID, routes...); err != nil {
				panic(err)
			}

			if _, err := cli.AddSensor(
				ctx,
				tracker.ID,
				"sensor-"+random.String(5),
				1, between(rnd, 2, 25), 4, 0,
			); err != nil {
				panic(err)
			}

			if err := cli.StartTracker(ctx, tracker.ID); err != nil {
				panic(err)
			}

			fmt.Printf("added tracker %s for country %s\n", tracker.ID, countryName)
			total++
		}
	})
	return total
}

func parseArgs() *arguments {
	fs := flag.NewFlagSet("gpsgend-random", flag.ExitOnError)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	args := new(arguments)
	fs.IntVar(&args.numTrackers, "trackers", 10, "Number of trackers")
	fs.IntVar(&args.routesPerTracker, "maxRoutes", 3, "Max number of routes for each track, between 1-10")
	fs.StringVar(&args.countryCode, "country", "us", "Country based on its two-letter country code, [us, ru, zw, ...]")
	fs.BoolVar(&args.fillMode, "fillmode", false, "Trackers are generated for each country")
	fs.StringVar(&args.addr, "addr", "0.0.0.0:15015", "Server address")
	fs.Parse(os.Args[1:])
	args.prepare()
	return args
}

func newOpts(r *rand.Rand) *transportgrpc.AddTrackerOptions {
	opts := transportgrpc.NewAddTrackerOptions()
	opts.CustomID = uuid.NewString()
	opts.Descr = "Descr-" + random.String(16)
	opts.Speed.Min = 1
	opts.Speed.Max = between(r, 2, 11)
	opts.Elevation.Min = 1
	opts.Elevation.Max = between(r, 5, 500)
	opts.Battery.Min = 1
	opts.Battery.Max = between(r, 5, 100)
	return opts
}

func between(r *rand.Rand, min, max float64) float64 {
	return min + r.Float64()*(max-min)
}

func randomRouteLevel(r *rand.Rand) int {
	min := 0
	max := len(routeLevels)
	return routeLevels[rand.Intn(max-min)+min]
}

func randomNumTracks(r *rand.Rand) int {
	min := 1
	max := 5
	return rand.Intn(max-min) + min
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
