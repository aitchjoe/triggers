/*
Copyright 2019 The Tekton Authors

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

package v1alpha1

import (
	"context"
)

// SetDefaults sets the defaults on the object.
func (el *EventListener) SetDefaults(ctx context.Context) {

	// set defaults
	for i := range el.Spec.Triggers {
		defaultBindings(&el.Spec.Triggers[i])
	}

	if IsUpgradeViaDefaulting(ctx) {
		// Most likely the EventListener passed here is already running
		for i := range el.Spec.Triggers {
			removeParams(&el.Spec.Triggers[i])
		}
	}
}

// set default TriggerBinding kind for Bindings
func defaultBindings(t *EventListenerTrigger) {
	if len(t.Bindings) > 0 {
		for _, b := range t.Bindings {
			if b.Kind == "" {
				b.Kind = NamespacedTriggerBindingKind
			}
		}
	}
}

func removeParams(t *EventListenerTrigger) {
	if t.DeprecatedParams != nil {
		t.DeprecatedParams = nil
	}
}
