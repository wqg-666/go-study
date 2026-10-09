package main

import (
	"fmt"
)

func main() {
	//整数型
	var age = 12 //默认int.64位，匹配自己的电脑
	fmt.Println(age)
	var u8 uint8 = 255
	//var r uint8 = 256 //报错
	//fmt.Println(r)
	fmt.Println(u8)
	//uint8
	//u表示非负整数。8表示8位二进制，
	//11111111=255，2^8-1
	//00000000=0
	//int8
	//0 0000000=0
	//0 1111111=127,2^7-1
	//1 0000001=-1
	//1 0000000=-128,2^7
	//浮点型
	//float32  范围3.4e38
	//float64 默认 范围1.8e308
	//字符型
	//byte，单字节，asc里面的字符
	var a byte = 'a'
	fmt.Printf("%c %d\n", a, a) //输出a,97(在ASCII里的位置)
	var a1 uint8 = 97
	fmt.Printf("%c %d\n", a1, a1) //输出a,97(在ASCII里的位置)
	//ASC里256个字符，uint8(0~255)
	//rune ，多字节字符,中文，emoji,其他国家文字
	var z rune = '中'
	fmt.Printf("%c %d\n", z, z) //输出 中 20013
	//字符串
	var s string = "wqgnh"
	fmt.Println(s)
	//转义字符
	fmt.Println("wqg\tnihao")   //wqg	nihao
	fmt.Println("'wqg'nihao")   //'wqg'nihao
	fmt.Println("\"wqg\"nihao") //"wqg"nihao
	fmt.Println("\\wqg\\nihao") //\wqg\nihao
	fmt.Println("wqg\nnihao")   //换行
	//多行字符串
	//``中。转义字符无效
	fmt.Println(`hello
world
你好
`)
	//布尔类型
	//true,false(默认)
	//不能将整型0，1强制转换为bool型
	//不参与数值运算，不能与其他类型转换
	//零值问题,基本数据类型
	var a3 int
	fmt.Printf("%#v\n", a3) //0
	var a2 string
	fmt.Printf("%#v\n", a2) //""
	var a4 bool
	fmt.Printf("%#v\n", a4) //false

}
