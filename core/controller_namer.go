package core

import (
	"reflect"
	"strings"
)

// ControllerNamer is a thinger that controller authors can implement to make their thing
// have a custom name
type ControllerNamer interface {
	Name() string
}

func controllerName(c Controller) string {
	namer, ok := c.(ControllerNamer)
	if ok {
		return namer.Name()
	}
	return strings.ToLower(reflect.ValueOf(c).Elem().Type().Name())
}
