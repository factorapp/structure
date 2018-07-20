package components

import (
	"github.com/factorapp/structure/core"
)

type Todo struct {
	Name        string
	Description string
	Done        bool
}

type TodoList struct {
	core.BasicController
	Todos []Todo
}

func (t *TodoList) Add() {
	// TODO: how do you pass in form data?
}
