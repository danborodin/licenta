package main

import "fmt"

type (
	Value      float64
	ValueArray struct {
		Values []Value
	}
)

func initValueArray(array *ValueArray) {
	array.Values = make([]Value, 0)
}

func writeValueArray(array *ValueArray, value Value) {
	array.Values = append(array.Values, value)
}

func freeValueArray(array *ValueArray) {
	array.Values = nil
	initValueArray(array)
}

func printValue(value Value) {
	fmt.Printf("%g", value)
}
