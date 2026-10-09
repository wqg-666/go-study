package main

import (
	"fmt"
)

// 无返回值，不用写return
func sayhello() {
	fmt.Println("hello")
}
func param1(id string) { //使用时要传参
	fmt.Println(id)
}
func param2(id string, userName string) {
	fmt.Println(id)
}
func param3(id, userName string) { //同类型，前一个string可以省略
	fmt.Println(id)
}
func add(numberList ...int) { //多个参数
	var sum int
	//第一个返回值：下标索引。第二个返回值：元素本身（就是数组里的数字）
	for _, i2 := range numberList {
		sum += i2
	}
	fmt.Println(sum)
}
func r1() {
	return //没有返回值,或者可以省略不写
}

// 有返回值要用return
func r2() bool {
	return true
}
func r3() int {
	var ok int
	return ok
}
func r6() (ok bool) {
	return ok //或者直接return,bool默认返回值为false
}
func r4() (string, bool) {
	return "", true
}
func r5() (val string, ok bool) {
	return val, ok
	//或者return
}
func main() {
	sayhello()
	param1("2")
	param2("1", "2")
	param3("1", "2")
	add(1, 2)
	add(1, 2, 3)
	add(1, 2, 3, 4)
	fmt.Println(r6())
	fmt.Println(r4())
	//匿名函数，在函数内创建函数
	var getName = func() string {
		return "qq"
	}
	fmt.Println(getName())
	var setName = func(name string) {
		fmt.Println(name)
		return
	}
	setName("张三")
	//高阶函数
	fmt.Println("请输入你要执行的操作：")
	fmt.Println("1 登录")
	fmt.Println("2 注册")
	fmt.Println("3 个人中心")
	var index int
	fmt.Scan(&index)
	/////////////////////
	var funMap = map[int]func(){ //函数类型的 map
		1: login, //代表函数本身，把函数当作值存入 map
		2: register,
		3: userCenter,
	}
	fun, ok := funMap[index]
	//如果用户输入1 → fun = login，对应的函数，ok=true,是否存在
	//ok为true说明这个数字有对应的函数，fun()就执行这个函数。
	if ok {
		fun()
	}
	//////////////////////
	//switch可以选择函数
	//switch index {
	//case 1:
	//	login()
	//case 2:
	//	register()
	//case 3:
	//	userCenter()
	//}

}
func login() {
	fmt.Println("登录")
}
func register() {
	fmt.Println("注册")
}
func userCenter() {
	fmt.Println("个人中心")
}
