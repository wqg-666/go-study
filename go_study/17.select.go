package main

import (
	"fmt"
	"sync"
	"time"
)

var moneyChan1 = make(chan int) //声明并初始化一个缓冲容量为0的信道
var nameChan1 = make(chan string)
var doneChan = make(chan struct{}) //声明一个用于关闭的信道，专门传“结束信号”
func send(name string, money int, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	moneyChan1 <- money //把钱存入信道
	nameChan1 <- name
	wait.Done()
}
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	wait.Add(3) //3个协程线程
	//主线程结束，协程线程跟着结束
	go send("张三", 2, &wait)
	go send("王五", 3, &wait)
	go send("李四", 5, &wait)
	// 单独协程：等待所有购物协程完成，关闭通道,防止死锁
	go func() {
		defer close(moneyChan1)
		defer close(nameChan1)
		defer close(doneChan)
		wait.Wait()
	}()

	var moneyList []int
	var nameList []string
	var event = func() {
		for {
			// select 可以同时监听多个 channel
			//并行监听全部 case，就绪的 case 候选，多个就绪就随机选一个
			select {
			case money := <-moneyChan1:
				// 在 select 里，监听 moneyChan1 通道，一旦通道有数据
				//就读取数据赋值给 money，执行这个 case 内代码。
				//从 moneyChan1 里面取出一个数据，放到 money。
				moneyList = append(moneyList, money)
			case name := <-nameChan1:
				nameList = append(nameList, name)
			case <-doneChan:
				//我们不关心传过来的东西是什么，只关心“有没有收到信号”。
				fmt.Println("结束")
				return
			}
		}
	}
	event()
	fmt.Println("购买完成", time.Since(startTime))
	fmt.Println("moneyList", moneyList)
	fmt.Println("nameList", nameList)
}

//main
//│
//├── wait.Add(3)
//│
//├── go send("张三")
//│
//├── go send("王五")
//│
//├── go send("李四")
//│
//├── go 等待协程
//│       │
//│       └── wait.Wait()
//│
//└── event()
//│
//↓
//select
//│
//┌───┼────────┐
//↓   ↓        ↓
//钱  姓名      done
//│   │        │
//↓   ↓        ↓
//收集 收集    等待
//│
//│
//三个人全部完成
//↓
//WaitGroup = 0
//↓
//wait.Wait()返回
//↓
//close(doneChan)
//↓
//done case 就绪
//↓
//打印“结束”
//↓
//return
