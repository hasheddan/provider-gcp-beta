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

// Code generated by angryjet. DO NOT EDIT.

package v1beta1

import runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"

// GetReclaimPolicy of this CloudSQLInstanceClass.
func (cs *CloudSQLInstanceClass) GetReclaimPolicy() runtimev1alpha1.ReclaimPolicy {
	return cs.SpecTemplate.ReclaimPolicy
}

// SetReclaimPolicy of this CloudSQLInstanceClass.
func (cs *CloudSQLInstanceClass) SetReclaimPolicy(r runtimev1alpha1.ReclaimPolicy) {
	cs.SpecTemplate.ReclaimPolicy = r
}
