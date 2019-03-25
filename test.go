package main

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
)

func main() {
	args := []string{"-a", "/home/faraaz/rpc/ruby/code-20190325061029.rb"}

	// get Command struct instance by passing command name and arguments
	cmd := exec.Command("ruby", args...)

	// provide stdin to command
	cmd.Stdin = bytes.NewReader([]byte("hello world"))

	// store cmd.Stdout in a Bytes buffer
	var Stdout bytes.Buffer
	cmd.Stdout = &Stdout

	// run the command
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(string(Stdout.Bytes()))
}
