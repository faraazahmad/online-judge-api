package main

import (
	"context"
	"fmt"
	"net"
	"os/user"

	proto "../proto"
	"../wget"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

// User : struct to hold info about current user
type User struct {
	// Uid is the user ID.
	// On POSIX systems, this is a decimal number representing the uid.
	// On Windows, this is a security identifier (SID) in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Uid string
	// Gid is the primary group ID.
	// On POSIX systems, this is a decimal number representing the gid.
	// On Windows, this is a SID in a string format.
	// On Plan 9, this is the contents of /dev/user.
	Gid string
	// Username is the login name.
	Username string
	// Name is the user's real or display name.
	// It might be blank.
	// On POSIX systems, this is the first (or only) entry in the GECOS field
	// list.
	// On Windows, this is the user's display name.
	// On Plan 9, this is the contents of /dev/user.
	Name string
	// HomeDir is the path to the user's home directory (if they have one).
	HomeDir string
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

func (s *server) Ruby(ctx context.Context, request *proto.Request) (*proto.Response, error) {
	codeURL, params := request.GetCodeURL(), request.GetParams()

	/*
		Code has to be downloaded in
		/home/${whoami}/remote/temp.rb
	*/

	// get home directory of current user
	homeDir, err := user.Current()
	if err != nil {
		panic(err)
	}
	// generate string for destination
	destinationString := fmt.Sprintf("%s/remote/temp.rb", homeDir)

	// download file in the provided destination
	wget.Wget(codeURL, destinationString)

	// TODO: execute the code with provided params

	// TODO: Capture the output

	// TODO: check if there was an error

	// TODO: return the full response

	return nil, nil
}
