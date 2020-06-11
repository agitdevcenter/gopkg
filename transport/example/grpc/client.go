package main

import (
	"context"
	"github.com/agitdevcenter/gopkg/transport/example/grpc/handler"
	"google.golang.org/grpc"
	"log"
	"time"
)

func main() {
	conn, err := grpc.Dial(
		":2202",
		grpc.WithInsecure(),
	)

	helloClient := handler.NewHelloHandlerClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	helloRequest := &handler.HelloRequest{
		Name: "Bias Tegaralaga",
	}

	helloResponse, err := helloClient.Hello(ctx, helloRequest)
	if err != nil {
		log.Fatalf("Error calling hello, err : %+v", err)
	}

	log.Printf("Hello : %+v", helloResponse)

}
