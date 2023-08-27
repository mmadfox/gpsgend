package http

import (
	"log/slog"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/mmadfox/gpsgend/internal/generator"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
)

type Server struct {
	addr      string
	generator generator.Service
	handler   *fiber.App
}

func New(addr string, gen generator.Service, logger *slog.Logger) *Server {
	srv := Server{addr: addr, generator: gen}

	app := fiber.New(fiber.Config{
		ServerHeader:          "gpsgend",
		StrictRouting:         true,
		ReadTimeout:           readTimeout,
		WriteTimeout:          writeTimeout,
		ErrorHandler:          errorHandler,
		DisableStartupMessage: true,
	})

	app.Use(requestid.New())
	app.Use(recover.New())
	app.Use(pprof.New(pprof.Config{Prefix: "/debug"}))
	app.Use(LoggingMiddleware(logger))

	app.Get(v1("/trackers"), srv.searchTrackers)
	app.Post(v1("/trackers"), srv.addTracker)
	app.Delete(v1("/trackers/:trackerID"), srv.removeTracker)
	app.Patch(v1("/trackers/:trackerID"), srv.updateTracker)
	app.Get(v1("/trackers/:trackerID"), srv.findTracker)
	app.Patch(v1("/trackers/:trackerID/start"), srv.startTracker)
	app.Patch(v1("/trackers/:trackerID/stop"), srv.stopTracker)
	app.Patch(v1("/trackers/:trackerID/state"), srv.getTrackerState)
	app.Post(v1("/trackers/:trackerID/routes"), srv.addRoutes)
	app.Delete(v1("/trackers/:trackerID/routes"), srv.resetRoutes)
	app.Delete(v1("/trackers/:trackerID/routes/:routeID"), srv.removeRoute)
	app.Get(v1("/trackers/:trackerID/routes/:routeID"), srv.getRoute)
	app.Get(v1("/trackers/:trackerID/routes/at/:routeIndex"), srv.getRouteByIndex)
	app.Get(v1("/trackers/:trackerID/routes"), srv.getRoutes)
	app.Post(v1("/trackers/:trackerID/navigator/reset"), srv.resetNavigator)
	app.Post(v1("/trackers/:trackerID/navigator/next"), srv.toNextRoute)
	app.Post(v1("/trackers/:trackerID/navigator/prev"), srv.toPrevRoute)
	app.Post(v1("/trackers/:trackerID/navigator/move-to/at/:routeIndex"), srv.moveToRoute)
	app.Post(v1("/trackers/:trackerID/navigator/move-to/:routeID"), srv.moveToRouteByID)
	app.Post(v1("/trackers/:trackerID/navigator/move-to/at/:routeIndex/:trackIndex"), srv.moveToTrack)
	app.Post(v1("/trackers/:trackerID/navigator/move-to/:routeID/:trackID"), srv.moveToTrackByID)
	app.Post(v1("/trackers/:trackerID/navigator/move-to/at/:routeIndex/:trackIndex/:segmentIndex"), srv.moveToSegment)
	app.Post(v1("/trackers/:trackerID/sensors"), srv.addSensor)
	app.Delete(v1("/trackers/:trackerID/sensors/:sensorID"), srv.removeSensor)
	app.Get(v1("/trackers/:trackerID/sensors"), srv.getSensors)
	app.Patch(v1("/trackers/:trackerID/shutdown"), srv.shutdown)
	app.Patch(v1("/trackers/:trackerID/resume"), srv.resume)

	srv.handler = app

	return &srv
}

func (s *Server) Listen() error {
	return s.handler.Listen(s.addr)
}

func (s *Server) Close() error {
	return s.handler.Shutdown()
}

func v1(s string) string {
	return "/gpsgend/v1" + s
}

func (s *Server) searchTrackers(c *fiber.Ctx) error {
	filter, err := decodeSearchTrackersRequest(c)
	if err != nil {
		return err
	}

	result, err := s.generator.SearchTrackers(c.Context(), filter)
	if err != nil {
		return err
	}

	return c.JSON(result)
}

func (s *Server) addTracker(c *fiber.Ctx) error {
	opts, err := decodeAddTrackerRequest(c.Request())
	if err != nil {
		return err
	}

	newTracker, err := s.generator.NewTracker(c.Context(), opts)
	if err != nil {
		return err
	}

	return encodeTrackerResponse(c, newTracker)
}

