package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/gpsgend/internal/transport/codes"
)

func errorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	statusCode := http.StatusInternalServerError

	if stdErr, ok := err.(*fiber.Error); ok {
		c.Status(stdErr.Code)
		stdErr.Code = codes.CodeUnknown
		return c.JSON(stdErr)
	}

	httpErr := fiber.ErrInternalServerError

	code := codes.FromError(err)
	if code > 0 {
		statusCode = codes.ToHTTP(code)
		httpErr.Code = code
		httpErr.Message = err.Error()
	} else if errors.Is(err, errBadRequest) {
		httpErr.Code = codes.CodeUnknown
		httpErr.Message = err.Error()
		statusCode = http.StatusBadRequest
	}

	c.Status(statusCode)

	return c.JSON(httpErr)
}
