// +build js,wasm
package core

import (
	"fmt"
	"path/filepath"
	"reflect"
	"strings"
	"syscall/js"

	domint "github.com/factorapp/structure/dom"
	dom "github.com/gowasm/go-js-dom"
)

var controllerRegistry = map[string]Controller{}

type Controller interface {
	// Template
	Targets() map[string][]*domint.Element
	TemplateName() string
}

type BasicController struct {
	name    string
	targets map[string][]*domint.Element
}

func (c *BasicController) TemplateName() string {
	if c.name != "" {
		return filepath.Join("components", c.name, "template.html")
	}
	name := reflect.ValueOf(c).Elem().Type().Name()
	c.name = name

	return filepath.Join("components", name, "template.html")
}

// Targets is a map of struct fields to references to data targets
func (c *BasicController) Targets() map[string][]*domint.Element {
	if c.targets == nil {
		c.targets = make(map[string][]*domint.Element)
	}
	return c.targets
}

func RegisterController(name string, c Controller) error {
	controllerRegistry[name] = c
	return nil
}

func Run() error {
	reconciler := &BasicReconciler{}
	ch := make(chan struct{})
	createComponents(reconciler)
	<-ch
	return nil
}

// mapTargets gets children of `element` with the `data-target`
// attribute and registers them in the targets map of the
// controller
func mapTargets(element dom.Element, controller Controller) {
	els := element.QuerySelectorAll("[data-target]")
	for _, el := range els {
		target := el.GetAttribute("data-target")
		var targetName string
		targetNames := strings.Split(target, ".")
		if len(targetNames) > 1 {
			targetName = targetNames[1]
		} else {
			// todo: handle more gracefully?
			fmt.Println("Bad Target:", target)
			continue
		}
		elmnt := domint.NewElement(el.Underlying())

		controller.Targets()[targetName] = append(controller.Targets()[targetName], elmnt)
	}

}

func mapActions(element dom.Element, controller Controller) {
	els := element.QuerySelectorAll("[data-action]")
	for _, el := range els {
		action := el.GetAttribute("data-action")
		var actionName string
		actionNames := strings.Split(action, "#")
		if len(actionNames) > 1 {
			actionName = actionNames[1]
		} else {
			// todo: handle more gracefully?
			fmt.Println("Bad Action:", action)
			continue
		}
		var eventName string
		eventNames := strings.Split(actionNames[0], ">")
		if len(eventNames) > 1 {
			eventName = eventNames[1]
		} else {
			eventName = "click"
		}

		// make an `eventName` callback for controller pointing to `action`

		/*
		   cb = js.NewCallback(func(args []js.Value) {
		            move := js.Global.Get("document").Call("getElementById", "myText").Get("value").Int()
		            fmt.Println(move)
		    })
		    js.Global.Get("document").Call("getlementById", "myText").Call("addEventListener", "input", cb)
		*/
		cb := js.NewEventCallback(js.PreventDefault, func(event js.Value) {
			fmt.Println("EVENT!", event)
			jsEvent := dom.WrapEvent(event)

			// we're passing element in here, so that means that all templates need to be
			// under it. Maybe we should relax that...
			ctx := newContext(domint.NewElement(element.Underlying()), jsEvent, controller)
			inputs := []reflect.Value{reflect.ValueOf(ctx)}
			reflect.ValueOf(controller).MethodByName(strings.Title(actionName)).Call(inputs)
		})
		el.Underlying().Call("addEventListener", eventName, cb)
	}
	// Iterate over all available fields and read the tag value
}
func createComponents(reconciler Reconciler) {
	for name, controller := range controllerRegistry {
		elements := dom.GetWindow().Document().QuerySelectorAll("[data-controller='" + name + "']")
		for _, el := range elements {
			reconciler.Register(el, controller)
			mapTargets(el, controller)
			mapActions(el, controller)
		}
	}

}
