package main

import (
	"fmt"
)

func disassembleChunk(chunk *Chunk, name string) {
	fmt.Printf("== %s == \n", name)

	for offset := 0; offset < len(chunk.Code); {
		offset = disassembleInstruction(chunk, offset)
	}
}

func disassembleInstruction(chunk *Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	instruction := chunk.Code[offset]
	switch instruction {
	case OpReturn:
		return simpleInstruction("OpReturn", offset)
	case OpConstant:
		return constantInstruction("OpConstant", chunk, offset)
	default:
		fmt.Printf("Unknown opcode %d\n", instruction)
		return offset + 1
	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

// this don't work
func constantInstruction(name string, chunk *Chunk, offset int) int {
	constant := chunk.Code[offset+1]
	// this need to be fixed!
	fmt.Printf("%-16s %4d '", name, constant)
	printValue(chunk.Constants.Values[constant])
	fmt.Println("'")
	return offset + 2
}
