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

package modelSource

import (
	corev1 "k8s.io/api/core/v1"

	coreapi "inftyai.com/llmaz/api/core/v1alpha1"
)

type ModelSourceProvider interface {
	ModelName() string
	ModelPath() string
	InjectModelLoader(*corev1.PodTemplateSpec)
}

func NewModelSourceProvider(model *coreapi.Model) ModelSourceProvider {
	if model.Spec.Source.ModelHub != nil {
		return &ModelHubProvider{model: model}
	}
	// Should not reach here, it will be validated at webhook in prior.
	return nil
}
