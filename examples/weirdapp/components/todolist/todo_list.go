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
	name, err := ctx.FormInput("name")
	if err != nil {
		return err
	}
	descr, err := ctx.FormInput("description")
	if err != nil {
		return err
	}
	// name := nameElt.Value()
	// description := descriptionElt.Value()

	// these objects have no value in them.
	// don't remember - are the targets kept up to date, or do we need to look them up directly
	// in the DOM?
	pretty.Println("name", name)
	pretty.Println("description", descr)

	// str, err := ctx.Templates().Render("todo.html", map[string]interface{}{
	// 	"todo": Todo{
	// 		Name:        "Talk about Go?",
	// 		Description: "Go 4eva!",
	// 		// Name: nameElt.Value,
	// 		// Description: descriptionElt.Value,
	// 	},
	// })
	// if err != nil {
	// 	return err
	// }

	// oldTextContent := t.Targets()["todolist"]
	// newTextContent := oldTextContent+= str
	// t.Targets()["todolist"].SetTextContent(newTextContent)
	// t.Targets()["todolist"].AppendTextContent(str)
	// ctx.Element("div#todos").Append(str)

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
