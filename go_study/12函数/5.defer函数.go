package main

import "fmt"

// 关键字 defer 用于注册延迟调用
// 这些调用直到 return 前才被执。因此，可以用来做资源清理
// 多个defer语句，按先进后出的方式执行，谁离return近谁先执行
// defer语句中的变量，在defer声明时就决定了
func main() {
	defer func() {
		fmt.Println("defer2")
	}()
	//defer fmt.Println("defer2")
	defer func() {
		fmt.Println("defer1")
	}()
	//defer fmt.Println("defer1")
	//输出defer1,defer2,谁靠近return谁先执行
	var age = 18
	defer func() {
		//fmt.Println(name)//不能打印，defer语句中的变量，在defer之前声明可以用，之后不能用
		fmt.Println(age)
	}()
	//var name ="qq"

	return
}
