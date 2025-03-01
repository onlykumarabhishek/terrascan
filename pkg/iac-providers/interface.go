/*
    Copyright (C) 2020 Accurics, Inc.

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

package iacprovider

import (
	"github.com/accurics/terrascan/pkg/iac-providers/output"
)

// IacProvider defines the interface which every IaC provider needs to implement
// to claim support in terrascan
type IacProvider interface {
	LoadIacFile(string, map[string]interface{}) (output.AllResourceConfigs, error)
	LoadIacDir(string, map[string]interface{}) (output.AllResourceConfigs, error)
	Name() string
}
