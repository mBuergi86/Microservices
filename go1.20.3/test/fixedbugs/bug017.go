// run

// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

func main() {
	var s2 string = "\a\b\f\n\r\t\v" // \r is miscompiled
	_ = s2
}

/*
server.go.c: In function ‘main_main’:
server.go.c:20: error: missing terminating " character
server.go.c:21: error: missing terminating " character
server.go.c:24: error: ‘def’ undeclared (first use in this function)
server.go.c:24: error: (Each undeclared identifier is reported only once
server.go.c:24: error: for each function it appears in.)
server.go.c:24: error: syntax error before ‘def’
server.go.c:24: error: missing terminating " character
server.go.c:25: warning: excess elements in struct initializer
server.go.c:25: warning: (near initialization for ‘slit’)
server.go.c:36: error: syntax error at end of input
*/
