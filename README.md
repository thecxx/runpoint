<!--
 Copyright 2022 Kami
 
 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at
 
     http://www.apache.org/licenses/LICENSE-2.0
 
 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
-->

# Introduction

`runpoint` is a package for getting runtime environment.

# Getting started

```
package main

import (
    "fmt"

    "github.com/goentf/runpoint"
)

func main() {
	fmt.Printf("Current package: %s\n", runpoint.Package())
	fmt.Printf("Current file: %s\n", runpoint.File())
	fmt.Printf("Current line: %d\n", runpoint.Line())

	pc := runpoint.PC()
	fmt.Printf("Current package: %s\n", pc.Package())
	fmt.Printf("Current file: %s\n", pc.File())
	fmt.Printf("Current line: %d\n", pc.Line())

	pc.Frames(func(f runpoint.Frame) {
		fmt.Printf("Current package: %s\n", f.Package())
		fmt.Printf("Current file: %s\n", f.File())
		fmt.Printf("Current line: %d\n", f.Line())
	})

	// Others:
	// runpoint.PackFull() // example: github.com/goentf/runpoint
	// runpoint.Package()  // example: runpoint
	// runpoint.FuncFull() // example: github.com/goentf/runpoint.FuncFull
	// runpoint.Receiver() // example: PCounter
	// runpoint.FuncLong() // example: (PCounter).FuncFull
	// runpoint.Function() // example: FuncFull
	// runpoint.Dir()
	// runpoint.File()
	// runpoint.Filename()
	// runpoint.Line()
}
```