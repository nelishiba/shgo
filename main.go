package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/negli0/shgo/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hello %s! Thin is the SHGO programming language!\n", user.Username)

	repl.Start(os.Stdin, os.Stdout)
}
