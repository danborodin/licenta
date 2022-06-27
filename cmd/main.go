package main

import (
	fileinput "bdlang/fileInput"
	"bdlang/repl"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		file := os.Args[1]
		fileinput.Start(file)
	} else {
		fmt.Printf("bdlang REPL alpha version\n")
		repl.Start(os.Stdin, os.Stdout)
	}
}
