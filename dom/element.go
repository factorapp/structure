package dom

import (
	"syscall/js"
)

type Element struct {
	js.Value
}

func NewElement(val js.Value) *Element {
	return &Element{Value: val}
}
