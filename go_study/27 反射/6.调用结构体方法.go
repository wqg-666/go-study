package main

import (
	"fmt"
	"reflect"
)

type User1 struct {
}

// 属于 User1 这个结构体的方法,(User1)：接收者，代表这个方法绑定在 User1 结构体上，方法名 Call，首字母大写才能被反射找到
func (User1) Call(name string) {
	fmt.Println("我被调用了", name)
}

// User1.Call 是绑定在 User1 上的方法 ,Call 是包下的普通函数, 同名完全没问题，不会冲突。
// 普通函数 function，包下面独立的普通函数
func Call(obj any) {
	////反射逻辑，找到 obj 里面的 Call【方法】，然后调用
	// obj传进来是 &s，*User指针。 Elem()解指针
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	//v.NumMethod()：这个结构体一共有多少个【可导出】的方法（大写方法）
	for i := 0; i < v.NumMethod(); i++ {
		//t.Method(i)：拿到第 i 个方法的【类型信息】，方法名、参数、返回值等
		m := t.Method(i)
		//判断方法名字是不是 "Call"，不是就跳过
		if m.Name != "Call" {
			continue
		}
		// v.Method(i)：拿到第 i 个方法的【值对象】，用来执行调用
		method := v.Method(i)
		// method.Call([]reflect.Value{})：反射调用方法，参数必须包装成 []reflect.Value 切片，里面每一项对应方法的入参
		//这一行，就是调用结构体的 Call 方法！
		method.Call([]reflect.Value{
			reflect.ValueOf("qq"), //把字符串 "qq" 包装成 reflect.Value，作为第一个参数
		})
	}
}

func main() {
	s := User1{}
	s.Call("aaa") //调用【方法】.我被调用了 aaa
	Call(&s)      //我被调用了 qq
	//s.Call("aaa")：硬编码，编译就确定调用，简单高效。
	//Call(&s)：调用你写的反射工具函数，运行时去对象里面搜 Call 方法，然后执行，可以支持多种不同类型。
}
