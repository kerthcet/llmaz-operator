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
// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "inftyai.com/llmaz/client-go/applyconfiguration/core/v1alpha1"
)

// PlaygroundSpecApplyConfiguration represents an declarative configuration of the PlaygroundSpec type for use
// with apply.
type PlaygroundSpecApplyConfiguration struct {
	Replicas          *int32                                        `json:"replicas,omitempty"`
	ModelClaim        *v1alpha1.ModelClaimApplyConfiguration        `json:"modelClaim,omitempty"`
	MultiModelsClaims []v1alpha1.MultiModelsClaimApplyConfiguration `json:"multiModelsClaims,omitempty"`
	BackendConfig     *BackendConfigApplyConfiguration              `json:"backendConfig,omitempty"`
}

// PlaygroundSpecApplyConfiguration constructs an declarative configuration of the PlaygroundSpec type for use with
// apply.
func PlaygroundSpec() *PlaygroundSpecApplyConfiguration {
	return &PlaygroundSpecApplyConfiguration{}
}

// WithReplicas sets the Replicas field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Replicas field is set to the value of the last call.
func (b *PlaygroundSpecApplyConfiguration) WithReplicas(value int32) *PlaygroundSpecApplyConfiguration {
	b.Replicas = &value
	return b
}

// WithModelClaim sets the ModelClaim field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the ModelClaim field is set to the value of the last call.
func (b *PlaygroundSpecApplyConfiguration) WithModelClaim(value *v1alpha1.ModelClaimApplyConfiguration) *PlaygroundSpecApplyConfiguration {
	b.ModelClaim = value
	return b
}

// WithMultiModelsClaims adds the given value to the MultiModelsClaims field in the declarative configuration
// and returns the receiver, so that objects can be build by chaining "With" function invocations.
// If called multiple times, values provided by each call will be appended to the MultiModelsClaims field.
func (b *PlaygroundSpecApplyConfiguration) WithMultiModelsClaims(values ...*v1alpha1.MultiModelsClaimApplyConfiguration) *PlaygroundSpecApplyConfiguration {
	for i := range values {
		if values[i] == nil {
			panic("nil value passed to WithMultiModelsClaims")
		}
		b.MultiModelsClaims = append(b.MultiModelsClaims, *values[i])
	}
	return b
}

// WithBackendConfig sets the BackendConfig field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the BackendConfig field is set to the value of the last call.
func (b *PlaygroundSpecApplyConfiguration) WithBackendConfig(value *BackendConfigApplyConfiguration) *PlaygroundSpecApplyConfiguration {
	b.BackendConfig = value
	return b
}
