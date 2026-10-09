package main

import (
	"fmt"
	"sync"
	"time"
)

func shopping(name string, wait *sync.WaitGroup) {
	fmt.Printf("%s 开始购物\n", name)
	time.Sleep(1 * time.Second)
	fmt.Printf("%s 购物结束\n", name)
	wait.Done()
}
func main() {
	var wait sync.WaitGroup
	startTime := time.Now()
	//现在的模式是购物接力
	//shopping("张三")
	//shopping("王五")
	//shopping("李四")
	wait.Add(3) //3个协程线程
	//主线程结束，协程线程跟着结束
	go shopping("张三", &wait)
	go shopping("王五", &wait)
	go shopping("李四", &wait)
	//主线程，main函数，main协程开始运行，记录startTime
	//连续执行 3 条go shopping()：
	//只是创建 3 个 goroutine，放入调度队列，不会阻塞 main 函数
	//这 3 个协程并发执行（几乎同时启动）
	//然后 main 协程执行 time.Sleep(2*time.Second)：main 协程休眠 1s
	//三个 shopping 协程：
	//几乎同一时刻打印 xx 开始购物
	//然后执行 time.Sleep(1*time.Second)，协程休眠 1 秒
	//关键点：
	//三个 shopping 协程 sleep 的 1 秒 和 main 协程 sleep 的 1 秒是并行的
	//当 1 秒时间到：
	//三个 shopping 协程被唤醒，打印xx 购物结束
	//main 协程也被唤醒，打印购买完成
	//time.Sleep(2 * time.Second)
	wait.Wait()
	//阻塞 main 协程，一直等到计数器减到 0（所有子 goroutine 都调用完 Done ()）之后，main 才继续往下执行。
	fmt.Println("购买完成", time.Since(startTime))
}
