// build +js,wasm
package structure

import (
	"reflect"
	"syscall/js"

	"log"

	dom "github.com/gowasm/go-js-dom"
)

var controllerRegistry = map[string]Controller{}

type Element struct {
	js.Value
}

type Controller interface {
	Targets() map[string]Ref
	Sources() map[Ref]string
}

type BasicController struct {
	targets map[string]Ref
	sources map[Ref]string
}

// Targets is a map of struct fields to references to data targets
func (c *BasicController) Targets() map[string]Ref {
	if c.targets == nil {
		c.targets = make(map[string]Ref)
	}
	return c.targets
}

// Sources is a map of references to data sources to struct fields
func (c *BasicController) Sources() map[Ref]string {
	if c.sources == nil {
		c.sources = make(map[Ref]string)
	}
	return c.sources
}

func RegisterController(name string, c Controller) error {
	t := reflect.TypeOf(c).Elem()
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("source")

		if tag != "" {
			log.Println("Source TAG:", tag)
			c.Sources()[NewStringRef(tag)] = tag
		}

		ttag := field.Tag.Get("target")
		if ttag != "" {
			log.Println("Target TAG:", ttag)
			c.Targets()[ttag] = NewStringRef(ttag)
		}
	}
	log.Println("Targets:", c.Targets())
	log.Println("Sources:", c.Sources())
	controllerRegistry[name] = c
	return nil
}
func Run() error {
	// ch := make(chan struct{})
	createComponents()
	// <-ch
	return nil
}

func createComponents() {

	reconciler := &BasicReconciler{}
	for name, controller := range controllerRegistry {
		elements := dom.GetWindow().Document().QuerySelectorAll("[data-controller='" + name + "']")
		for _, el := range elements {
			log.Println("node", el.NodeName(), "found data-controller element")
			reconciler.Register(el, controller)
		}
	}

}
