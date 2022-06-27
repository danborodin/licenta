package main

import (
	fileinput "bdlang/fileInput"
	"os"
)

func main() {
	// user, err := user.Current()
	// if err != nil {
	// 	panic(err)
	// }
	//if len(os.Args) > 2 {
	file := os.Args[1]
	fileinput.Start(file)
	// } else {
	// 	fmt.Printf("bdlang REPL alpha version\n")
	// 	repl.Start(os.Stdin, os.Stdout)
	// }
}
