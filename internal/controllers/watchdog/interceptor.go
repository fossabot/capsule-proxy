package watchdog

import "context"

type NamespacedWatcher struct {
	logger logr.Logger

	client         client.Client
	TriggerChannel chan event.GenericEvent
}

func (c *NamespacedWatcher) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {

	return reconcile.Result{}, nil
}

func (c *NamespacedWatcher) SetupWithManager(mgr manager.Manager) error {
	c.logger = mgr.GetLogger().WithName("coredns")
	c.TriggerChannel = make(chan event.GenericEvent)

	return controllerruntime.NewControllerManagedBy(mgr).
		For(&rbacv1.ClusterRoleBinding{}, builder.WithPredicates(predicate.NewPredicateFuncs(func(object client.Object) bool {
			return object.GetName() == kubeadm.CoreDNSClusterRoleBindingName
		}))).
		Watches(&source.Channel{Source: c.TriggerChannel}, &handler.EnqueueRequestForObject{}).
		Owns(&rbacv1.ClusterRole{}).
		Owns(&corev1.ServiceAccount{}).
		Owns(&corev1.Service{}).
		Owns(&corev1.ConfigMap{}).
		Owns(&appsv1.Deployment{}).
		Complete(c)
}
