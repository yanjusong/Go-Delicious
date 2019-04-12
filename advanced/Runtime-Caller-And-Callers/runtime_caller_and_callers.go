package main

import (
	"fmt"
	"runtime"
)

func main() {
	for skip := 0; ; skip++ {
		pc, file, line, ok := runtime.Caller(skip)
		if !ok {
			break
		}
		fmt.Printf("skip = %v, pc = %v, file = %v, line = %v\n", skip, pc, file, line)
	}
	// Output:
	// skip = 0, pc = 4198456, file = caller.go, line = 10
	// skip = 1, pc = 4280962, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 220
	// skip = 2, pc = 4290608, file = $(GOROOT)/src/pkg/runtime/proc.c, line = 1394
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pc)
		if n <= 0 {
			break
		}
		fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
	}
	// Output:
	// skip = 0, pc = [4305334 4198635 4280962 4290608]
	// skip = 1, pc = [4198635 4280962 4290608]
	// skip = 2, pc = [4280962 4290608]
	// skip = 3, pc = [4290608]
}

// Output from self PC:
// skip = 0, pc = 17534487, file = /Users/jusong.yan/go/src/Gogoyan/advanced/Runtime-Caller-And-Callers/runtime_caller_and_callers.go, line = 10
// skip = 1, pc = 16965086, file = /usr/local/go/src/runtime/proc.go, line = 200
// skip = 2, pc = 17138864, file = /usr/local/go/src/runtime/asm_amd64.s, line = 1337
// skip = 0, pc = [16808161 17535248 16965087 17138865]
// skip = 1, pc = [17535248 16965087 17138865]
// skip = 2, pc = [16965087 17138865]
// skip = 3, pc = [17138865]
