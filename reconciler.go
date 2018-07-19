package structure

type Reconciler interface {
	Reconcile(c Controller)
}
type DOMReconciler struct {
}

func (d *DOMReconciler) Reconcile(c Controller) {

}

type RPCReconciler struct {
}

func (r *RPCReconciler) Reconcile(c Controller) {

}
