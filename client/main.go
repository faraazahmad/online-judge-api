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
	g.GET("/ruby", func(ctx *gin.Context) {
		// get codeURL from request body
		codeURL := ctx.MustGet("url").(string)

		// get args from request body and split into []string
		args := strings.Split(ctx.MustGet("args").(string), ",")

		// get stdin from request body
		stdin := ctx.MustGet("stdin").(string)

		log.Println(ctx)
		log.Println(ctx)

		// create protobuf request
		req := &proto.Request{CodeURL: codeURL, Args: args, Stdin: []byte(stdin)}

		// send request and get response, error
		if response, err := client.Ruby(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprintf("%s", response.Body),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
	})

	g.GET("/test", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, ctx)
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
