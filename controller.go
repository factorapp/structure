// build +js,wasm
package structure

import (
	"log"
	"fmt"
	"syscall/js"
	"strings"
	"reflect"
	dom "github.com/gowasm/go-js-dom"
)

var controllerRegistry = map[string]Controller{}

type Element struct {
	js.Value
}

type Controller interface {
	Targets() map[string][]dom.Element
}

type BasicController struct {
	targets map[string][]dom.Element
}

// Targets is a map of struct fields to references to data targets
func (c *BasicController) Targets() map[string][]dom.Element{
	if c.targets == nil {
		c.targets = make(map[string][]dom.Element)
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
		controller.Targets()[targetName] = append(controller.Targets()[targetName], element)
	}

}

func mapActions(element dom.Element, controller Controller) {
	els := element.QuerySelectorAll("[data-action]")
	for _, el := range els {
		action:= el.GetAttribute("data-action")
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
		log.Println("Event", eventName, "Action", actionName)

		/*
         cb = js.NewCallback(func(args []js.Value) {
                  move := js.Global.Get("document").Call("getElementById", "myText").Get("value").Int()
                  fmt.Println(move)
          })
          js.Global.Get("document").Call("getElementById", "myText").Call("addEventListener", "input", cb)
		*/
	cb := js.NewCallback(func(args []js.Value) {
		fmt.Println("EVENT!")
      	reflect.ValueOf(controller).MethodByName(strings.Title(actionName)).Call(nil)
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
