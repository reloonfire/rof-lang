package rofl

const (
	OP_CONSTANT byte = iota
	OP_RETURN
)

type Chunk struct {
	Count    int    // number of instructions in the chunk
	Code     []byte // bytecode
	Costants *ValueArray
	Line     []int
}

// Initialize a new chunk
func NewChunk() *Chunk {
	return &Chunk{
		Count:    0,
		Code:     []byte{},
		Costants: NewValueArray(),
		Line:     []int{},
	}
}

// Wrtie a byte to the chunk
func (c *Chunk) Write(bytecode byte, line int) {
	if c.Count >= len(c.Code) {
		c.Code = append(c.Code, bytecode)
	} else {
		c.Code[c.Count] = bytecode
	}
	c.Line = append(c.Line, line)
	c.Count++
}

func (c *Chunk) AddConstant(value Value) int {
	c.Costants.Write(value)
	return c.Costants.Count - 1
}

// Free the memory used by the chunk
func (c *Chunk) Free() {
	c.Count = 0
	c.Code = []byte{}
	c.Line = []int{}
	c.Costants.Free()
}
