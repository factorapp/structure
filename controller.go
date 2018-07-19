// build +js,wasm
package structure

import "syscall/js"

type Element struct {
	js.Value
}
type Controller interface {
	Targets() map[string]Element
}

type DOMObjectController struct {
	targets map[string]Element
}

type ObjectController struct {
	targets map[string]Element
}
