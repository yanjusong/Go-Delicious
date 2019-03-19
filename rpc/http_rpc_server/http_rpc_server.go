package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

type Arith int

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

func main() {

	// 服务端会调用（用于HTTP服务）
	arith := new(Arith)
	/*
		Register在server注册并公布rcvr的方法集中满足如下要求的方法：

		- 方法是导出的
		- 方法有两个参数，都是导出类型或内建类型
		- 方法的第二个参数是指针
		- 方法只有一个error接口类型的返回值
		如果rcvr不是一个导出类型的值，或者该类型没有满足要求的方法，Register会返回错误。Register也会使用log包将错误写入日志。客户端可以使用格式为"Type.Method"的字符串访问这些方法，其中Type是rcvr的具体类型。
	*/
	rpc.Register(arith)
	/*
		HandleHTTP函数注册DefaultServer的RPC信息HTTP处理器对应到DefaultRPCPath，和DefaultServer的debug处理器对应到DefaultDebugPath。HandleHTTP函数会注册到http.DefaultServeMux。之后，仍需要调用http.Serve()，一般会另开线程："go http.Serve(l, nil)"
	*/
	rpc.HandleHTTP()

	/*
		ListenAndServe监听srv.Addr指定的TCP地址，并且会调用Serve方法接收到的连接。如果srv.Addr为空字符串，会使用":http"。
		ListenAndServe使用指定的监听地址和处理器启动一个HTTP服务端。处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。Handle和HandleFunc函数可以向DefaultServeMux添加处理器。
	*/
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}
