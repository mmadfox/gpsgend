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
	"github.com/mmcloughlin/spherand"
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

func main() {
	fs := flag.NewFlagSet("gpsgend-random", flag.ExitOnError)
	fs.Usage = usageFor(fs, os.Args[0]+" [flags]")
	numTrackers := fs.Int("trackers", 10, "Number of trackers")
	routesPerTracker := fs.Int("maxRoutes", 3, "Max number of routes for each track, Between 1-10")
	gpsgendAddr := fs.String("addr", "0.0.0.0:15015", "Server address")
	fs.Parse(os.Args[1:])

	if *numTrackers <= 0 {
		*numTrackers = 10
	}

	if *routesPerTracker <= 0 {
		*routesPerTracker = 3
	}
	if *routesPerTracker > 10 {
		*routesPerTracker = 10
	}

	fmt.Println("connecting to gpsgend server...")

	conn, err := grpc.Dial(*gpsgendAddr, []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	fmt.Println("starting...")

	cli := transportgrpc.New(conn)
	ctx := context.Background()
	defer conn.Close()
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for i := 0; i < *numTrackers; i++ {
		tracker, err := cli.AddTracker(ctx, newOpts(rnd))
		if err != nil {
			panic(err)
		}

		routes := make([]*gpsgen.Route, 0, *routesPerTracker)
		for r := 0; r < *routesPerTracker; r++ {
			lat, lon := spherand.Geographical()
			route := gpsgen.RandomRoute(lon, lat, randomNumTracks(rnd), randomRouteLevel(rnd))
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

		fmt.Println("added tracker", tracker.ID)
	}

	fmt.Printf("successfully added: %d trackers\n", *numTrackers)
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
