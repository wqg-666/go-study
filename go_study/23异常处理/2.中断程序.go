package main

import (
	"fmt"
	"os"
)

func init() {
	//初始化出错
	_, err := os.ReadFile("12344")
	//data → 文件内容 //err  → 有没有发生错误
	if err != nil { //读取文件失败
		//return//这个 return 只负责退出 init()。 它不是退出整个程序。输出main
		//log.Fatalln("错误了") //输出。2026/09/16 22:42:26 错误了
		panic("错误了") //输出，panic: 错误了
	}
}
func main() {
	fmt.Println("main")
}
