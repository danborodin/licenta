package builtin

import (
	"bdlang/object"
)

var Functions = map[string]*object.Builtin{
	"len": {
		Fn: internalLen,
	},
	"print": {
		Fn: internalPrint,
	},
	"println": {
		Fn: internalPrintln,
	},
	"pow": {
		Fn: internalPow,
	},
	"wd": {
		Fn: internalWd,
	},
	"readln": {
		Fn: internalReadln,
	},
}
