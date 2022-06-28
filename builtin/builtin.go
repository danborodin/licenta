package builtin

import (
	"bdlang/object"
	"io"
	"log"
	"math"
	"os"
	"strconv"
)

var Functions = map[string]*object.Builtin{
	"len": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return object.NewError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return object.NewError("argument to `len` not supported, got %s", args[0].Type())
			}
		},
	},
	"print": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"println": {
		Fn: func(args ...object.Object) object.Object {
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
		},
	},
	"pow": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) <= 1 {
				return object.NewError("not enoughth arguments in function 'pow' call, got=%d, want 2", len(args))
			}

			var x int
			var y int

			switch arg := args[0].(type) {
			case *object.Integer:
				x = int(arg.Value)
			default:
				return object.NewError("first argument is not an integer, got=%s", arg.Type())
			}

			switch arg := args[1].(type) {
			case *object.Integer:
				y = int(arg.Value)
			default:
				return object.NewError("second argument is not an integer, got=%s", arg.Type())
			}

			return &object.Integer{Value: int64(math.Pow(float64(x), float64(y)))}
		},
	},
	"wd": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 0 {
				return object.NewError("pwd function should have 0 arguments, got=%d", len(args))
			}

			wd, err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}

			return &object.String{Value: wd}
		},
	},
}
