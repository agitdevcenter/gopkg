package handler

import (
	"google.golang.org/grpc"
)

type Handler struct {
	Hello HelloHandlerServer
}

func New() *Handler {
	return &Handler{Hello: SetupHelloHandler().Validate()}
}

func (h *Handler) Register(server *grpc.Server) {
	RegisterHelloHandlerServer(server, h.Hello)
}
