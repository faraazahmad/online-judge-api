package main

import (
	"context"
	"net"

	proto "../proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func main() {
	// create a tcp listener at port 4040
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// create new grpc server and register
	srv := grpc.NewServer()
	proto.RegisterExecServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}

func (s *server) Ruby(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	codeURL, params := request.GetCodeURL(), request.GetParams()

	// TODO: Donload the code

	// TODO: execute the code with provided params

	// TODO: Capture the output

	// TODO: check if there was an error

	// TODO: return the full response

	return nil, nil
}
