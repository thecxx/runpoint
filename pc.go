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
	PCounter struct {
		st []uintptr
		fr runtime.Frame
	}
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
	p := PCounter{st: stack(s+2, int(depth))}
	if len(p.st) > 0 {
		p.fr, _ = runtime.CallersFrames(p.st[0:1]).Next()
	}
	return p
}

// Func returns the name of the function.
func (p PCounter) Func() (name string) {
	if p.fr.PC != 0 {
		name = p.fr.Function
	}
	return
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func (p PCounter) Dir() (dir string) {
	if p.fr.PC != 0 {
		dir = path.Dir(p.fr.File)
	}
	return
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func (p PCounter) File() (file string) {
	if p.fr.PC != 0 {
		file = p.fr.File
	}
	return
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (p PCounter) Line() (line int) {
	if p.fr.PC != 0 {
		line = p.fr.Line
	}
	return
}

// Frames is used to get all the stack frame.
func (p PCounter) Frames(fun func(Frame)) (num int) {
	if fun == nil {
		return
	}
	if len(p.st) < 1 {
		return
	}

	frames := runtime.CallersFrames(p.st)
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

// Func returns the name of the function.
func Func() string {
	return frame(2).Function
}

// Dir returns the directory name of the
// source code corresponding to the program counter pc.
func Dir() string {
	return path.Dir(frame(2).File)
}

// File returns the file name of the
// source code corresponding to the program counter pc.
func File() string {
	return frame(2).File
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func Line() int {
	return frame(2).Line
}

func frame(skip int) (f runtime.Frame) {
	f, _ = runtime.CallersFrames(stack(skip+1, 1)).Next()
	return
}
