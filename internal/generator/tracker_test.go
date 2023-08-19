package generator_test

import (
	"testing"

	"github.com/mmadfox/go-gpsgen/geo"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/stretchr/testify/require"
)

var (
	track300m1segment, _ = navigator.NewTrack([]geo.LatLonPoint{
		{Lon: 106.49331396675268, Lat: 29.5299004724652},
		{Lon: 106.49523863664103, Lat: 29.532016484207674},
	})
)

func TestTracker_AddRoutes(t *testing.T) {
	type args struct {
		newRoutes func() []*navigator.Route
	}
	tests := []struct {
		name    string
		args    args
		arrange func(*generator.Tracker)
		assert  func(*generator.Tracker, []*navigator.Route)
		wantErr bool
	}{
		{
			name: "should return error when max number of routes exceeded",
			args: args{
				newRoutes: func() []*navigator.Route {
					routes := make([]*navigator.Route, 0)
					for i := 0; i < 3; i++ {
						r1 := navigator.RouteFromTracks(track300m1segment)
						routes = append(routes, r1)
					}
					return routes
				},
			},
			arrange: func(trk *generator.Tracker) {
				routes := make([]*navigator.Route, 0)
				for i := 0; i < generator.MaxRoutesPerTracker-1; i++ {
					r1 := navigator.RouteFromTracks(track300m1segment)
					routes = append(routes, r1)
				}
				_, err := trk.AddRoutes(routes)
				require.NoError(t, err)
			},
			assert: func(trk *generator.Tracker, routes []*navigator.Route) {
				require.Nil(t, routes)
				require.Equal(t, generator.MaxRoutesPerTracker-1, trk.NumRoutes())
			},
			wantErr: true,
		},
		{
			name: "should return error when max number of tracks exceeded",
			args: args{
				newRoutes: func() []*navigator.Route {
					tracks := make([]*navigator.Track, 0)
					for i := 0; i < generator.MaxTracksPerRoute+1; i++ {
						tracks = append(tracks, track300m1segment)
					}
					route := navigator.RouteFromTracks(tracks...)
					return []*navigator.Route{route}
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when max number of segments exceeded",
			args: args{
				newRoutes: func() []*navigator.Route {
					points := make([]geo.LatLonPoint, 0)
					for i := 0; i < generator.MaxSegmentsPerTrack+10; i++ {
						points = append(points, geo.LatLonPoint{Lon: 106.49331396675268, Lat: 29.5299004724652})
					}
					track, err := navigator.NewTrack(points)
					require.NoError(t, err)
					route := navigator.RouteFromTracks(track)
					return []*navigator.Route{route}
				},
			},
			wantErr: true,
		},
		{
			name: "should return error when list of routes is empty",
			args: args{
				newRoutes: func() []*navigator.Route { return nil },
			},
			wantErr: true,
		},
		{
			name: "should add new routes when tracker init",
			args: args{
				newRoutes: func() []*navigator.Route {
					r1 := navigator.RouteFromTracks(track300m1segment)
					return []*navigator.Route{r1, r1, r1}
				},
			},
			assert: func(trk *generator.Tracker, routes []*navigator.Route) {
				require.Equal(t, 1, trk.NumRoutes())
				require.Len(t, routes, 1)
				actualRoutes, err := trk.Routes()
				require.NoError(t, err)
				require.Equal(t, routes, actualRoutes)
			},
			wantErr: false,
		},
		{
			name: "should add new routes when routes already exists",
			args: args{
				newRoutes: func() []*navigator.Route {
					r1 := navigator.RouteFromTracks(track300m1segment)
					return []*navigator.Route{r1}
				},
			},
			arrange: func(trk *generator.Tracker) {
				r1 := navigator.RouteFromTracks(track300m1segment)
				trk.AddRoutes([]*navigator.Route{r1})
			},
			assert: func(trk *generator.Tracker, routes []*navigator.Route) {
				require.Equal(t, 2, trk.NumRoutes())
				require.Len(t, routes, 2)
				actualRoutes, err := trk.Routes()
				require.NoError(t, err)
				require.Equal(t, routes, actualRoutes)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := new(generator.Tracker)
			if tt.arrange != nil {
				tt.arrange(tr)
			}
			got, err := tr.AddRoutes(tt.args.newRoutes())
			if (err != nil) != tt.wantErr {
				t.Errorf("Tracker.AddRoutes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.assert != nil {
				tt.assert(tr, got)
			}
		})
	}
}
