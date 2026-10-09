package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	//向 http://127.0.0.1:80 发送GET请求
	response, err := http.Get("http://127.0.0.1:80")
	if err != nil {
		fmt.Println(err)
		return
	}
	// 读取响应体里面全部字节数据
	byteDate, _ := io.ReadAll(response.Body)
	fmt.Println(string(byteDate))
}
