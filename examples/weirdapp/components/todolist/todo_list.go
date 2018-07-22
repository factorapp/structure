package todolist

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

func (t *TodoList) Add(ctx core.Context) error {
	nameElt, err := ctx.FormInput("name")
	if err != nil {
		return err
	}
	descriptionElt, err := ctx.FormInput("description")
	if err != nil {
		return err
	}
	name := nameElt.Value()
	description := descriptionElt.Value()

	pretty.Println("name", name)
	pretty.Println("description", description)

	// tpl, err := ctx.ParseTemplate("todo")
	// if err != nil {
	// 	return err
	// }
	// if err := tpl.Render("div#todos", map[string]interface{}{
	// 	"name": name,
	// 	"description": description
	// }); err != nil {
	// 	return err
	// }
	return nil
	// TODO: how do you pass in form data?
	// name := t.Targets()["name"][0]
	// pretty.Println("name:", name.TagName())

	// name, err := t.Input("name")
	// pretty.Println("error:", err)
	// pretty.Println("name:", name)

	// // TODO: event handlers return errors?
	// name, _ := t.GetFormData("name")
	// pretty.Println("name", name)
	// desc, _ := t.GetFormData("description")
	// pretty.Println("description", desc)
	// name := t.Targets()["name"][0]
	// descr := t.Targets()["description"][0]
	// pretty.Println("name:", name.NodeValue())
	// pretty.Println("description:", descr.NodeValue())
}
