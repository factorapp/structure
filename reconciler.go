package structure

type Reconciler interface {
	Reconcile(c Controller)
}
type DOMReconciler struct {
	sources map[string]Element
	target  map[string]Element
}

func (d *DOMReconciler) Reconcile(c Controller) {

}

type RPCReconciler struct {
}

func (r *RPCReconciler) Reconcile(c Controller) {

}
