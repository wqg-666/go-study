package main

import (
	"fmt"
	"io"
	"net"
)

func main() {
	// 创建tcp的监听地址
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:8080")
	// tcp监听
	listen, _ := net.ListenTCP("tcp", tcpAddr)
	for {
		// 等待连接
		conn, err := listen.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		// 获取客户端的地址
		fmt.Println(conn.RemoteAddr().String() + " 进来了")
		// 读取客户端传来的数据
		for {
			//make([]byte,1024) 创建缓冲区存放网络字节
			var buf []byte = make([]byte, 1024)
			//conn.Read(buf) 阻塞读取客户端发送的数据
			n, err := conn.Read(buf)
			//如果 err == io.EOF：代表客户端关闭连接，打印客户端离开，跳出内层循环，继续等待下一个客户端
			if err == io.EOF {
				fmt.Println(conn.RemoteAddr().String() + " 出去了")
				break
			}
			//打印收到的客户端消息
			fmt.Println(string(buf[0:n]))
			//conn.Write() 回复消息给客户端：你说的是xxx?
			conn.Write([]byte("你说的是" + string(buf[0:n]) + "?"))
		}
	}
}
