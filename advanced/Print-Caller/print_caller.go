package main

import (
	"fmt"
	"runtime"
)

var a = PrintCallerName(0, "main.a")
var b = PrintCallerName(0, "main.b")

func init() {
	a = PrintCallerName(0, "main.init.a")
}

func init() {
	b = PrintCallerName(0, "main.init.b")
	func() {
		b = PrintCallerName(0, "main.init.b[1]")
	}()
}

func main() {
	a = PrintCallerName(0, "main.main.a")
	b = PrintCallerName(0, "main.main.b")
	func() {
		b = PrintCallerName(0, "main.main.b[1]")
		func() {
			b = PrintCallerName(0, "main.main.b[1][1]")
		}()
		b = PrintCallerName(0, "main.main.b[2]")
	}()
}

// CallerName ...
func CallerName(skip int) (name, file string, line int, ok bool) {
	var pc uintptr
	if pc, file, line, ok = runtime.Caller(skip + 1); !ok {
		return
	}
	name = runtime.FuncForPC(pc).Name()
	return
}

// PrintCallerName ...
func PrintCallerName(skip int, comment string) bool {
	name, file, line, ok := CallerName(skip + 1)
	if !ok {
		return false
	}
	fmt.Printf("skip = %v, comment = %s\n", skip, comment)
	fmt.Printf("  file = %v, line = %d\n", file, line)
	fmt.Printf("  name = %v\n", name)
	return true
}
