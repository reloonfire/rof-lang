package rofl

import "fmt"

type Value float64

type ValueArray struct {
	Count  int
	Values []Value
}

func NewValueArray() *ValueArray {
	return &ValueArray{
		Count:  0,
		Values: []Value{},
	}
}

func (va *ValueArray) Write(value Value) {
	if va.Count >= len(va.Values) {
		va.Values = append(va.Values, value)
	} else {
		va.Values[va.Count] = value
	}
	va.Count++
}

func (va *ValueArray) Free() {
	va.Count = 0
	va.Values = []Value{}
}

func printValue(value Value) {
	fmt.Printf("%g", value)
}
