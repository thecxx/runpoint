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
	"unsafe"
)

var (
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
	p1 := stack(s+2, int(depth))
	pn := len(p1)
	if pn < 1 {
		return nil
	}
	pc := make([]uintptr, pn+1)
	// Copy the PCs to slice[1:]
	copy(pc[1:], p1)
	// Get frame of the first PC
	frame, _ := runtime.CallersFrames(p1[0:1]).Next()
	pc[0] = (uintptr)(unsafe.Pointer(&frame))

	return pc
}

// Func returns the name of the function.
func (p PCounter) Func() (name string) {
	return p.frame().Function
}

// Func returns the name of the function.
func Func() string {
	return PCounter(stack(2, 1)).Func()
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func (p PCounter) Dir() (dir string) {
	return path.Dir(p.frame().File)
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func Dir() string {
	return PCounter(stack(2, 1)).Dir()
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func (p PCounter) File() (file string) {
	return p.frame().File
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func File() string {
	return PCounter(stack(2, 1)).File()
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (p PCounter) Line() (line int) {
	return p.frame().Line
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
	if len(p) < 1 {
		return
	}

	frames := runtime.CallersFrames(p[1:])
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

func (p PCounter) frame() runtime.Frame {
	if len(p) < 1 || p[0] == 0 {
		return runtime.Frame{}
	}
	return *(*runtime.Frame)(unsafe.Pointer(p[0]))
}
