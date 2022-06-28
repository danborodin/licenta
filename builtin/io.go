package builtin

import (
	"bdlang/object"
	"bufio"
	"io"
	"os"
	"strconv"
)

func internalPrint(args ...object.Object) object.Object {
	if len(args) <= 0 {
		return object.NewError("not enoughth arguments in function 'print' call, got=%d, want 1 or more", len(args))
	}

	var str string = ""

	for i := range args {
		switch arg := args[i].(type) {
		case *object.String:
			str = str + arg.Value
		case *object.Integer:
			str = str + strconv.Itoa(int(arg.Value))
		case *object.Boolean:
			str = str + strconv.FormatBool(arg.Value)
		}
	}

	io.WriteString(os.Stdout, str)

	return nil
}

func internalPrintln(args ...object.Object) object.Object {
	if len(args) <= 0 {
		return object.NewError("not enoughth arguments in function 'println' call, got=%d, want 1 or more", len(args))
	}

	var str string = ""

	for i := range args {
		switch arg := args[i].(type) {
		case *object.String:
			str = str + arg.Value
		case *object.Integer:
			str = str + strconv.Itoa(int(arg.Value))
		case *object.Boolean:
			str = str + strconv.FormatBool(arg.Value)
		}
	}

	io.WriteString(os.Stdout, str)
	io.WriteString(os.Stdout, "\n")

	return nil
}

func internalReadln(args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("readln function should have 0 arguments, got=%d", len(args))
	}

	scanner := bufio.NewScanner(os.Stdin)
	scanned := scanner.Scan()
	if !scanned {
		return nil
	}

	line := scanner.Text()

	number, err := strconv.Atoi(line)
	if err == nil {
		return &object.Integer{Value: int64(number)}
	}

	boolean, err := strconv.ParseBool(line)
	if err == nil {
		if boolean {
			return object.TRUE
		} else {
			return object.FALSE
		}
	}

	return &object.String{Value: line}
}
