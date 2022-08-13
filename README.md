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

## Use case 1

Simplified functions.

```
package main

import (
    "fmt"

    "github.com/thecxx/runpoint"
)

func main() {
	fmt.Printf("Current package: %s\n", runpoint.Package())
	fmt.Printf("Current file: %s\n", runpoint.File())
	fmt.Printf("Current line: %d\n", runpoint.Line())
}
```

## Use case 2

Generate `runpoint.PCounter` object at run point via `runpoint.PC()`.

```
package main

import (
    "fmt"

    "github.com/thecxx/runpoint"
)

func main() {
	pc := runpoint.PC()
	fmt.Printf("Package: %s\n", pc.Package())
	fmt.Printf("File: %s\n", pc.File())
	fmt.Printf("Line: %d\n", pc.Line())

	pc.Frames(func(f runpoint.Frame) {
		fmt.Printf("Package: %s\n", f.Package())
		fmt.Printf("File: %s\n", f.File())
		fmt.Printf("Line: %d\n", f.Line())
	})
}
```

## Examples

```
// PackFull() // example: github.com/thecxx/runpoint
// Package()  // example: runpoint
// FuncFull() // example: github.com/thecxx/runpoint.(PCounter).FuncFull
// Receiver() // example: PCounter
// FuncLong() // example: (PCounter).FuncFull
// Function() // example: FuncFull
// Dir()
// File()
// Filename()
// Line()
```