package device

import (
	"github.com/google/uuid"
	"github.com/lucasb-eyer/go-colorful"
	"github.com/mmadfox/go-gpsgen/navigator"
)

type Route struct {
	id    uuid.UUID
	route *navigator.Route
	color colorful.Color
}

func RouteFrom(id uuid.UUID, c colorful.Color, r *navigator.Route) *Route {
	return &Route{
		id:    id,
		color: c,
		route: r,
	}
}

func NewRoute(r *navigator.Route) *Route {
	return &Route{id: uuid.New(), route: r, color: randomColor()}
}

func (r *Route) ID() uuid.UUID {
	return r.id
}

func (r *Route) Route() *navigator.Route {
	return r.route
}

func (r *Route) Color() colorful.Color {
	return r.color
}
