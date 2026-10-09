package main

import (
	"fmt"
	"time"
)

func awaitAdd(awaitSecond int) func(...int) int {
	//1.time.Sleep(time.Duration(awaitSecond) * time.Second)
	return func(List ...int) int {
		//2.time.Sleep(time.Duration(awaitSecond) * time.Second)
		//这里的awaitSecond调用了awaitAdd(awaitSecond int)的awaitSecond。是闭包
		var sum int
		for _, i2 := range List {
			sum += i2
		}
		return sum
	}
}

func main() {
	//先传一个参数表示延时，后面再次传参数就是将参数求和
	//fun(2)(1,2,3),延时2秒，求1+2+3
	t1 := time.Now() // // 获取【当前时间点】，记录开始时刻
	sum := awaitAdd(2)(1, 2, 3)
	subTime := time.Since(t1) //计算从t1这个时间点，到现在这一刻，过去了多长时间
	fmt.Println(sum, subTime) //1,2都是这个情况。6 2.0003373s
	//1.2位置不同的影响
	add := awaitAdd(2) //2。外层函数飞快执行，只返回闭包，**不会sleep**
	//1。执行awaitAdd(2)，这里就会sleep足足2秒！
	t3 := time.Now()
	sum2 := add(1, 2, 3) //2。调用闭包！这里触发sleep 2秒
	//1。add只是求和，没有延时
	subTime2 := time.Since(t3)
	fmt.Println(sum2, subTime2) //2。是6 2.0003077s，，1。是6 0s
}
