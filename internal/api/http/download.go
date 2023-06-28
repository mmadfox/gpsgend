package http

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/mmadfox/go-gpsgen/geojson"
	"github.com/mmadfox/go-gpsgen/navigator"
	"github.com/mmadfox/gpsgend/internal/device"
)

func downloadFileFromURL(addr string) ([]byte, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
		CheckRedirect: func(r *http.Request, via []*http.Request) error {
			r.URL.Opaque = r.URL.Path
			return nil
		},
	}
	resp, err := client.Get(addr)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to download file %s", addr)
	}
	buf := bytes.NewBuffer(nil)
	io.Copy(buf, resp.Body)
	return buf.Bytes(), nil
}

func writeDownloadRoutesResponse(ctx *fiber.Ctx, deviceID uuid.UUID, routes []*device.Route) error {
	originRoutes := make([]*navigator.Route, len(routes))
	for i := 0; i < len(routes); i++ {
		originRoutes[i] = routes[i].Route()
	}

	data, err := geojson.Encode(originRoutes)
	if err != nil {
		return err
	}

	filename := fmt.Sprintf("device_%s_routes.geojson", deviceID.String())
	ctx.Response().Header.Set("Content-Disposition", "attachment; filename="+filename)
	ctx.Response().Header.Set("Content-Type", "application/geo+json")
	ctx.Response().Header.Set("Content-Length", strconv.Itoa(len(data)))
	_, err = ctx.Write(data)
	return err
}
