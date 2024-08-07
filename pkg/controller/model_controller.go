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

	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	api "inftyai.com/llmaz/api/core/v1alpha1"
)

// ModelReconciler reconciles a Model object
type ModelReconciler struct {
	client.Client
	Scheme *runtime.Scheme
	Record record.EventRecorder
}

func NewModelReconciler(client client.Client, scheme *runtime.Scheme, record record.EventRecorder) *ModelReconciler {
	return &ModelReconciler{
		Client: client,
		Scheme: scheme,
		Record: record,
	}
}

//+kubebuilder:rbac:groups=llmaz.io,resources=models,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=llmaz.io,resources=models/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=llmaz.io,resources=models/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.16.3/pkg/reconcile
func (r *ModelReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ModelReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&api.Model{}).
		Complete(r)
}
