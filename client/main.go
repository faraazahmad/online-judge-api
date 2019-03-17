package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	proto "../proto"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:4040", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}

	client := proto.NewExecServiceClient(conn)
	g := gin.Default()

	g.GET("/ruby/:url/:args", func(ctx *gin.Context) {
		url := ctx.Param("url")
		args := strings.Split(ctx.Param("args"), ",")

		req := &proto.Request{CodeURL: url, Args: args}
		if response, err := client.Ruby(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprintf("%d", response.Result),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
