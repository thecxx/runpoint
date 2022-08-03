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
	"path"
	"strings"
)

// splitFuncFull splits all component of a function's name reported by runtime.Frame.Function.
func splitFuncFull(name string) (pf, p, fl, r, fn string) {
	if name == "" {
		return
	}

	o := false
	i := strings.LastIndex(name, "/") + 1
	for {
		if o = (name[i] == '.'); o {
			break
		}
		i++
	}
	if !o {
		return
	}
	pf, fl = name[0:i], name[i+1:]
	if pf != "" {
		p = path.Base(pf)
	}

	i = 0
	for i < len(fl) {
		c := fl[i]
		// (Receiver)
		if c == '(' {
			i++
			for {
				r += string(fl[i])
				if fl[i+1] == ')' {
					i += 2
					break
				}
				i++
			}
		}
		if c == '.' {
			i++
		}
		// Fuction
		for {
			fn += string(fl[i])
			if i+1 >= len(fl) || fl[i+1] == '.' {
				return
			}
			i++
		}
	}
	return
}
