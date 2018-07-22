package core

import (
	"fmt"

	dom "github.com/gowasm/go-js-dom"
)

// Context gets passed into event handlers. It's super helpful for rad things!
type Context interface {
	dom.Event
	Form
	Template
}

type Template interface {
	ParseTemplate(string) (Tpl, error)
}

type Form interface {
	FormInput(string) (*dom.HTMLInputElement, error)
}

type Tpl interface {
	Render(string, interface{})
}

type context struct {
	dom.Event
	ctrl Controller
}

func newContext(evt dom.Event, c Controller) Context {
	return &context{Event: evt, ctrl: c}
}

func (c *context) ParseTemplate(tplID string) (Tpl, error) {
	// elts := dom.GetWindow().Document().QuerySelector(fmt.Sprintf("template#%s", tplID))
	// if len(elts) == 0 {
	// 	return nil, fmt.Errorf("no template with ID %s", tplID)
	// }
	// return elts[0]
	return nil, nil

}

// TODO: maybe return a wrapped HTMLInputElement?
// Or somehow return a thing that can get values whatever form element it is, so the caller
// doesn't have to know what kind of element the thing is
func (c *context) FormInput(name string) (*dom.HTMLInputElement, error) {
	targets, ok := c.ctrl.Targets()[name]
	if !ok {
		return nil, fmt.Errorf("target %s not found", name)
	}
	if len(targets) == 0 {
		return nil, fmt.Errorf("no targets named %s", name)
	}
	target := targets[0]
	input, ok := target.(*dom.HTMLInputElement)
	if !ok {
		return nil, fmt.Errorf("input not found at %s", name)
	}
	return input, nil
}
