package main

import "fmt"

func main() {
	var age int
	fmt.Println("请输入你的年龄")
	fmt.Scan(&age)
	//中断式，卫语句
	if age <= 0 {
		fmt.Println("未出生")
		return
	}
	if age <= 18 {
		fmt.Println("未成年")
		return
	}
	if age <= 35 {
		fmt.Println("青年")
		return
	}
	fmt.Println("中年")
	//嵌套式
	if age <= 18 {
		if age <= 0 {
			fmt.Println("未出生")
		} else {
			fmt.Println("未成年")
		}
	} else {
		if age <= 35 {
			fmt.Println("青年")
		} else {
			fmt.Println("中年")
		}
	}
	//多条件
	if age <= 0 {
		fmt.Println("未出生")
	}
	if age <= 18 && age > 0 {
		fmt.Println("未成年")
	}
	if age <= 35 && age > 18 {
		fmt.Println("青年")
	}
	if age > 35 {
		fmt.Println("中年")
	}
	//&&,||，！。用法与c++一样
	//逻辑短路
	//&&。第一个条件为false,后面的条件就不会走了
	//||。第一个条件为true，后面的条件就不会走了
}
