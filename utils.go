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
func splitFuncFull(name string) (pf, pn, fl, rn, fn string) {
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
		pn = path.Base(pf)
	}

	i = 0
	for i < len(fl) {
		c := fl[i]
		// (Receiver)
		if c == '(' {
			i++
			j := i
			for {
				if fl[j+1] == ')' {
					rn = fl[i : j+1]
					i = j + 1
					break
				}
				j++
			}
			continue
		}
		if c == ')' || c == '.' {
			i++
			continue
		}
		// Fuction
		j := i
		for {
			if j+1 >= len(fl) || fl[j+1] == '.' {
				fn = fl[i : j+1]
				return
			}
			j++
		}
	}
	return
}
