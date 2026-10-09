package main

//导包
import (
	"fmt"
	name4 "study1/go_study/name"
)

func hello() {
	fmt.Println("hello world")
	//fmt.Println(age)
}

// 全局变量
var age = 12
var (
	s1 string = "1"
	s2 string = "2"
)

// 常量
const version string = "2.0.1"

func main() {
	//首字母大写的变量、函数。方法，属性可在包外进行访问
	fmt.Println(name4.Name)
	//fmt.Println(name4.version) //
	//先声明
	var name string
	//后赋值
	name = "qq"
	fmt.Println(name)

	//声明并赋值
	var name1 string = "qq"
	fmt.Println(name1)
	//省略类型
	var name2 = "qq"
	fmt.Println(name2)
	//声明并赋值 短声明符合 ：=
	name3 := "qq"
	fmt.Println(name3)
	hello()
	var a1, a2 = 1, 2
	fmt.Println(a1, a2)
	var (
		s1 string = "1"
		s2 string = "2"
	)
	fmt.Println(s1, s2)
	a11 := 1
	a11, a12 := 2, 3
	fmt.Println(a11, a12) //输出2，3
}
