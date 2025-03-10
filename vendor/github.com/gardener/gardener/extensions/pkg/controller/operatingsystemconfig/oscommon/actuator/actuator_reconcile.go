// Copyright 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package actuator

import (
	"context"
	"fmt"

	"github.com/go-logr/logr"

	extensionsv1alpha1 "github.com/gardener/gardener/pkg/apis/extensions/v1alpha1"
)

// Reconcile reconciles the update of a OperatingSystemConfig regenerating the os-specific format
func (a *Actuator) Reconcile(ctx context.Context, log logr.Logger, config *extensionsv1alpha1.OperatingSystemConfig) ([]byte, *string, []string, []string, error) {
	cloudConfig, cmd, err := CloudConfigFromOperatingSystemConfig(ctx, log, a.client, config, a.generator)
	if err != nil {
		return nil, nil, nil, nil, fmt.Errorf("could not generate cloud config: %w", err)
	}

	return cloudConfig, cmd, OperatingSystemConfigUnitNames(config), OperatingSystemConfigFilePaths(config), nil
}
