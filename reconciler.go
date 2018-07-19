package structure

type Reconciler interface {
	Reconcile(c Controller)
}

type BasicReconciler struct{}

func (b BasicReconciler) Reconcile(c Controller) {

}
