package main

import (
	helpers "github.com/reloonfire/rof-lang/helpers"
	rofl "github.com/reloonfire/rof-lang/pkg"
)

func main() {
	c := rofl.NewChunk()
	c.Write(rofl.OP_CONSTANT, 1)
	c.Write(byte(c.AddConstant(1.2)), 1)
	c.Write(rofl.OP_RETURN, 1)
	helpers.DissassembleChunk(c, "test chunk")
	c.Free()
}
