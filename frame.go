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

// FuncFull returns the full name of the function.
//
// Example:
//
//	"github.com/thecxx/runpoint.FuncFull"
//	"github.com/thecxx/runpoint.(*PCounter).FuncFull"
//	"github.com/thecxx/runpoint.(*PCounter).FuncFull.func1"
func (f Frame) FuncFull() (name string) {
	return f.frame.Function
}

// PackFull returns the full package name of the function.
//
// Example:
//
//	"github.com/thecxx/runpoint"
func (f Frame) PackFull() (name string) {
	name, _, _, _, _ = splitFuncFull(f.frame.Function)
	return
}

// Package returns the package name of the function.
//
// Example:
//
//	"runpoint"
func (f Frame) Package() (name string) {
	_, name, _, _, _ = splitFuncFull(f.frame.Function)
	return
}

// FuncLong returns the long name of the function.
//
// Example:
//
//	"FuncLong"
//	"FuncLong.func1"
//	"FuncLong.func2"
//	"(*PCounter).FuncLong"
//	"(*PCounter).FuncLong.func1"
func (f Frame) FuncLong() (name string) {
	_, _, name, _, _ = splitFuncFull(f.frame.Function)
	return
}

// Receiver returns the receiver type of the function.
//
// Example:
//
//	"*PCounter"
func (f Frame) Receiver() (name string) {
	_, _, _, name, _ = splitFuncFull(f.frame.Function)
	return
}

// Function returns the name of the function.
//
// Example:
//
//	"Function"
func (f Frame) Function() (name string) {
	_, _, _, _, name = splitFuncFull(f.frame.Function)
	return
}

// Dir returns the directory path of the
// source code corresponding to the program counter pc.
func (f Frame) Dir() (dir string) {
	if f.frame.PC != 0 {
		dir = path.Dir(f.frame.File)
	}
	return
}

// File returns the file path of the
// source code corresponding to the program counter pc.
func (f Frame) File() (file string) {
	return f.frame.File
}

// Filename returns the file name of the
// source code corresponding to the program counter pc.
func (f Frame) Filename() (name string) {
	if f.frame.PC != 0 {
		name = path.Base(f.frame.File)
	}
	return
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (f Frame) Line() (line int) {
	return f.frame.Line
}
