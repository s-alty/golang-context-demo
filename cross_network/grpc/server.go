package main

import (
	"net"
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"cross_network/grpc/service"
)


type server struct {
	service.UnimplementedNoopServer
}

// see if the value was present on the request context
// see if the deadline was set on the request context
// just do nothing until the context gets canceled
func (s *server) Noop(ctx context.Context, req *service.NoopRequest) (*service.NoopResponse, error){

	d, _ := metadata.FromIncomingContext(ctx)
	value := d["example"]
	fmt.Printf("Value retrieved from context for key 'example': %v\n", value)
	_, ok := ctx.Deadline()
	fmt.Printf("A deadline was present on the context?: %v\n", ok)

	<- ctx.Done() // wait for cancellation
	fmt.Println("request was cancelled")
	return &service.NoopResponse{}, nil
}

func main(){
	g := grpc.NewServer()
	service.RegisterNoopServer(g, &server{})
	listener, _ := net.Listen("tcp", "127.0.0.1:8082")
	g.Serve(listener)
}
