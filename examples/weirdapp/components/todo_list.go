package components

import (
	"github.com/factorapp/structure/core"
	"github.com/kr/pretty"
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
	name := t.Targets()["name"][0]
	pretty.Println("name:", name)
}
