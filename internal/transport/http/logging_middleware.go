package http

import (
	"context"
	"net/http"
	"time"

	"log/slog"

	"github.com/gofiber/fiber/v2"
)

func LoggingMiddleware(logger *slog.Logger) fiber.Handler {
	return func(c *fiber.Ctx) error {
		start := time.Now()
		path := c.Path()
		err := c.Next()
		end := time.Now()
		latency := end.Sub(start)

		attributes := []slog.Attr{
			slog.Int("status", c.Response().StatusCode()),
			slog.String("method", string(c.Context().Method())),
			slog.String("path", path),
			slog.String("ip", c.Context().RemoteIP().String()),
			slog.Duration("latency", latency),
			slog.String("user-agent", string(c.Context().UserAgent())),
			slog.Time("time", end),
		}

		attributes = append(attributes, slog.Any("requestid", c.Locals("requestid")))
		ctx := context.Background()

		switch {
		case c.Response().StatusCode() >= http.StatusBadRequest && c.Response().StatusCode() < http.StatusInternalServerError:
			logger.LogAttrs(ctx, slog.LevelWarn, err.Error(), attributes...)
		case c.Response().StatusCode() >= http.StatusInternalServerError:
			logger.LogAttrs(ctx, slog.LevelError, err.Error(), attributes...)
		default:
			logger.LogAttrs(ctx, slog.LevelInfo, "Incoming request", attributes...)
		}

		if err != nil {
			logger.LogAttrs(ctx, slog.LevelError, err.Error(), attributes...)
		}

		return err
	}
}
