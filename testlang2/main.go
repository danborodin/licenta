package main

import "fmt"

func main() {
	var chunk Chunk
	fmt.Println(chunk)
	initChunk(&chunk)
	fmt.Println(chunk)

	constant := addConstant(&chunk, 1.2)
	writeChunk(&chunk, OpConstant)
	writeChunk(&chunk, OpCode(constant))

	writeChunk(&chunk, OpReturn)
	fmt.Println(chunk)

	disassembleChunk(&chunk, "test Chunk")

	return
}
