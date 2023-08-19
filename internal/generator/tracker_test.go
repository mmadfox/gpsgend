package generator_test

import (
	"testing"

	"github.com/mmadfox/go-gpsgen"
	"github.com/mmadfox/go-gpsgen/geo"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/generator"
	"github.com/mmadfox/gpsgend/internal/types"
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

func TestTracker_RouteAt(t *testing.T) {
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		args    args
		arrange func(*generator.Tracker)
		assert  func(*generator.Tracker, *navigator.Route)
		wantErr bool
	}{
		{
			name: "should return error when routeIndex < 0",
			args: args{
				index: -1,
			},
			wantErr: true,
		},
		{
			name: "should return error when routeIndex > maxRoutes",
			args: args{
				index: 100,
			},
			arrange: func(trk *generator.Tracker) {
				trk.AddRoute(gpsgen.RandomRouteForMoscow())
			},
			wantErr: true,
		},
		{
			name: "should return error when routeSnapshot is invalid",
			args: args{
				index: 1,
			},
			arrange: func(trk *generator.Tracker) {
				trk.AddRoute(gpsgen.RandomRouteForMoscow())
				// inject invalid routesSnapshot field data
				generator.Debug_InjectInvalidDatatForTracker(trk, "routesSnapshot")
			},
			wantErr: true,
		},
		{
			name: "should return route by index when all params are valid",
			arrange: func(trk *generator.Tracker) {
				trk.AddRoute(gpsgen.RandomRouteForNewYork())
			},
			args: args{
				index: 1,
			},
			assert: func(trk *generator.Tracker, actualRoute *navigator.Route) {
				require.Equal(t, 1, trk.NumRoutes())
				expectedRoutes, err := trk.Routes()
				require.NoError(t, err)
				require.Equal(t, expectedRoutes[0], actualRoute)
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
			got, err := tr.RouteAt(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tracker.RouteAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.assert != nil {
				tt.assert(tr, got)
			}
		})
	}
}

func TestTracker_RouteByID(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(*generator.Tracker) string
		assert  func(*generator.Tracker, *navigator.Route)
		wantErr bool
	}{
		{
			name:    "should return error when tracker has no routes",
			wantErr: true,
		},
		{
			name: "should return error when routeSnapshot is invalid",
			arrange: func(trk *generator.Tracker) string {
				route := gpsgen.RandomRouteForParis()
				trk.AddRoute(route)
				// inject invalid routesSnapshot field data
				generator.Debug_InjectInvalidDatatForTracker(trk, "routesSnapshot")
				return route.ID()
			},
			assert: func(trk *generator.Tracker, route *navigator.Route) {
				require.Equal(t, 1, trk.NumRoutes())
			},
			wantErr: true,
		},
		{
			name: "should return error when route not found",
			arrange: func(trk *generator.Tracker) string {
				expectedRoute := gpsgen.RandomRouteForParis()
				trk.AddRoute(expectedRoute)
				otherRoute := gpsgen.RandomRouteForMoscow()
				return otherRoute.ID()
			},
			wantErr: true,
		},
		{
			name: "should return route when all params are valid",
			arrange: func(trk *generator.Tracker) string {
				expectedRoute := gpsgen.RandomRouteForParis()
				trk.AddRoute(expectedRoute)
				return expectedRoute.ID()
			},
			assert: func(trk *generator.Tracker, actual *navigator.Route) {
				require.Equal(t, 1, trk.NumRoutes())
				routes, err := trk.Routes()
				require.NoError(t, err)
				expected := routes[0]
				require.Equal(t, expected.ID(), actual.ID())
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := new(generator.Tracker)
			var routeID types.ID
			if tt.arrange != nil {
				if expectedRouteID := tt.arrange(tr); len(expectedRouteID) > 0 {
					rid, err := types.ParseID(expectedRouteID)
					require.NoError(t, err)
					routeID = rid
				}
			}
			got, err := tr.RouteByID(routeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tracker.RouteByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.assert != nil {
				tt.assert(tr, got)
			}
		})
	}
}

func TestTracker_RemoveRoute(t *testing.T) {
	tests := []struct {
		name    string
		arrange func(*generator.Tracker) string
		assert  func(*generator.Tracker, types.ID)
		wantErr bool
	}{
		{
			name:    "should return error when tracker has no routes",
			wantErr: true,
		},
		{
			name: "should return error when routeSnapshot is invalid",
			arrange: func(trk *generator.Tracker) string {
				route := gpsgen.RandomRouteForMoscow()
				trk.AddRoute(route)
				// inject invalid routesSnapshot field data
				generator.Debug_InjectInvalidDatatForTracker(trk, "routesSnapshot")
				return route.ID()
			},
			wantErr: true,
		},
		{
			name: "should return error when route not found",
			arrange: func(trk *generator.Tracker) string {
				route := gpsgen.RandomRouteForMoscow()
				trk.AddRoute(route)
				return types.NewID().String()
			},
			assert: func(trk *generator.Tracker, id types.ID) {
				require.Equal(t, 1, trk.NumRoutes())
			},
			wantErr: true,
		},
		{
			name: "should not return error when all params are valid",
			arrange: func(trk *generator.Tracker) string {
				route := gpsgen.RandomRouteForMoscow()
				trk.AddRoute(route)
				return route.ID()
			},
			assert: func(trk *generator.Tracker, _ types.ID) {
				require.Equal(t, 0, trk.NumRoutes())
				routes, err := trk.Routes()
				require.NoError(t, err)
				require.Len(t, routes, 0)
			},
			wantErr: false,
		},
		{
			name: "should not return error when all params are valid",
			arrange: func(trk *generator.Tracker) string {
				trk.AddRoute(gpsgen.RandomRouteForMoscow())
				route := gpsgen.RandomRouteForMoscow()
				trk.AddRoute(route)
				return route.ID()
			},
			assert: func(trk *generator.Tracker, _ types.ID) {
				require.Equal(t, 1, trk.NumRoutes())
				routes, err := trk.Routes()
				require.NoError(t, err)
				require.Len(t, routes, 1)
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := new(generator.Tracker)
			var routeID types.ID
			if tt.arrange != nil {
				if expectedRouteID := tt.arrange(tr); len(expectedRouteID) > 0 {
					rid, err := types.ParseID(expectedRouteID)
					require.NoError(t, err)
					routeID = rid
				}
			}
			err := tr.RemoveRoute(routeID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Tracker.RemoveRoute() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if tt.wantErr {
				return
			}
			if tt.assert != nil {
				tt.assert(tr, routeID)
			}
		})
	}
}

func TestTracker_ResetRoutes(t *testing.T) {
	trk := new(generator.Tracker)
	trk.AddRoute(gpsgen.RandomRouteForMoscow())
	trk.AddRoute(gpsgen.RandomRouteForMoscow())
	trk.AddRoute(gpsgen.RandomRouteForMoscow())
	routes, err := trk.Routes()
	require.NoError(t, err)
	require.Len(t, routes, 3)
	require.Equal(t, 3, trk.NumRoutes())

	trk.ResetRoutes()

	require.Equal(t, 0, trk.NumRoutes())
	routes, err = trk.Routes()
	require.NoError(t, err)
	require.Len(t, routes, 0)
	require.True(t, trk.HasNoRoutes())
}
