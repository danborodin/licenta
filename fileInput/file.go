package fileinput

import (
	"bdlang/evaluator"
	"bdlang/lexer"
	"bdlang/object"
	"bdlang/parser"
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func Start(fileName string) {
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	env := object.NewEnvironment()

	var line string = ""
	for {
		scanned := scanner.Scan()
		if !scanned {
			break
		}
		l := scanner.Text()
		line = line + l
	}

	line = strings.TrimSpace(line)

	l := lexer.New(line)
	p := parser.New(l)

	program := p.ParseProgram()
	if len(p.Errors()) != 0 {
		printParseErrors(os.Stdout, p.Errors())
	}

	evaluated := evaluator.Eval(program, env)
	switch evaluated.(type) {
	case *object.Error:
		io.WriteString(os.Stdout, evaluated.Inspect())
		io.WriteString(os.Stdout, "\n")
		os.Exit(0)
	}
}

func printParseErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\t")
	}
}
