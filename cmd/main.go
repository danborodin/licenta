package main

import (
	fileinput "bdlang/fileInput"
	"bdlang/repl"
	"fmt"
	"os"
)

func main() {
	// start := time.Now()
	if len(os.Args) == 2 {
		file := os.Args[1]
		fileinput.Start(file)
	} else {
		fmt.Printf("bdlang REPL alpha version\n")
		repl.Start(os.Stdin, os.Stdout)
	}

	// end := time.Since(start)
	// fmt.Println(end)
}
