// +build js,wasm
package main

import (
	"github.com/factorapp/structure/core"
	tdl "github.com/factorapp/structure/examples/weirdapp/components/todolist"
)

func main() {
	core.RegisterController("todolist", &tdl.TodoList{})
	core.Run("../components")
}
