package main

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/exec"
	"os/user"
	"time"

	proto "../proto"
	"../wget"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// create a tcp listener at port 4040
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}

	// create new grpc server and register
	srv := grpc.NewServer()
	proto.RegisterExecServiceServer(srv, &Server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
