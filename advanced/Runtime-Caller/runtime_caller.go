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
}

// Output:
// skip = 0, pc = 17534359, file = /Users/jusong.yan/go/src/Gogoyan/advanced/Tuntime-Caller/runtime_caller.go, line = 10
// skip = 1, pc = 16964958, file = /usr/local/go/src/runtime/proc.go, line = 200
// skip = 2, pc = 17138736, file = /usr/local/go/src/runtime/asm_amd64.s, line = 1337

// Explanation:
// 1.runtime.goexit 为真正的函数入口(并不是main.main)
// 2.然后 runtime.goexit 调用 runtime.main 函数
// 3.最终 runtime.main 调用用户编写的 main.main 函数
