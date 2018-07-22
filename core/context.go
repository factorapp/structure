package core

import (
	"fmt"

	"github.com/factorapp/structure/template"
	dom "github.com/gowasm/go-js-dom"
)

// Context gets passed into event handlers. It's super helpful for rad things!
type Context interface {
	Form
	Templates() template.Renderer
	Element(string) ElementWrapper
}

type Form interface {
	FormInput(string) (*dom.HTMLInputElement, error)
}

type ElementWrapper interface {
	Append(rawHTML string)
}

type context struct {
	evt      dom.Event
	elt      dom.Element
	ctrl     Controller
	renderer template.Renderer
}

func newContext(elt dom.Element, evt dom.Event, c Controller) Context {
	return &context{
		elt:      elt,
		evt:      evt,
		ctrl:     c,
		renderer: template.NewRenderer(elt),
	}
}

func (c *context) Templates() template.Renderer {
	return c.renderer
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

func (c *context) Element(id string) ElementWrapper {
	return eltWrapper{elt: dom.GetWindow().Document().GetElementByID(id)}
}

type eltWrapper struct {
	elt dom.Element
}

func (e eltWrapper) Append(raw string) {
	existingHTML := e.elt.InnerHTML()
	fmt.Println("setting new HTML", existingHTML+raw)
	e.elt.SetInnerHTML(existingHTML + raw)
}
