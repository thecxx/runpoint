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
	"sync"
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
	// Program counter
	PCounter struct {
		st []uintptr
		lz sync.Once
		fr runtime.Frame
		// Cache the result of function splitFuncFull
		pf, pn, fl, rn, fn string
	}
)

// PC returns a PCounter.
func PC(skip ...int) *PCounter {
	s := 0
	if len(skip) == 1 {
		if skip[0] < 0 {
			panic(errors.New("invalid skip value"))
		}
		s = skip[0]
	}
	return &PCounter{st: stack(s+2, int(depth))}
}

// FuncFull returns the full name of the function.
//
// Example:
// 		"github.com/goentf/runpoint.FuncFull"
// 		"github.com/goentf/runpoint.(*PCounter).FuncFull"
// 		"github.com/goentf/runpoint.(*PCounter).FuncFull.func1"
func (p *PCounter) FuncFull() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.fr.Function
}

// PackFull returns the full package name of the function.
//
// Example:
// 		"github.com/goentf/runpoint"
func (p *PCounter) PackFull() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.pf
}

// Package returns the package name of the function.
//
// Example:
// 		"runpoint"
func (p *PCounter) Package() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.pn
}

// FuncLong returns the long name of the function.
//
// Example:
//		"FuncLong"
//		"FuncLong.func1"
//		"FuncLong.func2"
// 		"(*PCounter).FuncLong"
// 		"(*PCounter).FuncLong.func1"
func (p *PCounter) FuncLong() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.fl
}

// Receiver returns the receiver type of the function.
//
// Example:
//		"*PCounter"
func (p *PCounter) Receiver() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.rn
}

// Function returns the name of the function.
//
// Example:
//		"Function"
func (p *PCounter) Function() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.fn
}

// Dir returns the directory path of the
// source code corresponding to the program counter pc.
func (p *PCounter) Dir() (dir string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	if p.fr.PC != 0 {
		dir = path.Dir(p.fr.File)
	}
	return
}

// File returns the file path of the
// source code corresponding to the program counter pc.
func (p *PCounter) File() (file string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.fr.File
}

// Filename returns the file name of the
// source code corresponding to the program counter pc.
func (p *PCounter) Filename() (name string) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	if p.fr.PC != 0 {
		name = path.Base(p.fr.File)
	}
	return
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func (p *PCounter) Line() (line int) {
	if p.fr.PC == 0 {
		p.lazyLoad()
	}
	return p.fr.Line
}

// Frames is used to get all the stack frames.
func (p *PCounter) Frames(fun func(Frame)) (num int) {
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

func (p *PCounter) lazyLoad() {
	p.lz.Do(func() {
		if len(p.st) < 1 {
			return
		}
		p.fr, _ = runtime.CallersFrames(p.st[0:1]).Next()
		if p.fr.PC != 0 {
			p.pf, p.pn, p.fl, p.rn, p.fn = splitFuncFull(p.fr.Function)
		}
	})
}

// FuncFull returns the full name of the function.
//
// Example:
// 		"github.com/goentf/runpoint.FuncFull"
// 		"github.com/goentf/runpoint.(*PCounter).FuncFull"
// 		"github.com/goentf/runpoint.(*PCounter).FuncFull.func1"
func FuncFull() (name string) {
	return frame(2).Function
}

// PackFull returns the full package name of the function.
//
// Example:
// 		"github.com/goentf/runpoint"
func PackFull() (name string) {
	name, _, _, _, _ = splitFuncFull(frame(2).Function)
	return
}

// Package returns the package name of the function.
//
// Example:
// 		"runpoint"
func Package() (name string) {
	_, name, _, _, _ = splitFuncFull(frame(2).Function)
	return
}

// FuncLong returns the long name of the function.
//
// Example:
//		"FuncLong"
//		"FuncLong.func1"
//		"FuncLong.func2"
// 		"(*PCounter).FuncLong"
// 		"(*PCounter).FuncLong.func1"
func FuncLong() (name string) {
	_, _, name, _, _ = splitFuncFull(frame(2).Function)
	return
}

// Receiver returns the receiver type of the function.
//
// Example:
//		"*PCounter"
func Receiver() (name string) {
	_, _, _, name, _ = splitFuncFull(frame(2).Function)
	return
}

// Function returns the name of the function.
//
// Example:
//		"Function"
func Function() (name string) {
	_, _, _, _, name = splitFuncFull(frame(2).Function)
	return
}

// Dir returns the directory path of the
// source code corresponding to the program counter pc.
func Dir() (dir string) {
	if frame := frame(2); frame.PC != 0 {
		dir = path.Dir(frame.File)
	}
	return
}

// File returns the file path of the
// source code corresponding to the program counter pc.
func File() (file string) {
	return frame(2).File
}

// Filename returns the file name of the
// source code corresponding to the program counter pc.
func Filename() (name string) {
	if frame := frame(2); frame.PC != 0 {
		name = path.Base(frame.File)
	}
	return
}

// Line returns the line number of the
// source code corresponding to the program counter pc.
func Line() (line int) {
	return frame(2).Line
}
