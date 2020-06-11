package handler

import (
	Response "github.com/agitdevcenter/gopkg/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

const (
	ShowContentKey   = "content"
	ShowContentValue = "true"
)

type ResponseData struct {
	Elapsed int64 `json:"elapsed"`
}

type healthHandler struct{}

func NewHealthHandler() *healthHandler {
	return &healthHandler{}
}

func (h *healthHandler) defaultResponse(ctx echo.Context) *ResponseData {
	requestTime := ctx.Get("RequestTime").(time.Time)
	stop := time.Now()

	return &ResponseData{Elapsed: stop.Sub(requestTime).Nanoseconds() / 1000000}
}

func (h *healthHandler) Ping(ctx echo.Context) (err error) {
	content := ctx.QueryParam(ShowContentKey)

	if content == ShowContentValue {
		return ctx.JSON(http.StatusOK, &Response.DefaultResponse{
			Data: h.defaultResponse(ctx),
			Response: Response.Response{
				Status:  Response.SuccessCode,
				Message: http.StatusText(http.StatusOK),
			},
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}

func (h *healthHandler) PingError(ctx echo.Context) (err error) {
	content := ctx.QueryParam(ShowContentKey)

	var app []string

	if content == app[3] {
		return ctx.JSON(http.StatusOK, &Response.DefaultResponse{
			Data: h.defaultResponse(ctx),
			Response: Response.Response{
				Status:  Response.SuccessCode,
				Message: http.StatusText(http.StatusOK),
			},
		})
	}

	return ctx.NoContent(http.StatusNoContent)
}
