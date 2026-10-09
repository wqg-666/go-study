package main

import "fmt"

type SingInterface interface {
	Sing()           //没有返回值
	GetName() string //string是它的返回值类型
}

// Chicken 和 Cat 都写了这两个方法，所以它们都自动实现了 SingInterface
type Chicken struct {
	Name string
}
type Cat struct {
	Name string
}

// func (接收者 类型) 方法名(参数) 返回值/1. Chicken的方法
func (c Chicken) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Chicken) GetName() string {
	return c.Name
}
func (c Cat) Sing() {
	fmt.Println(c.Name, "在唱歌")
}
func (c Cat) GetName() string {
	return c.Name
}

// 空接口1,任何类型都实现了空接口的定义
type EmptyInterface interface {
}

func print(val EmptyInterface) {
	fmt.Println(val)
}

// 空接口2
func print2(val interface{}) {
	fmt.Println(val)
}

// 空接口3
func print3(val any) {
	fmt.Println(val)
}

// func 函数名(参数) 返回值/2. 普通函数
func sing(c SingInterface) {
	//SingInterface。所以cat,chicken都可以用这个方法
	//指定断言。。{ik} true   {} false
	ch, ok := c.(Chicken)
	fmt.Println(ch, ok)
	//指定断言。。{ik}   {mimi}
	switch server := c.(type) {
	case Cat:
		fmt.Println(server)
	case Chicken:
		fmt.Println(server)
	default:
		fmt.Println("其他")
	}
	c.Sing()
	fmt.Println(c.GetName())
}
func main() {
	ch := Chicken{Name: "ik"}
	ca := Cat{Name: "mimi"}
	sing(ch) //ik 在唱歌。ik
	//ch.Sing() //c的方法，ik 在唱歌。
	sing(ca) //mimi 在唱歌。mimi。
	//ca.ch
	print(ca) //{mimi}
	print(ch) //{ik}
}
