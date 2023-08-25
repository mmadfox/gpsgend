package grpc_test

import (
	context "context"
	"errors"
	"io"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/google/uuid"
	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	"github.com/mmadfox/gpsgend/internal/types"
	"github.com/stretchr/testify/require"
)

func TestTrackerServer_SubscribeOneClient(t *testing.T) {
	tc, clean := setup()
	t.Cleanup(clean)

	ctx := context.Background()
	cli := gpsgendproto.NewTrackerServiceClient(tc.grpcConn)
	stream, err := cli.Subscribe(ctx, &gpsgendproto.SubscribeRequest{ClientId: uuid.NewString()})
	require.NoError(t, err)
	require.NotNil(t, stream)
	done := make(chan struct{})
	wantEvents := 5

	go func() {
		got := 0
		for {
			if got >= wantEvents {
				break
			}
			resp, err := stream.Recv()
			require.NoError(t, err)
			require.NotEmpty(t, resp.Event)
			got++
		}
		close(done)
		require.NoError(t, stream.CloseSend())
	}()

	for i := 0; i < wantEvents; i++ {
		tc.broker.PublishTrackerRemoved(ctx, types.NewID())
		time.Sleep(3 * time.Millisecond)
	}

	<-done
}

func TestTrackerServer_SubscribeMultiClient(t *testing.T) {
	tc, clean := setup()
	t.Cleanup(clean)

	ctx := context.Background()
	wantEvents := 5
	wantClients := 5
	var wg sync.WaitGroup
	var wgdone sync.WaitGroup

	var got uint32

	for i := 0; i < wantClients; i++ {
		wg.Add(1)
		wgdone.Add(1)
		go func() {
			defer wg.Done()

			cli := gpsgendproto.NewTrackerServiceClient(tc.grpcConn)
			stream, err := cli.Subscribe(ctx, &gpsgendproto.SubscribeRequest{ClientId: uuid.NewString()})
			require.NoError(t, err)
			require.NotNil(t, stream)

			wgdone.Done()

			cnt := 0
			for {
				if cnt >= wantEvents {
					break
				}
				resp, err := stream.Recv()
				require.NoError(t, err)
				require.NotEmpty(t, resp.Event)
				cnt++
			}

			atomic.AddUint32(&got, uint32(cnt))
		}()
	}

	wgdone.Wait()

	for i := 0; i <= wantEvents; i++ {
		tc.broker.PublishTrackerRemoved(ctx, types.NewID())
		time.Sleep(5 * time.Millisecond)
	}

	wg.Wait()

	require.Equal(t, uint32(wantClients*wantEvents), atomic.LoadUint32(&got))
}

func TestTrackerServer_UnsubscribeMultiClient(t *testing.T) {
	tc, clean := setup()
	t.Cleanup(clean)

	ctx := context.Background()
	wantClients := 5
	var wg sync.WaitGroup
	var wgdone sync.WaitGroup

	type client struct {
		id  string
		cli gpsgendproto.TrackerServiceClient
	}

	ids := make(chan client, wantClients)

	for i := 0; i < wantClients; i++ {
		wg.Add(1)
		wgdone.Add(1)

		go func() {
			defer wg.Done()

			cli := gpsgendproto.NewTrackerServiceClient(tc.grpcConn)
			id := uuid.NewString()
			stream, err := cli.Subscribe(ctx, &gpsgendproto.SubscribeRequest{ClientId: id})
			require.NoError(t, err)
			require.NotNil(t, stream)

			ids <- client{id: id, cli: cli}
			wgdone.Done()

			for {
				_, err := stream.Recv()
				if errors.Is(err, io.EOF) {
					break
				}
			}
		}()
	}

	wgdone.Wait()
	time.Sleep(10 * time.Millisecond)

	for i := 0; i < wantClients; i++ {
		cli := <-ids
		resp, err := cli.cli.Unsubscribe(ctx, &gpsgendproto.UnsubscribeRequest{ClientId: cli.id})
		require.NoError(t, err)
		require.Nil(t, resp.Error)
	}

	wg.Wait()
}
