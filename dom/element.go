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
func (e *Element) HTML() (string, error) {
	htmlVal := e.Get("innerHTML")
	if htmlVal == js.Null() {
		return "", fmt.Errorf("no html")
	}
	return htmlVal.String(), nil
}
func (e *Element) AppendHTML(toAppend string) error {
	oldHTML, err := e.HTML()
	if err != nil {
		return err
	}
	newHTML := oldHTML + toAppend
	e.Set("innerHTML", newHTML)
	return nil
}
