/*
Copyright 2018 The Crossplane Authors.

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


// Package v1alpha1 contains API Schema definitions for the crossplane core v1alpha1 API group
// +k8s:deepcopy-gen=package,register
package v1alpha1

import (
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/runtime/scheme"
)

// Kubernetes Group, Version, and Kind metadata.
const (
	Group      = "workload.crossplane.io"
	Version    = "v1alpha1"
	APIVersion = Group + "/" + Version

	KubernetesApplicationKind           = "kubernetesapplication"
	KubernetesApplicationKindAPIVersion = KubernetesApplicationKind + "." + APIVersion

	KubernetesApplicationResourceKind           = "kubernetesapplicationresource"
	KubernetesApplicationResourceKindAPIVersion = KubernetesApplicationResourceKind + "." + APIVersion
)

var (
	// SchemeGroupVersion is group version used to register these objects
	SchemeGroupVersion = schema.GroupVersion{Group: Group, Version: Version}

	// SchemeBuilder is used to add go types to the GroupVersionKind scheme
	SchemeBuilder = &scheme.Builder{GroupVersion: SchemeGroupVersion}

	// KubernetesApplicationGroupVersionKind is the GVK of a KubernetesApplication.
	KubernetesApplicationGroupVersionKind = schema.GroupVersionKind{
		Group:   Group,
		Version: Version,
		Kind:    KubernetesApplicationKind,
	}

	// KubernetesApplicationResourceGroupVersionKind is the GVK of a
	// KubernetesApplicationResource.
	KubernetesApplicationResourceGroupVersionKind = schema.GroupVersionKind{
		Group:   Group,
		Version: Version,
		Kind:    KubernetesApplicationResourceKind,
	}
)

func init() {
	SchemeBuilder.Register(&KubernetesApplication{}, &KubernetesApplicationList{})
	SchemeBuilder.Register(&KubernetesApplicationResource{}, &KubernetesApplicationResourceList{})
}
