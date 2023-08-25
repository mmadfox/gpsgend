package websocket

import (
	"time"

	"log/slog"

	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	transporthttp "github.com/mmadfox/gpsgend/internal/transport/http"
)

const (
	readTimeout  = 5 * time.Second
	writeTimeout = 10 * time.Second
	writeWait    = 10 * time.Second
	pingPeriod   = (pongWait * 9) / 10
	pongWait     = 60 * time.Second
)

type Server struct {
	addr   string
	broker Broker

	handler *fiber.App
}

func New(
	addr string,
	broker Broker,
	logger *slog.Logger,
) *Server {
	srv := Server{
		addr:   addr,
		broker: broker,
	}

	app := fiber.New(fiber.Config{
		ServerHeader:          "gpsgend",
		StrictRouting:         true,
		ReadTimeout:           readTimeout,
		WriteTimeout:          writeTimeout,
		DisableStartupMessage: true,
	})

	app.Use(transporthttp.LoggingMiddleware(logger))

	app.Use("/gpsgend/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	app.Get("/gpsgend/ws", websocket.New(func(c *websocket.Conn) {
		ticker := time.NewTicker(pingPeriod)

		cid := uuid.New()
		cli := newClient()
		broker.RegisterClient(cid, cli)

		defer func() {
			broker.Unregister(cid)
			ticker.Stop()
			c.Close()
		}()

		for {
			select {
			case <-cli.closeCh:
				return
			case msg, ok := <-cli.out:
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
					logger.Error("failed to write data message to websocket", "err", err)
					return
				}
			case <-ticker.C:
				logger.Debug("ping message")

				c.SetWriteDeadline(time.Now().Add(writeWait))
				if err := c.WriteMessage(websocket.PingMessage, nil); err != nil {
					logger.Error("filed to write ping message to websocket", "err", err)
					return
				}
			}
		}
	}))

	srv.handler = app

	return &srv
}

func (s *Server) Listen() error {
	return s.handler.Listen(s.addr)
}

func (s *Server) Close() error {
	return s.handler.Shutdown()
}
