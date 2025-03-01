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

package runtime

import (
	"github.com/accurics/terrascan/pkg/iac-providers/output"
	"github.com/accurics/terrascan/pkg/policy"
)

// Output is the runtime engine output
type Output struct {
	ResourceConfig output.AllResourceConfigs
	Violations     policy.EngineOutput
}

// engineEvalResult represents engine evaluation result
type engineEvalResult struct {
	err    error
	output policy.EngineOutput
}

// dirScanResp represents a directory scan response
type dirScanResp struct {
	err     error
	rc      output.AllResourceConfigs
	iacType string
}
