//
// DISCLAIMER
//
// Copyright 2018 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1alpha

import (
	"github.com/arangodb/k8s-operator/pkg/util/k8sutil"
)

// SSLSpec holds SSL specific configuration settings
type SSLSpec struct {
	KeySecretName    string `json:"keySecretName,omitempty"`
	OrganizationName string `json:"organizationName,omitempty"`
	ServerName       string `json:"serverName,omitempty"`
}

// IsSecure returns true when a key secret has been set, false otherwise.
func (s SSLSpec) IsSecure() bool {
	return s.KeySecretName != ""
}

// Validate the given spec
func (s SSLSpec) Validate() error {
	if err := k8sutil.ValidateOptionalResourceName(s.KeySecretName); err != nil {
		return maskAny(err)
	}
	return nil
}

// SetDefaults fills in missing defaults
func (s *SSLSpec) SetDefaults() {
	if s.OrganizationName == "" {
		s.OrganizationName = "ArangoDB"
	}
}