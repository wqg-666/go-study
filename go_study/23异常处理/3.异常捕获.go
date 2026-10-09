package main

import (
	"fmt"
	"runtime/debug"
)

func read() {

	var list []int = []int{1, 2}
	fmt.Println(list[2])
}
func main1() {
	//而 defer 的作用是：让这个匿名函数不要现在执行，而是在当前函数即将退出时执行。
	defer func() {
		err := recover() //recover = 收尾的人负责把 panic 接住
		if err != nil {
			fmt.Println(err)
			fmt.Println(string(debug.Stack())) //可以看到哪里报错了
		}
	}()
	read()
}
func main() {
	main1()
	fmt.Println("正常逻辑")
}
