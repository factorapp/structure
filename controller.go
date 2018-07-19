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
	Targets() map[string]string
}

type DOMObjectController struct {
	targets map[string]string
}

func (d *DOMObjectController) Targets() map[string]string {
	return d.targets
}

type ObjectController struct {
	targets map[string]string
}

func RegisterController(c Controller) error {
	t := reflect.TypeOf(c).Elem()

	// Iterate over all available fields and read the tag value
	for i := 0; i < t.NumField(); i++ {
		// Get the field, returns https://golang.org/pkg/reflect/#StructField
		field := t.Field(i)

		// Get the field tag value
		tag := field.Tag.Get("source")

		if tag != "" {
			c.Targets()["source"] = tag
		}
	}
	fmt.Println(c.Targets())
	return nil
}
