package server

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"time"

	proto "github.com/faraazahmad/online_judge_api/proto"
	"github.com/faraazahmad/online_judge_api/wget"
)

// Interpreted - function to execute interpreted code and return its result
func (s *Server) Interpreted(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	// extract the code URL, args, and stdin from request
	codeURL, args, stdin := request.GetCodeURL(), request.GetArgs(), request.GetStdin()

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

	/*
		If no arguments were provided, only leave the
		destinationString in the args slice otherwise
		append location of file to arguments list
	*/
	if args[0] == "" {
		args = []string{destinationString}
	} else {
		args = append(args, destinationString)
	}

	// get Command struct instance by passing command name and arguments
	cmd := exec.Command("ruby", args...)

	// provide stdin to command
	cmd.Stdin = bytes.NewReader(stdin)

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
