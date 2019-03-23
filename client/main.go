package main

import (
	"log"

	proto "../proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	// connect to localhost:4040 without HTTPS
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewExecServiceClient(conn)
	g := gin.Default()

	/*
		Get codeURL, args and stdin from body
	*/
	g.GET("/ruby", ruby.handler)

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
