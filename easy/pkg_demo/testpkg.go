package main

import "Gogoyan/easy/pkg_demo/mytest"

func main() {
	mytest.PrintAdd(1, 2)
	mytest.PrintSub(1, 2)
}

// init 函数调用顺序
// 1. 如果包互相依赖，则先调用最早被依赖的包的init函数；否则按照import顺序调用各自包的init函数；
// 2. 同一个包内的文件，按照文件名的字母序依次调用各个文件内的init函数；
// 3. 同一个文件内，按照顺序调用init函数。
