package handler

import "google.golang.org/grpc"

type Handler interface {
	Register(server *grpc.Server)
}
