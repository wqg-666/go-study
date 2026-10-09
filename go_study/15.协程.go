package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan = make(chan int) //初始化长度为0的信道里面存int类型,无缓冲通道（长度 0）
func pay(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan <- money //把钱存入信道
	wait.Done()
}
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(3) //3个协程线程
	//主线程结束，协程线程跟着结束
	go pay("张三", 2, &wait)
	go pay("王五", 3, &wait)
	go pay("李四", 5, &wait)
	// 单独协程：等待所有购物协程完成，关闭通道
	go func() {
		//1或者2作用
		//1.defer close(moneyChan)
		wait.Wait()
		//2.close(moneyChan)
	}()
	//通道可以用 for range 自动循环，通道关闭自动退出
	var moneyList []int
	for money := range moneyChan {
		fmt.Println(money)
		moneyList = append(moneyList, money)
	}
	//for {
	//	money, ok := <-moneyChan
	//	fmt.Println(money, ok)
	//	if !ok {
	//		break
	//	}
	//}

	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
}
