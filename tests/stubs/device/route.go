package stubdevice

import (
	"github.com/mmadfox/go-gpsgen/route"
	"github.com/mmadfox/gpsgend/internal/device"
)

func Routes() []*device.Route {
	originRoutes, err := route.RoutesForRussia()
	if err != nil {
		panic(err)
	}
	routes := make([]*device.Route, len(originRoutes))
	for i := 0; i < len(originRoutes); i++ {
		routes[i] = device.NewRoute(originRoutes[i])
	}
	return routes
}
