package http

import (
	"net/http"
	"time"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
	"github.com/mmadfox/gpsgend/internal/device"
	"github.com/mmadfox/gpsgend/internal/publisher"
	"golang.org/x/exp/slog"
)

const (
	minUploadFileSize = 8
	maxUploadFileSize = 3000000
	readTimeout       = 5 * time.Second
	writeTimeout      = 10 * time.Second
	writeWait         = 10 * time.Second
	pingPeriod        = (pongWait * 9) / 10
	pongWait          = 60 * time.Second
)

type Server struct {
	handler *fiber.App
}

func NewServer(
	deviceUseCase device.UseCase,
	deviceQuery device.Query,
	publisher *publisher.Publisher,
	logger *slog.Logger,
) *Server {
	server := &Server{
		handler: fiber.New(fiber.Config{
			ServerHeader:          "gpsgend",
			StrictRouting:         true,
			ReadTimeout:           readTimeout,
			WriteTimeout:          writeTimeout,
			ErrorHandler:          errorHandler,
			DisableStartupMessage: true,
		}),
	}

	server.handler.Use(requestid.New())
	server.handler.Use(recover.New())
	server.handler.Use(pprof.New(pprof.Config{Prefix: "/debug"}))
	server.handler.Use(loggingMiddleware(logger))
	server.handler.Use("/gpsgend/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	server.handler.Get("/gpsgend/ws", websocket.New(func(c *websocket.Conn) {
		ticker := time.NewTicker(pingPeriod)
		client := publisher.NewClient()

		defer func() {
			ticker.Stop()
			publisher.CloseClient(client.ID)
			c.Close()
		}()

		for {
			select {
			case msg, ok := <-client.Out:
				c.SetWriteDeadline(time.Now().Add(writeWait))
				if !ok {
					c.WriteMessage(websocket.CloseMessage, nil)
					return
				}

				w, err := c.NextWriter(websocket.BinaryMessage)
				if err != nil {
					return
				}

				if _, err := w.Write(msg); err != nil {
					logger.Error("unable to write data message to websocket", "err", err)
					return
				}
			case <-ticker.C:
				logger.Debug("ping message")

				c.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
					logger.Error("unable to write ping message to websocket", "err", err)
					return
				}
			}
		}
	}))

	server.handler.Get("/gpsgend/v1/devices", searchDevicesHandler(deviceQuery))

	server.handler.Post("/gpsgend/v1/devices", createDeviceHandler(deviceUseCase))
	server.handler.Patch("/gpsgend/v1/devices/:deviceID", updateDeviceHandler(deviceUseCase))
	server.handler.Delete("/gpsgend/v1/devices/:deviceID", removeDeviceHandler(deviceUseCase))
	server.handler.Patch("/gpsgend/v1/devices/:deviceID/run", runDeviceHandler(deviceUseCase))
	server.handler.Patch("/gpsgend/v1/devices/:deviceID/stop", stopDeviceHandler(deviceUseCase))
	server.handler.Patch("/gpsgend/v1/devices/:deviceID/pause", pauseDeviceHandler(deviceUseCase))
	server.handler.Patch("/gpsgend/v1/devices/:deviceID/resume", resumeDeviceHandler(deviceUseCase))
	server.handler.Post("/gpsgend/v1/devices/:deviceID/sensors", addSensorHandler(deviceUseCase))
	server.handler.Get("/gpsgend/v1/devices/:deviceID/sensors", sensorsHandler(deviceUseCase))
	server.handler.Delete("/gpsgend/v1/devices/:deviceID/sensors/:sensorID", removeSensorHandler(deviceUseCase))
	server.handler.Post("/gpsgend/v1/devices/:deviceID/routes/upload", uploadRouteHandler(deviceUseCase))
	server.handler.Post("/gpsgend/v1/devices/:deviceID/routes/import", importRouteHandler(deviceUseCase))
	server.handler.Get("/gpsgend/v1/devices/:deviceID/routes/download", downloadRouteHandler(deviceUseCase))
	server.handler.Get("/gpsgend/v1/devices/:deviceID/routes", routesHandler(deviceUseCase))
	server.handler.Delete("/gpsgend/v1/devices/:deviceID/routes/:routeID", removeRouteHandler(deviceUseCase))
	server.handler.Get("/gpsgend/v1/generator/routes/random", randomRouteHandler())

	return server
}

func (s *Server) Listen(addr string) error {
	return s.handler.Listen(addr)
}

func (s *Server) Close() error {
	if s.handler == nil {
		return nil
	}
	return s.handler.Shutdown()
}

func (s *Server) invoke(r *http.Request) (*http.Response, error) {
	return s.handler.Test(r, 50)
}
