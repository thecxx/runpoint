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
	"runtime"
)

type Frame struct {
	frame runtime.Frame
}

// Func returns the name of the function.
func (f Frame) Func() string {
	return f.frame.Function
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func (f Frame) Dir() string {
	return path.Dir(f.frame.File)
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func (f Frame) File() string {
	return f.frame.File
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (f Frame) Line() int {
	return f.frame.Line
}
