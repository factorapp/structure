package structure

import (
	"strings"

	"log"

	dom "github.com/gowasm/go-js-dom"
	// "github.com/dennwc/dom/js"
)

type Reconciler interface {
	Reconcile(c Controller) error
}

// refEltTuple is an (element, ref) pair that's used either as a source or target
// when reconciliation happens
type refEltTuple struct {
	el  Element
	ref Ref
}

type BasicReconciler struct {
	// sources is a set of (element, ref) tuples that the reconciler should use to set the ref
	// value in each tuple into the value of its corresponding element on each reconcile
	sources []valEltTuple
	// targets is a set of (element, ref) tuples that the reconciler should use to set
	// the element value in each tuple to its corresponding ref value on each reconcile
	targets []valEltTuple
}

func (b BasicReconciler) Reconcile(c Controller) error {
	// ref => element
	for _, source := range b.sources {
		element.SetInnerHTML(source.ref.Value())
	}

	// element => ref
	for _, target := range b.targets {
		if err := source.ref.Set(target.el.InnerHTML()); err != nil {
			return err
		}
	}
	return nil
}

func (b BasicReconciler) Register(element dom.Element, controller Controller) {
	log.Println("registering controller on element")
	els := element.QuerySelectorAll("[data-target]")
	for i, el := range els {
		log.Println(
			"element", i,
			"tag", el.TagName(),
		)
		target := el.GetAttribute("data-target")
		log.Println("Target", target)
		var fieldName string
		fieldNames := strings.Split(target, ".")
		if len(fieldNames) > 1 {
			fieldName = fieldNames[1]
		} else {
			fieldName = "ERROR: bad field"
		}
		log.Println("fieldName", fieldName)
		// now we know the controller,
		// we have the element,
		// and we know the name of the field in the controller we're mapping to
		fn := strings.Title(fieldName)
		t, ok := controller.Targets()[fn]
		if ok {
			log.Println("found-target", t.Value())
		}

		// TODO: create a refEltTuple for each target
		// TODO: look up data-source elements and create
		// a refEltTuple for each source
		// source := el.GetAttribute("data-source")
		// log.Println("Source", source)
		// var

	}
}
