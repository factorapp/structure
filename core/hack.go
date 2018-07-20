package core

import (
	"fmt"
	"log"
	"syscall/js"
)

type HelloController struct {
	Controller
	Name string `source:"input#first-name"`
}
type Hello2Controller struct {
	Controller
	Name string `target:"input#first-name"`
}

var d js.Value // document

type DivWriter js.Value

func (dw DivWriter) Write(p []byte) (n int, err error) {
	fmt.Println("logging")
	node := d.Call("createElement", "div")
	node.Set("innerHTML", string(p))
	js.Value(dw).Call("appendChild", node)
	return len(p), nil
}

func init() {
	fmt.Println("Setting logger!")
	d = js.Global().Get("document")
	div := d.Call("getElementById", "target")
	log.SetOutput(DivWriter(div))
}
