package builtin

import (
	"bdlang/object"
	"log"
	"os"
)

func internalWd(args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("pwd function should have 0 arguments, got=%d", len(args))
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &object.String{Value: wd}
}
