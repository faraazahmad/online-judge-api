package main

import "net"

type server struct{}

func main() {
	// create a tcp listener at port 4040
	listener, err := net.Listen("tcp", ":4040")
	if err != nil {
		panic(err)
	}
}
