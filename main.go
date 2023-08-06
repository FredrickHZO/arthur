package main

import (
	"arthur/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s. Arthur version 0.1\n", user.Username)
	fmt.Printf("Type in commands\n\n")
	repl.Init(os.Stdin, os.Stdout)
}
