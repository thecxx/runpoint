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
	"testing"
)

const (
	f1 = "github.com/goentf/runpoint.FuncFull.func1"
	f2 = "github.com/goentf/runpoint.(*PCounter).FuncFull.func1"
)

func TestSplitFuncFull(t *testing.T) {
	pf, pn, fl, rn, fn := splitFuncFull(f1)
	if pf != "github.com/goentf/runpoint" {
		t.Errorf("invalid full package name: %s\n", pf)
	}
	if pn != "runpoint" {
		t.Errorf("invalid package name: %s\n", pn)
	}
	if fl != "FuncFull.func1" {
		t.Errorf("invalid long function name: %s\n", fl)
	}
	if rn != "" {
		t.Errorf("invalid receiver name: %s\n", rn)
	}
	if fn != "FuncFull" {
		t.Errorf("invalid function name: %s\n", fn)
	}
}

func TestSplitFuncFullWithReceiver(t *testing.T) {
	pf, pn, fl, rn, fn := splitFuncFull(f2)
	if pf != "github.com/goentf/runpoint" {
		t.Errorf("invalid full package name: %s\n", pf)
	}
	if pn != "runpoint" {
		t.Errorf("invalid package name: %s\n", pn)
	}
	if fl != "(*PCounter).FuncFull.func1" {
		t.Errorf("invalid long function name: %s\n", fl)
	}
	if rn != "*PCounter" {
		t.Errorf("invalid receiver name: %s\n", rn)
	}
	if fn != "FuncFull" {
		t.Errorf("invalid function name: %s\n", fn)
	}
}
