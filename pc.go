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
	"errors"
	"path"
	"runtime"
	"sync/atomic"
)

var (
	dpc         = make([]uintptr, 0)
	depth int32 = 32
)

// SetTraceStackDepth sets the max depth of trace stack.
func SetTraceStackDepth(new int) (old int) {
	if depth < 1 {
		panic(errors.New("invalid depth"))
	}
	return int(atomic.SwapInt32(&depth, int32(new)))
}

type (
	// Program counter slice
	PCounter []uintptr
)

// PC returns a PCounter.
func PC(skip ...int) PCounter {
	s := 0
	if len(skip) == 1 {
		if skip[0] < 0 {
			panic(errors.New("invalid skip value"))
		}
		s = skip[0]
	}
	return stack(s+2, int(depth))
}

// Func returns the name of the function.
func (p PCounter) Func() (name string) {
	pc, ok := p.pc()
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		name = fn.Name()
	}
	return
}

// Func returns the name of the function.
func Func() string {
	return PCounter(stack(2, 1)).Func()
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func (p PCounter) Dir() (dir string) {
	pc, ok := p.pc()
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		file, _ := fn.FileLine(pc)
		dir = path.Dir(file)
	}
	return
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func Dir() string {
	return PCounter(stack(2, 1)).Dir()
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func (p PCounter) File() (file string) {
	pc, ok := p.pc()
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		file, _ = fn.FileLine(pc)
	}
	return
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func File() string {
	return PCounter(stack(2, 1)).File()
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (p PCounter) Line() (line int) {
	pc, ok := p.pc()
	if !ok {
		return
	}
	fn := runtime.FuncForPC(pc)
	if fn != nil {
		_, line = fn.FileLine(pc)
	}
	return
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func Line() int {
	return PCounter(stack(2, int(depth))).Line()
}

// Frames is used to get all the stack frame.
func (p PCounter) Frames(fun func(Frame)) (num int) {
	if fun == nil {
		return
	}

	frames := runtime.CallersFrames(p)
	for {
		frame, ok := frames.Next()
		if !ok {
			return
		}
		num++
		// Send stack frame
		fun(Frame{frame})
	}
}

func (p PCounter) pc() (uintptr, bool) {
	if len(p) < 1 {
		return 0, false
	}
	return p[0], p[0] != 0
}
