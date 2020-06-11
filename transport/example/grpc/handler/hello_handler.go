package handler

import (
	"context"
	Response "github.com/agitdevcenter/gopkg/response"
)

type helloHandler struct{}

func SetupHelloHandler() *helloHandler {
	return &helloHandler{}
}

func (h *helloHandler) Validate() HelloHandlerServer {
	return h
}

func (h *helloHandler) Hello(ctx context.Context, request *HelloRequest) (response *HelloResponse, err error) {
	response = &HelloResponse{
		Status:  Response.SuccessCode,
		Message: "Success",
		Data: HelloData{
			Hello: request.Name,
		},
	}
	return
}
