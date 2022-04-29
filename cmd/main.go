package main

import (
	"bdlang/repl"
	"fmt"
	"os"
	"os/user"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Hi %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}
