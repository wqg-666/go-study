package main

import "fmt"

// 值传递
func Copy(name string) {
	fmt.Println("infunc %p\n", &name)
	name = "ww" //这个的值修改不会影响主函数中的name的值
}

// 引用传递，传递内存地址
func Set(name *string) {
	//这个name是内存地址
	//*name,通过内存地址去找值

	*name = "nihao" //主函数中的name也变了
	fmt.Println("infunc %p\n", name)
}
func main() {
	var name = "qq"
	//输出内存地址
	fmt.Println("%p\n", &name)
	//&取地址，*解引用，去这个地址的值
	Set(&name)
	Copy(name) //这个会产生一个新的name地址
	fmt.Println(name)

}
