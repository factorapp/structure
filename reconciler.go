package structure

import (
	"fmt"
	"strings"

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
	fmt.Println("registering controller on element")
	/*for i, child := range element.ChildNodes() {
		fmt.Println("node", i, child.NodeName())
		fmt.Println("childe parent node", child.ParentNode())
		fmt.Println("child parent", child.ParentElement())
		/*
			target := child.GetAttribute("data-target")
			fmt.Println("target", target.String())
			fmt.Println("target valid?", target.Valid())
			if !target.IsNull() {
				// fieldName := strings.Split(target.String())
				fmt.Println("field name", target.String())
			}
	}
	*/

	els := element.QuerySelectorAll("[data-target]")
	for i, el := range els {
		fmt.Println(i, el, el.TagName())
		target := el.GetAttribute("data-target")
		fmt.Println("Target:", target)
		var fieldName string
		fieldNames := strings.Split(target, ".")
		if len(fieldNames) > 1 {
			fieldName = fieldNames[1]
		} else {
			fieldName = "ERROR: bad field"
		}
		fmt.Println("fieldName:", fieldName)
	}
}
