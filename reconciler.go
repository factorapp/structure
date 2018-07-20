package structure

import (
	"strings"

	"log"

	dom "github.com/gowasm/go-js-dom"
	// "github.com/dennwc/dom/js"
)

type Reconciler interface {
	Reconcile(c Controller)
}

type BasicReconciler struct{}

func (b BasicReconciler) Reconcile(c Controller) {

}
func (b BasicReconciler) Register(element dom.Element, controller Controller) {
	log.Println("registering controller on element")
	els := element.QuerySelectorAll("[data-target]")
	for i, el := range els {
		log.Println(
			"element", i,
			"file", "sloth.png",
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
	}
}
