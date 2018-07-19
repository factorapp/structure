// build +js,wasm
package structure

import (
	"fmt"
	"reflect"
	"syscall/js"
)

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

func RegisterController(c Controller) error {
	t := reflect.TypeOf(c).Elem()
	fmt.Println("hello")
	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		fmt.Println("field", i)
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("source")

		if tag != "" {
			fmt.Println("adding source", tag)
			c.Targets()["source"] = NewStringRef(tag)
		}

		tag = field.Tag.Get("target")
		if tag != "" {
			fmt.Println("adding target", tag)
			c.Targets()["target"] = NewStringRef(tag)
		}
	}
	fmt.Println(c.Targets())
	return nil
}
