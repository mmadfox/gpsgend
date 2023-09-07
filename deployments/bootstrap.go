package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/random"
	transportgrpc "github.com/mmadfox/gpsgend/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("BOOTSTRAP")

	time.Sleep(5 * time.Second)

	conn, err := grpc.Dial("0.0.0.0:15015", []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	cli := transportgrpc.New(conn)
	ctx := context.Background()
	defer conn.Close()

	for i := 0; i < 1000; i++ {
		opts := transportgrpc.NewAddTrackerOptions()
		opts.CustomID = uuid.NewString()
		opts.Descr = "Descr-" + random.String(16)
		tracker, err := cli.AddTracker(ctx, opts)
		if err != nil {
			panic(err)
		}

		var route *gpsgen.Route
		if i < 500 {
			route = gpsgen.RandomRouteForNewYork()
		} else {
			route = gpsgen.RandomRouteForMoscow()
		}

		if err := cli.AddRoutes(ctx, tracker.ID, route); err != nil {
			panic(err)
		}

		if _, err := cli.AddSensor(ctx, tracker.ID, "s1", 1, 100, 8, 0); err != nil {
			panic(err)
		}

		if err := cli.StartTracker(ctx, tracker.ID); err != nil {
			panic(err)
		}
	}

	fmt.Println("Done OK!")
}
