package main

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/mmadfox/go-gpsgen"
	tranportgrpc "github.com/mmadfox/gpsgend/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	fmt.Println("BOOTSTRAP")

	time.Sleep(5 * time.Second)

	conn, err := grpc.Dial("gpsgend:15015", []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		fmt.Println("[ERROR]", err)
		os.Exit(1)
	}

	cli := tranportgrpc.New(conn)
	ctx := context.Background()
	defer conn.Close()

	for i := 0; i < 350; i++ {
		opts := tranportgrpc.NewAddTrackerOptions()
		tracker, err := cli.AddTracker(ctx, opts)
		if err != nil {
			panic(err)
		}

		route := gpsgen.RandomRouteForNewYork()
		if err := cli.AddRoutes(ctx, tracker.ID, route); err != nil {
			panic(err)
		}

		if err := cli.StartTracker(ctx, tracker.ID); err != nil {
			panic(err)
		}
	}

	fmt.Println("Done OK!")
}
