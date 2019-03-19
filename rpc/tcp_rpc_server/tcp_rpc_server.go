package main

import (
	"errors"
	"fmt"
	"net"
	"net/rpc"
	"os"
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
		ResolveTCPAddr将addr作为TCP地址解析并返回。参数addr格式为"host:port"或"[ipv6-host%zone]:port"，解析得到网络名和端口名；net必须是"tcp"、"tcp4"或"tcp6"。

		IPv6地址字面值/名称必须用方括号包起来，如"[::1]:80"、"[ipv6-host]:http"或"[ipv6-host%zone]:80"。
	*/
	tcpAddr, err := net.ResolveTCPAddr("tcp", ":1234")
	checkError(err)
	/*
		ListenTCP在本地TCP地址laddr上声明并返回一个*TCPListener，net参数必须是"tcp"、"tcp4"、"tcp6"，如果laddr的端口字段为0，函数将选择一个当前可用的端口，可以用Listener的Addr方法获得该端口。
	*/
	listener, err := net.ListenTCP("tcp", tcpAddr)
	checkError(err)

	for {
		// AcceptTCP接收下一个呼叫，并返回一个新的*TCPConn。
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		/*
			ServeConn在单个连接上执行DefaultServer。ServeConn会阻塞，服务该连接直到客户端挂起。调用者一般应另开线程调用本函数："go ServeConn(conn)"。ServeConn在该连接使用gob（参见encoding/gob包）有线格式。要使用其他的编解码器，可调用ServeCodec方法。
		*/
		go rpc.ServeConn(conn)
	}

}

func checkError(err error) {
	if err != nil {
		fmt.Println("Fatal error ", err.Error())
		os.Exit(1)
	}
}
