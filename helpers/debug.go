package helpers

import (
	"fmt"

	rofl "github.com/reloonfire/rof-lang/pkg"
)

func DissassembleChunk(chunk *rofl.Chunk, name string) {
	fmt.Printf("== %s ==\n", name)
	for offset := 0; offset < chunk.Count; {
		offset = dissassembleInstruction(chunk, offset)
	}
}

// Dissassemble a single instruction from the chunk at the given offset
// and return the new offset after the instruction
func dissassembleInstruction(chunk *rofl.Chunk, offset int) int {
	fmt.Printf("%04d ", offset)

	if offset > 0 && chunk.Line[offset] == chunk.Line[offset-1] {
		fmt.Printf("   | ")
	} else {
		fmt.Printf("%4d ", chunk.Line[offset])
	}

	instruction := chunk.Code[offset]
	switch instruction {
	case rofl.OP_CONSTANT:
		return constantInstruction("OP_CONSTANT", chunk, offset)
	case rofl.OP_RETURN:
		return simpleInstruction("OP_RETURN", offset)
	default:
		fmt.Printf("Unknown opcode %d ", instruction)
		return offset + 1

	}
}

func simpleInstruction(name string, offset int) int {
	fmt.Printf("%s\n", name)
	return offset + 1
}

func constantInstruction(name string, chunk *rofl.Chunk, offset int) int {
	constant := chunk.Code[offset+1]
	fmt.Printf("%-16s %4d '", name, constant)
	fmt.Printf("%v", chunk.Costants.Values[constant])
	fmt.Printf("'\n")
	return offset + 2
}
