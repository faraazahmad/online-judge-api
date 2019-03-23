package ruby

import (
	"net/http"
	"strings"

	proto "../proto"
	"../wget"
	"github.com/gin-gonic/gin"

	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

type server struct{}

func (s *server) run(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	// extract the code URL and its arguments from the request
	codeURL, args := request.GetCodeURL(), request.GetArgs()

	/*
		Code has to be downloaded in
		/home/${whoami}/remote/ruby/temp.rb
	*/

	// get home directory of current user
	homeDir, err := user.Current()
	if err != nil {
		return nil, err
	}
	// generate string for destination (in UNIX based systems)
	destinationString := fmt.Sprintf("%s/remote/ruby/temp.rb", homeDir)

	// download file in the provided destination
	wget.Wget(codeURL, destinationString)

	// get Command struct instance by passing command name and arguments
	cmd := exec.Command("ruby", args...)

	// provide stdin to command
	cmd.Stdin = bytes.NewReader(request.GetStdin())

	// store cmd.Stdout in a Bytes buffer
	var Stdout bytes.Buffer
	cmd.Stdout = &Stdout

	// run the command
	err = cmd.Run()
	if err != nil {
		return nil, err
	}

	// delete the code file
	err = os.Remove(destinationString)
	if err != nil {
		log.Println(err)
	}

	// return full response
	return &proto.Response{Body: Stdout.Bytes()}, nil
}

func handler(ctx *gin.Context) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// get codeURL from request body
		codeURL := ctx.PostForm("url")

		// get args from request body and split into []string
		args := strings.Split(ctx.PostForm("args"), ",")

		// get stdin from request body
		stdin := ctx.PostForm("stdin")

		// create protobuf request
		req := &proto.Request{CodeURL: codeURL, Args: args, Stdin: []byte(stdin)}

		// send request and get response, error
		if response, err := run(ctx, req); err == nil {
			ctx.JSON(http.StatusOK, gin.H{
				"result": fmt.Sprintf("%s", response.Body),
			})
		} else {
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error()})
		}
	}
}
