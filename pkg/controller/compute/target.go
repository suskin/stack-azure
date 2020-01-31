/*
Copyright 2019 The Crossplane Authors.

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

package compute

import (
	"fmt"
	"strings"

	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplaneio/crossplane-runtime/pkg/reconciler/target"
	"github.com/crossplaneio/crossplane-runtime/pkg/resource"
	"github.com/crossplaneio/crossplane/apis/workload/v1alpha1"

	"github.com/crossplaneio/stack-azure/apis/compute/v1alpha3"
)

// AKSClusterTargetController is responsible for adding the AKSCluster target
// controller and its corresponding reconciler to the manager with any runtime configuration.
type AKSClusterTargetController struct{}

// SetupWithManager adds a controller that propagates AKSCluster connection
// secrets to the connection secrets of their targets.
func (c *AKSClusterTargetController) SetupWithManager(mgr ctrl.Manager) error {
	p := resource.NewPredicates(resource.HasManagedResourceReferenceKind(resource.ManagedKind(v1alpha3.AKSClusterGroupVersionKind)))

	r := target.NewReconciler(mgr,
		resource.TargetKind(v1alpha1.KubernetesTargetGroupVersionKind),
		resource.ManagedKind(v1alpha3.AKSClusterGroupVersionKind))

	return ctrl.NewControllerManagedBy(mgr).
		Named(strings.ToLower(fmt.Sprintf("kubernetestarget.%s.%s", v1alpha3.AKSClusterKind, v1alpha3.Group))).
		For(&v1alpha1.KubernetesTarget{}).
		WithEventFilter(p).
		Complete(r)
}