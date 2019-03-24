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

// Request : A structre to hold request info
type Request struct {
	URL   string `json:"url"`
	Args  string `json:"args"`
	Stdin string `json:"string"`
}

// DRY function to handle errors
func checkError(err error, errorCode int, ctx *gin.Context) {
	if err != nil {
		ctx.JSON(errorCode, gin.H{
			"error": err.Error(),
		})
	}
}

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
		// struct instance to store request body
		var requestJSON Request

		// return error if there was an error in binding JSON
		err := ctx.BindJSON(&requestJSON)
		checkError(err, http.StatusInternalServerError, ctx)

		// assign code URL
		codeURL := requestJSON.URL

		// get args and split into []string
		args := strings.Split(requestJSON.Args, ",")

		// get stdin
		stdin := requestJSON.Stdin

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

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}

}
