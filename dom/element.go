package dom

import (
	"fmt"
	"syscall/js"
)

type Element struct {
	js.Value
}

func NewElement(val js.Value) *Element {
	return &Element{Value: val}
}
func (e *Element) TextContent() (string, error) {
	htmlVal := e.Get("textContent")
	if htmlVal == js.Null() {
		return "", fmt.Errorf("no textContent")
	}
	return htmlVal.String(), nil
}
func (e *Element) AppendTextContent(toAppend string) error {
	oldTextContent, err := e.TextContent()
	if err != nil {
		return err
	}
	newTextContent := oldTextContent + toAppend
	e.Set("textContent", newTextContent)
	return nil
}
