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
	file := os.Args[1]
	//fmt.Println(file)
	fileinput.Start(file)

	//fmt.Printf("bdlang REPL alpha version\n")
	//repl.Start(os.Stdin, os.Stdout)
}
