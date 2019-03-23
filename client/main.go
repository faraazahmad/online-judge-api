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

	_client := proto.NewExecServiceClient(conn)
	g := gin.Default()

	/*
		Get codeURL, args and stdin from body
	*/
	g.GET("/ruby", func(ctx *gin.Context) {
		/*
			codeURL := ctx.Param("url")
			args := strings.Split(ctx.Param("args"), ",")

			req := &proto.Request{CodeURL: codeURL, Args: args}
			if response, err := client.Ruby(ctx, req); err == nil {
				ctx.JSON(http.StatusOK, gin.H{
					"result": fmt.Sprintf("%s", response.Body),
				})
			} else {
				ctx.JSON(http.StatusInternalServerError, gin.H{
					"error": err.Error()})
			}
		*/

		// get codeURL from request body
		codeURL := ctx.PostForm("url")

		// get args from request body
		args := ctx.PostForm("args")

		// get stdin from request body
		Stdin := ctx.PostForm("stdin")
	})

	if err := g.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
