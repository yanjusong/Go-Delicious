package main

import (
	"fmt"
	"runtime"
)

func main() {
	pc := make([]uintptr, 1024)
	for skip := 0; ; skip++ {
		n := runtime.Callers(skip, pc)
		if n <= 0 {
			break
		}
		fmt.Printf("skip = %v, pc = %v\n", skip, pc[:n])
	}
}

// Output:
// skip = 0, pc = [16808161 17534566 16965087 17138865]
// skip = 1, pc = [17534566 16965087 17138865]
// skip = 2, pc = [16965087 17138865]
// skip = 3, pc = [17138865]

// Explanation:
// 比`runtime.Caller`多了一步调用，因为`runtime.Callers`多了一步内部调用。
