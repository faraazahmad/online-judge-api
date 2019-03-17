package main

import (
	proto "../proto"
	"../wget"

	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
)

func (s *server) Ruby(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	codeURL, args := request.GetCodeURL(), request.GetArgs()

	/*
		Code has to be downloaded in
		/home/${whoami}/remote/temp.rb
	*/

	// get home directory of current user
	homeDir, err := user.Current()
	if err != nil {
		return nil, err
	}
	// generate string for destination
	destinationString := fmt.Sprintf("%s/remote/temp.rb", homeDir)

	// download file in the provided destination
	wget.Wget(codeURL, destinationString)

	// TODO: execute the code with provided params
	// get Command struct instance by passing command name and arguments
	cmd := exec.Command("ruby", args...)

	var Stdout bytes.Buffer
	// point cmd.Stdout to output buffer
	cmd.Stdout = &Stdout

	// run the command and capture output
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
