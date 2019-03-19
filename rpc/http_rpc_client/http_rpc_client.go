package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

func main() {
	// if len(os.Args) != 2 {
	// 	fmt.Println("Usage: ", os.Args[0], "server")
	// 	os.Exit(1)
	// }
	// serverAddress := os.Args[1]

	/*
		DialHTTP在指定的网络和地址与在默认HTTP RPC路径监听的HTTP RPC服务端连接。
	*/
	client, err := rpc.DialHTTP("tcp", "127.0.0.1"+":1234")
	if err != nil {
		/*
			Fatal等价于{log.Print(v...); os.Exit(1)}

			Print调用log.Output将生成的格式化字符串输出到logger，参数用和fmt.Print相同的方法处理。

			Output写入输出一次日志事件。参数s包含在Logger根据选项生成的前缀之后要打印的文本。如果s末尾没有换行会添加换行符。calldepth用于恢复PC，出于一般性而提供，但目前在所有预定义的路径上它的值都为2。

			Exit让当前程序以给出的状态码code退出。一般来说，状态码0表示成功，非0表示出错。程序会立刻终止，defer的函数不会被执行。
		*/
		log.Fatal("dialing:", err)
	}
	// Synchronous call
	args := Args{17, 8}
	var reply int
	/*
		Call方法会等待远端调用完成，而Go方法异步的发送调用请求并使用返回的Call结构体类型的Done通道字段传递完成信号。
		Call调用指定的方法，等待调用返回，将结果写入reply，然后返回执行的错误状态。
	*/
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	var quot Quotient
	err = client.Call("Arith.Divide", args, &quot)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d/%d=%d remainder %d\n", args.A, args.B, quot.Quo, quot.Rem)
}
