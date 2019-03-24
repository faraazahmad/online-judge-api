package main

import (
	"os/user"
	"fmt"
)

func main() {
	fmt.Println(user.Current())
}
