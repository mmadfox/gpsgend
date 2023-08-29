package grpc_test

import (
	"errors"
	"log/slog"
	"net"
	"strconv"
	"time"

	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/broker"
	transportgrpc "github.com/mmadfox/gpsgend/internal/transport/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type conns struct {
	listener      net.Listener
	grpcConn      *grpc.ClientConn
	trackerServer *transportgrpc.TrackerServer
	broker        *broker.Broker
}

func freePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}
	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port, nil
}

func setup() (*conns, func()) {
	p, err := freePort()
	if err != nil {
		panic(err)
	}

	port := strconv.Itoa(p)
	addr := "0.0.0.0:" + port

	// server side
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	con := conns{
		broker: broker.New(nil),
	}
	con.listener = lis
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)

	con.trackerServer = transportgrpc.NewTrackServer(con.broker, slog.Default())
	gpsgendproto.RegisterTrackerServiceServer(grpcServer, con.trackerServer)

	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			if !errors.Is(err, net.ErrClosed) {
				panic(err)
			}
		}
	}()

	time.Sleep(300 * time.Millisecond)

	// client side
	conn, err := grpc.Dial(addr, []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		if con.listener != nil {
			_ = con.listener.Close()
		}
		panic(err)
	}
	con.grpcConn = conn
	go func() {
		con.broker.Run()
	}()

	return &con, func() {
		if con.grpcConn != nil {
			con.grpcConn.Close()
		}
		if con.listener != nil {
			con.listener.Close()
		}
		con.broker.Close()
	}
}
