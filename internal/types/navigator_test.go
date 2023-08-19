package types

import (
	"testing"

	"github.com/mmadfox/go-gpsgen"
	"github.com/stretchr/testify/require"
)

func TestNavigator_NavigatorFromProc(t *testing.T) {
	tracker := gpsgen.NewBicycleTracker()
	route := gpsgen.RandomRouteForMoscow()
	tracker.AddRoute(route)
	totalDistance := tracker.Distance()
	distanceInMeters := totalDistance / 2
	tracker.DestinationTo(distanceInMeters)
	nav := NavigatorFromProc(tracker)
	require.NotZero(t, nav.Distance)
	require.NotZero(t, nav.Lat)
	require.NotZero(t, nav.Lon)
	require.NotZero(t, nav.RouteDistance)
	require.NotZero(t, nav.TrackDistance)
	require.NotZero(t, nav.SegmentDistance)
	require.Zero(t, nav.RouteIndex)
	require.NotZero(t, nav.TrackIndex)
	require.NotZero(t, nav.SegmentIndex)
}
