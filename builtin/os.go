package builtin

import (
	"bdlang/object"
	"bytes"
	"log"
	"os"
	"os/exec"
)

func internalWd(args ...object.Object) object.Object {
	if len(args) != 0 {
		return object.NewError("wd function should have 0 arguments, got=%d", len(args))
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}

	return &object.String{Value: wd}
}

func internalCd(args ...object.Object) object.Object {
	if len(args) != 1 {
		return object.NewError("cd function should have 1 arguments, got=%d", len(args))
	}

	dir := args[0].(*object.String)

	err := os.Chdir(dir.Value)
	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func internalRunCmd(args ...object.Object) object.Object {

	if len(args) < 1 {
		return object.NewError("createCmd function should have at least 1 argument, got=%d", len(args))
	}

	cmd := args[0].(*object.String)
	var localArgs = make([]string, 0)

	if len(args) > 1 {
		for i, v := range args {
			if i > 0 {
				lArg := v.(*object.String)
				localArgs = append(localArgs, lArg.Value)
			}
		}
	}

	command := exec.Command(cmd.Value, localArgs...)

	var out bytes.Buffer
	command.Stdout = &out

	err := command.Run()

	if err != nil {
		log.Fatal(err)
	}

	return &object.String{Value: out.String()}

	//mt.Println(out.String())
	//return object.NULL
}
