package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/VivekAlhat/Nova/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to Novalang Playground, %s!\n", user.Username)
	fmt.Println("Start typing the code below.")
	repl.Start(os.Stdin, os.Stdout)
}
