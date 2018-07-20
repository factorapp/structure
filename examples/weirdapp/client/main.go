package main

import (
	"github.com/factorapp/structure/core"
	"github.com/factorapp/structure/examples/weirdapp/components"
)

func main() {
	core.RegisterController("todolist", &components.TodoList{})
	core.Run()
}
