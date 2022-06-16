package main

import (
	"bdlang/repl"
	"fmt"
	"os"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	fmt.Printf("bdlang REPL alpha version\n")
	repl.Start(os.Stdin, os.Stdout)
}