func (s *Server) removeTracker(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.RemoveTracker(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) updateTracker(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	opts, err := decodeUpdateTrackerRequest(c)
	if err != nil {
		return err
	}

	if err := s.generator.UpdateTracker(c.Context(), trackerID, opts); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) findTracker(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	tracker, err := s.generator.FindTracker(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return encodeTrackerResponse(c, tracker)
}

func (s *Server) startTracker(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.StartTracker(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) stopTracker(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.StopTracker(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) getTrackerState(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	state, err := s.generator.TrackerState(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return encodeTrackerStateResponse(c, state)
}

func (s *Server) addRoutes(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	routes, err := decodeAddRoutesRequest(c)
	if err != nil {
		return err
	}

	if err := s.generator.AddRoutes(c.Context(), trackerID, routes); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) removeRoute(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeID, err := decodeID(c, "routeID")
	if err != nil {
		return err
	}

	if err := s.generator.RemoveRoute(c.Context(), trackerID, routeID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) getRoutes(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	routes, err := s.generator.Routes(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return encodeRoutes(c, routes)
}

func (s *Server) getRoute(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeID, err := decodeID(c, "routeID")
	if err != nil {
		return err
	}

	route, err := s.generator.RouteByID(c.Context(), trackerID, routeID)
	if err != nil {
		return err
	}

	return encodeRoute(c, route)
}

func (s *Server) getRouteByIndex(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeIndex, err := decodeIndex(c, "routeIndex")
	if err != nil {
		return err
	}

	route, err := s.generator.RouteAt(c.Context(), trackerID, routeIndex)
	if err != nil {
		return err
	}

	return encodeRoute(c, route)
}

func (s *Server) resetRoutes(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	ok, err := s.generator.ResetRoutes(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return encodeFlagResponse(c, ok)
}

func (s *Server) resetNavigator(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.ResetNavigator(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) toNextRoute(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.ToNextRoute(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) toPrevRoute(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.ToPrevRoute(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) moveToRoute(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeIndex, err := decodeIndex(c, "routeIndex")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.MoveToRoute(c.Context(), trackerID, routeIndex)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) moveToRouteByID(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeID, err := decodeID(c, "routeID")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.MoveToRouteByID(c.Context(), trackerID, routeID)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) moveToTrack(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeIndex, err := decodeIndex(c, "routeIndex")
	if err != nil {
		return err
	}
	trackIndex, err := decodeIndex(c, "trackIndex")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.MoveToTrack(c.Context(), trackerID, routeIndex, trackIndex)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) moveToTrackByID(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeID, err := decodeID(c, "routeID")
	if err != nil {
		return err
	}
	trackID, err := decodeID(c, "trackID")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.MoveToTrackByID(c.Context(), trackerID, routeID, trackID)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) moveToSegment(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	routeIndex, err := decodeIndex(c, "routeIndex")
	if err != nil {
		return err
	}
	trackIndex, err := decodeIndex(c, "trackIndex")
	if err != nil {
		return err
	}
	segmentIndex, err := decodeIndex(c, "segmentIndex")
	if err != nil {
		return err
	}

	nav, ok, err := s.generator.MoveToSegment(
		c.Context(),
		trackerID,
		routeIndex,
		trackIndex,
		segmentIndex,
	)
	if err != nil {
		return err
	}

	return c.JSON(navigatorResponse{
		Navigator: nav,
		Ok:        ok,
	})
}

func (s *Server) addSensor(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	newSensor, err := decodeAddSensorRequest(c)
	if err != nil {
		return err
	}

	if err := s.generator.AddSensor(c.Context(), trackerID, newSensor); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) removeSensor(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}
	sensorID, err := decodeID(c, "sensorID")
	if err != nil {
		return err
	}

	if err := s.generator.RemoveSensor(c.Context(), trackerID, sensorID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) getSensors(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	sensors, err := s.generator.Sensors(c.Context(), trackerID)
	if err != nil {
		return err
	}

	return encodeSensorsResponse(c, sensors)
}

func (s *Server) shutdown(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.ShutdownTracker(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}

func (s *Server) resume(c *fiber.Ctx) error {
	trackerID, err := decodeID(c, "trackerID")
	if err != nil {
		return err
	}

	if err := s.generator.ResumeTracker(c.Context(), trackerID); err != nil {
		return err
	}

	return encodeSuccessResponse(c)
}
