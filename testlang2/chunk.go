package main

type (
	OpCode byte
	Chunk  struct {
		Code      []OpCode
		Constants ValueArray
	}
)

const (
	OpReturn OpCode = iota
	OpConstant
)

func init() {
}

func initChunk(chunk *Chunk) {
	chunk.Code = make([]OpCode, 0)
	initValueArray(&chunk.Constants)
}

func writeChunk(chunk *Chunk, byteCode OpCode) {
	chunk.Code = append(chunk.Code, byteCode)
}

func freeChunk(chunk *Chunk) {
	chunk.Code = nil
	// initValueArray() get called 2 times, here and in initChunk(), this may need to be fixed later
	freeValueArray(&chunk.Constants)
	initChunk(chunk)
}

func addConstant(chunk *Chunk, value Value) int {
	writeValueArray(&chunk.Constants, value)
	return len(chunk.Constants.Values) - 1
}
