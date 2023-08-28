package http

import (
	"errors"
	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/mmadfox/gpsgend/internal/transport/codes"
)

const serviceName = "gpsgend"

type httpErr struct {
	Service   string `json:"service"`
	Code      int    `json:"code"`
	Message   string `json:"message,omitempty"`
	RequestID string `json:"requestId"`
}

func errorHandler(c *fiber.Ctx, err error) error {
	c.Set(fiber.HeaderContentType, fiber.MIMEApplicationJSONCharsetUTF8)

	rid := c.Locals("requestid").(string)

	statusCode := http.StatusInternalServerError

	if stdErr, ok := err.(*fiber.Error); ok {
		c.Status(stdErr.Code)
		stdErr.Code = codes.CodeUnknown
		return c.JSON(httpErr{
			Service:   serviceName,
			Code:      codes.CodeUnknown,
			Message:   stdErr.Message,
			RequestID: rid,
		})
	}

	herr := httpErr{
		Service:   serviceName,
		Code:      http.StatusInternalServerError,
		RequestID: rid,
	}

	code := codes.FromError(err)
	if code > 0 {
		statusCode = codes.ToHTTP(code)
		herr.Code = code
		if uerr := errors.Unwrap(err); uerr != nil {
			herr.Message = uerr.Error()
		} else {
			herr.Message = err.Error()
		}
	} else if errors.Is(err, errBadRequest) {
		herr.Code = codes.CodeUnknown
		herr.Message = err.Error()
		statusCode = http.StatusBadRequest
	}

	c.Status(statusCode)

	return c.JSON(herr)
}
