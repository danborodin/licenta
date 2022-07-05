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
	"cd": {
		Fn: internalCd,
	},
	"readln": {
		Fn: internalReadln,
	},
	"first": {
		Fn: internalArrayFirst,
	},
	"last": {
		Fn: internalArrayLast,
	},
	"rest": {
		Fn: internalArrayRest,
	},
	"push": {
		Fn: internalArrayPush,
	},
	"runCmd": {
		Fn: internalRunCmd,
	},
}
