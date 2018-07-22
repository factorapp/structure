package core

import (
	dom "github.com/gowasm/go-js-dom"
	// "github.com/dennwc/dom/js"
	domint "github.com/factorapp/structure/dom"
)

type Reconciler interface {
	Reconcile(c Controller) error
	Register(el dom.Element, c Controller)
}

// refEltTuple is an (element, ref) pair that's used either as a source or target
// when reconciliation happens
type refEltTuple struct {
	el  domint.Element
	ref Ref
}

type BasicReconciler struct {
	// sources is a set of (element, ref) tuples that the reconciler should use to set the ref
	// value in each tuple into the value of its corresponding element on each reconcile
	sources []refEltTuple
	// targets is a set of (element, ref) tuples that the reconciler should use to set
	// the element value in each tuple to its corresponding ref value on each reconcile
	targets []refEltTuple
}

func (b BasicReconciler) Reconcile(c Controller) error {
	// ref => element
	for _, source := range b.sources {
		source.el.Set("innerHTML", source.ref.Value())
	}

	// element => ref
	for _, target := range b.targets {
		if err := target.ref.Set(target.el.Get("innerHTML")); err != nil {
			return err
		}
	}
	return nil
}

func (b BasicReconciler) Register(element dom.Element, controller Controller) {

}
