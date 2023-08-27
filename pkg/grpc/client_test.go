package grpc_test

import (
	"context"
	"os"
	"reflect"
	"testing"

	gpsgendproto "github.com/mmadfox/gpsgend/gen/proto/gpsgend/v1"
	gpsgendclient "github.com/mmadfox/gpsgend/pkg/grpc"
	"github.com/stretchr/testify/require"
)

func TestMain(m *testing.M) {
	setup()
	ret := m.Run()
	teardown()
	os.Exit(ret)
}

func TestClient_AddTracker(t *testing.T) {
	tracker, err := gpsgendclient.DecodeTracker(expectedTracker)
	require.NoError(t, err)

	tests := []struct {
		name    string
		want    *gpsgendclient.Tracker
		mock    func()
		wantErr error
	}{
		{
			name: "should return error when params are invalid",
			mock: func() {
				generatorHandler.wantErr = expectedErr
			},
			wantErr: expectedErr,
		},
		{
			name: "should return error when generator failure",
			mock: func() {
				generatorHandler.wantErrResp = expectedErr
			},
			wantErr: expectedErr,
		},
		{
			name: "should not return error when all params are valid",
			mock: func() {
				generatorHandler.want = expectedTracker
			},
			want:    tracker,
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newClient()
			if tt.mock != nil {
				tt.mock()
			}
			opts := gpsgendclient.NewAddTrackerOptions()
			got, err := c.AddTracker(context.Background(), opts)
			generatorHandler.reset()
			if tt.wantErr != nil {
				require.EqualValues(t, tt.wantErr, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.NewTracker() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClient_SearchTrackers(t *testing.T) {
	want := make([]*gpsgendclient.Tracker, len(expectedTrackers))
	for i := 0; i < len(expectedTrackers); i++ {
		trk, err := gpsgendclient.DecodeTracker(expectedTrackers[i])
		require.NoError(t, err)
		want[i] = trk
	}
	type args struct {
		f gpsgendclient.Filter
	}
	tests := []struct {
		name    string
		args    args
		mock    func(gpsgendclient.Filter)
		want    gpsgendclient.SearchResult
		wantErr error
	}{
		{
			name: "should return error when params are invalid",
			mock: func(gpsgendclient.Filter) {
				generatorHandler.wantErr = expectedErr
			},
			wantErr: expectedErr,
		},
		{
			name: "should return error when generator failure",
			mock: func(gpsgendclient.Filter) {
				generatorHandler.wantErrResp = expectedErr
			},
			wantErr: expectedErr,
		},
		{
			name: "should not return error when all params are valid",
			args: args{
				f: gpsgendclient.Filter{
					Term:   "term",
					Status: 1,
					Limit:  1,
					Offset: 1,
				},
			},
			mock: func(f gpsgendclient.Filter) {
				generatorHandler.assert = func(r any) {
					req, ok := r.(*gpsgendproto.SearchTrackersRequest)
					require.True(t, ok)
					require.NotNil(t, req.Filter)
				}
				generatorHandler.want = expectedTracker
			},
			want:    gpsgendclient.SearchResult{Trackers: want},
			wantErr: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := newClient()
			if tt.mock != nil {
				tt.mock(tt.args.f)
			}
			got, err := c.SearchTrackers(context.Background(), tt.args.f)
			generatorHandler.reset()
			if tt.wantErr != nil {
				require.EqualValues(t, tt.wantErr, err)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Client.SearchTrackers() = %v, want %v", got, tt.want)
			}
		})
	}
}
