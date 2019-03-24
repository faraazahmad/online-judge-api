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

type server struct{}

func (s *server) Ruby(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	// extract the code URL and its arguments from the request
	codeURL, args := request.GetCodeURL(), request.GetArgs()

	/*
		Code has to be downloaded in
		/home/${whoami}/remote/ruby/code-#{time.Now()}.rb
	*/

	// get home directory of current user
	currentUser, err := user.Current()
	if err != nil {
		return nil, err
	}

	// get and format current time for every request
	t := time.Now().Format("20060102150405")

	// generate string for destination (in UNIX based systems)
	destinationString := fmt.Sprintf("%s/rpc/ruby/code-%s.rb", currentUser.HomeDir, t)

	// download file in the provided destination
	wget.Wget(codeURL, destinationString)

	// generate main commmand (without args)
	mainCmd := fmt.Sprintf("ruby %s", destinationString)

	// get Command struct instance by passing command name and arguments
	cmd := exec.Command(mainCmd, args...)

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
