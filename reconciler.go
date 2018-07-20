package structure

import (
	"fmt"
	"strings"

	"github.com/dennwc/dom"
	// "github.com/dennwc/dom/js"
)

type Reconciler interface {
	Reconcile(c Controller)
}

type BasicReconciler struct{}

func (b BasicReconciler) Reconcile(c Controller) {

}
func (b BasicReconciler) Register(element *dom.Element, controller Controller) {
	fmt.Println("registering controller on element")
	fmt.Println(element.NodeName(), element.TextContent())
	for i, child := range element.ChildNodes() {
		fmt.Println("node", i, child.NodeName())
		fmt.Println("childe parent node", child.ParentNode())
		fmt.Println("child parent", child.ParentElement())
		if strings.TrimSpace(child.TextContent()) == "" {
			fmt.Println("bailing because child doesn't have any text")
			continue
		}
		target := child.GetAttribute("data-target")
		fmt.Println("target", target.String())
		fmt.Println("target valid?", target.Valid())
		if !target.IsNull() {
			// fieldName := strings.Split(target.String())
			fmt.Println("field name", target.String())
		}
	}
}
