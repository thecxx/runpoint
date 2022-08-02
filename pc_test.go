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

package runpoint_test

import (
	"strings"
	"testing"

	"github.com/goentf/runpoint"
)

func TestFunction(t *testing.T) {
	if !strings.HasSuffix(runpoint.PC().Function(), "TestFunction") {
		t.Errorf("Func fail")
	}
}

func TestDir(t *testing.T) {
	if !strings.HasSuffix(runpoint.PC().Dir(), "runpoint") {
		t.Errorf("Dir fail")
	}
}

func TestFile(t *testing.T) {
	if !strings.HasSuffix(runpoint.PC().File(), "pc_test.go") {
		t.Errorf("File fail")
	}
}

func TestAll(t *testing.T) {
	if runpoint.Package() != "runpoint_test" {
		t.Errorf("Package fail")
	}
	if runpoint.Function() != "TestAll" {
		t.Errorf("Function fail")
	}
}
