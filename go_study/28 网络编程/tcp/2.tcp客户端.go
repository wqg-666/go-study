package main

import (
	"fmt"
	"io"
	"net"
)

// 客户端向服务端发送消息
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println(err)
		return
	}
	//开启一个协程，专门用来【读取服务端返回的数据】,只负责读
	//单独开协程，收发可以同时进行，互不阻塞。
	go func() {
		for {
			var buf []byte = make([]byte, 1024)
			n, err := conn.Read(buf)
			//服务端关闭连接时触发EOF
			if err == io.EOF {
				break
			}

			//输出的是：服务端发回来的字符串
			fmt.Println(string(buf[0:n]))

		}
	}()
	//主协程：负责【控制台输入，向服务端发送消息】,只负责输入 + 发送
	for {
		var text string

		fmt.Printf("请输入你要输入的内容: ")

		fmt.Scan(&text)
		if text == "q" {
			conn.Close()
			break
		}
		//把输入的文字发给服务端
		conn.Write([]byte(text))
	}
}
