package main

import (
	"fmt"
)

func main() {
	var sum int
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		sum += i
		sum = sum + i
	}
	//死循环
	//for i := 1; true ; i++ {
	//
	//}
	//for true{
	//
	//}
	//for{
	//
	//}
	//while模式
	var sum1 int
	var i int = 1
	for i <= 100 { //先判断条件再走循环体
		sum1 += i
		i++
	}
	fmt.Println(sum1)
	//do-while模式
	var sum2 int
	var i2 int = 1
	for { //先走循环体
		sum2 += i2
		i2++
		if i2 == 101 { //再循环体里判断条件
			break
		}
	}
	fmt.Println(sum2)
	//遍历切片
	var list = []string{"ww", "qq", "gg"}
	for i := 0; i < len(list); i++ {
		fmt.Println(i, list[i])
	}
	for i, i2 := range list {
		fmt.Println(i, i2) //索引，元素
	}
	//循环map
	var userMap = map[string]string{"name": "qq"}
	for i, i2 := range userMap {
		fmt.Println(i, i2) //name qq
	}
	//break，continue
	//break用于跳出当前循环
	//continue用于跳过本轮循环
	for i := 1; i < 9; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println() //换行
	}
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if j > i {
				// 去掉 列比行大的数据
				continue
			}
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
	for i := 1; i <= 10; i++ {
		if i == 5 {
			continue //输出,1,2,3,4,6,7,8,9,10
			break    //输出1，2，3，4
		}
		fmt.Println(i)
	}

}
