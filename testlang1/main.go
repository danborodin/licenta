package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

var HadError bool

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: tl [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error readin from file")
	}
	run(string(bytes))
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, _, _ := reader.ReadLine()
		if string(line) == "" {
			break
		}
		run(string(line))
		HadError = false
	}
}

func run(source string) {
	//get tokens
	//print tokens
	if HadError {
		os.Exit(65)
	}
	fmt.Println(source)
}

func programError(line int, msg string) {
	report(line, "", msg)
}

func report(line int, where, msg string) {
	fmt.Println("[line ", line, "] Error", where, ": ", msg)
}
