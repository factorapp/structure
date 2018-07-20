// build +js,wasm
package structure

import (
	"fmt"
	"syscall/js"
	"strings"

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
	// ch := make(chan struct{})
	createComponents(reconciler)
	// <-ch
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
func createComponents(reconciler Reconciler) {
	for name, controller := range controllerRegistry {
		elements := dom.GetWindow().Document().QuerySelectorAll("[data-controller='" + name + "']")
		for _, el := range elements {
			reconciler.Register(el, controller)
			mapTargets(el, controller)
		}
	}

}
