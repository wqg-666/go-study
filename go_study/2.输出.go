package main

import "fmt"

func main() {
	//自带换行
	//两个字母.数字间有空格
	fmt.Println("h", "l")
	fmt.Println(1, "h")
	fmt.Println("nihao")
	//没有换行
	//\n,表示换行,需要占位符，%s,%d,%.2f
	fmt.Printf("%s 你好呀\n", "qq")
	fmt.Printf("%s 你好呀\n", "qq")
	fmt.Printf("%d\n", 30)
	fmt.Printf("%.2f\n", 2.31895) //输出两位小数2.32（四舍五入）
	fmt.Printf("%f\n", 2.335152)
	//打印类型
	fmt.Printf("%T %T\n", "你好", 2.5)
	//任意类型
	fmt.Printf("%v\n", "f")
	fmt.Printf("%#v\n", " ") //输出" "
	fmt.Printf("%v\n", " ")  //输出为空3
	var f = fmt.Sprintf("%.2f", 2.31895)
	fmt.Printf(f) //输出2.32
}
