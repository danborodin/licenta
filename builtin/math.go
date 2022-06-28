package builtin

import (
	"bdlang/object"
	"math"
)

func internalPow(args ...object.Object) object.Object {
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
}
