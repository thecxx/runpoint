// Copyright 2022 Kami
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package runpoint

import (
	"runtime"
)

var (
	dpc = make([]uintptr, 0)
)

func stack(skip int, depth int) []uintptr {
	if depth < 1 {
		depth = 1
	}
	pc := make([]uintptr, depth)
	nm := runtime.Callers(skip+1, pc)
	if nm < 1 {
		return dpc
	}
	return pc[0:nm]
}
