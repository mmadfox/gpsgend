package grpc_test

import (
	"context"
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/google/uuid"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	gpsgendclient "github.com/mmadfox/gpsgend/pkg/grpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var listener net.Listener
var grpcConn *grpc.ClientConn
var generatorHandler *generatorService

func setup() {
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
	listener = lis
	grpcServer := grpc.NewServer([]grpc.ServerOption{}...)
	generatorHandler = new(generatorService)
	gpsgendproto.RegisterGeneratorServiceServer(grpcServer, generatorHandler)
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			if !errors.Is(err, net.ErrClosed) {
				panic(err)
			}
		}
	}()

	time.Sleep(time.Second)

	// client side
	conn, err := grpc.Dial(addr, []grpc.DialOption{
		grpc.WithTransportCredentials(
			insecure.NewCredentials(),
		),
		grpc.WithBlock()}...,
	)
	if err != nil {
		if listener != nil {
			_ = listener.Close()
		}
		panic(err)
	}
	grpcConn = conn
}

func teardown() {
	if grpcConn != nil {
		grpcConn.Close()
	}
	if listener != nil {
		listener.Close()
	}
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

func newClient() *gpsgendclient.Client {
	return gpsgendclient.New(grpcConn)
}

var (
	expectedErr     = errors.New("error")
	expectedTracker = &gpsgendproto.Tracker{
		Id:       uuid.NewString(),
		CustomId: uuid.NewString(),
		Status: &gpsgendproto.Status{
			Id:   1,
			Name: "Running",
		},
		Model: "model",
		Color: "color",
		Descr: "descr",
		Offline: &gpsgendproto.Offline{
			Min: 1,
			Max: 10,
		},
		Elevation: &gpsgendproto.Elevation{
			Min:       1,
			Max:       100,
			Amplitude: 8,
			Mode:      0,
		},
		Battery: &gpsgendproto.Battery{
			Min:        1,
			Max:        5,
			ChargeTime: int64(time.Hour),
		},
		NumSensors:  1,
		NumRoutes:   1,
		SkipOffline: true,
		CreatedAt:   time.Now().Unix(),
		UpdatedAt:   time.Now().Unix(),
		RunningAt:   time.Now().Unix(),
		StoppedAt:   time.Now().Unix(),
	}
	expectedTrackers = []*gpsgendproto.Tracker{
		expectedTracker,
		expectedTracker,
		expectedTracker,
		expectedTracker,
	}
)

type generatorService struct {
	gpsgendproto.GeneratorServiceServer

	assert func(req any)

	wantErr     error
	wantErrResp error
	want        *gpsgendproto.Tracker
}

func (s *generatorService) reset() {
	s.want = nil
	s.wantErr = nil
	s.wantErrResp = nil
}

func (s *generatorService) NewTracker(ctx context.Context, req *gpsgendproto.NewTrackerRequest) (*gpsgendproto.NewTrackerResponse, error) {
	if s.wantErr != nil {
		return nil, s.wantErr
	}
	if s.wantErrResp != nil {
		return &gpsgendproto.NewTrackerResponse{
			Error: &gpsgendproto.Error{
				Msg: s.wantErrResp.Error(),
			},
		}, nil
	}
	if s.assert != nil {
		s.assert(req)
	}
	return &gpsgendproto.NewTrackerResponse{
		Tracker: s.want,
	}, nil
}

func (s *generatorService) SearchTrackers(ctx context.Context, req *gpsgendproto.SearchTrackersRequest) (*gpsgendproto.SearchTrackersResponse, error) {
	if s.wantErr != nil {
		return nil, s.wantErr
	}
	if s.wantErrResp != nil {
		return &gpsgendproto.SearchTrackersResponse{
			Error: &gpsgendproto.Error{
				Msg: s.wantErrResp.Error(),
			},
		}, nil
	}
	if s.assert != nil {
		s.assert(req)
	}
	return &gpsgendproto.SearchTrackersResponse{
		Trackers: expectedTrackers,
	}, nil
}
