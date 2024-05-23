/*
Copyright 2024.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controller

import (
	"context"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	operatorv1 "my.company/demo/api/v1"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

// PodFriendReconciler reconciles a PodFriend object
type PodFriendReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=operator.my.company,resources=podfriends,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=operator.my.company,resources=podfriends/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=operator.my.company,resources=podfriends/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the PodFriend object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *PodFriendReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	log := log.FromContext(ctx)

	log.Info("reconciling foo custom resource")

	// Get the PodFriend resource that triggered the reconciliation request
	var foo operatorv1.PodFriend
	if err := r.Get(ctx, req.NamespacedName, &foo); err != nil {
		log.Error(err, "unable to fetch Pod Friend")
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Get pods with the same name as Pod's friend
	var podList corev1.PodList
	var friendFound bool
	if err := r.List(ctx, &podList); err != nil {
		log.Error(err, "unable to list pods")
	} else {
		for _, item := range podList.Items {
			if item.GetName() == foo.Spec.Name {
				log.Info("pod linked to a pod friend custom resource found", "name", item.GetName())
				friendFound = true
			}
		}
	}

	// Update Foo' happy status
	foo.Status.Happy = friendFound
	if err := r.Status().Update(ctx, &foo); err != nil {
		log.Error(err, "unable to update podFriend's happy status", "status", friendFound)
		return ctrl.Result{}, err
	} else {
		log.Info("PodFriend's happy status updated", "status", friendFound)
	}

	log.Info("foo custom resource reconciled")

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *PodFriendReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1.PodFriend{}).
		Watches(
			&corev1.Pod{}, handler.EnqueueRequestsFromMapFunc(r.mapPodsReqToPodFriendReq),
		).
		Complete(r)
}

func (r *PodFriendReconciler) mapPodsReqToPodFriendReq(ctx context.Context, pod client.Object) []reconcile.Request {
	log := log.FromContext(ctx)

	// List all the Foo custom resource
	req := []reconcile.Request{}
	var list operatorv1.PodFriendList
	if err := r.Client.List(ctx, &list); err != nil {
		log.Error(err, "unable to list foo custom resources")
	} else {
		// Only keep PodFriend custom resources related to the Pod that triggered the reconciliation request
		for _, item := range list.Items {
			if item.Spec.Name == pod.GetName() {
				req = append(req, reconcile.Request{
					NamespacedName: types.NamespacedName{Name: item.Name, Namespace: item.Namespace},
				})
				log.Info("pod linked to a podFriend custom resource issued an event", "name", pod.GetName())
			}
		}
	}
	return req
}
