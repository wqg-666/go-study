package main

import (
	"fmt"
	"time"
)

var done = make(chan struct{})

func event() {
	fmt.Println("event执行开始")
	time.Sleep(2 * time.Second) //延时2s
	fmt.Println("event执行结束")
	close(done) //event 执行完毕，关闭 done channel，通知等待的人。
}
func main() {
	go event() //开一个新的 goroutine 执行 event()。
	select {
	case <-done:
		fmt.Println("协程执行完毕")
	case <-time.After(1 * time.Second):
		//创建一个 channel，1 秒之后往这个 channel 里提供一个时间值。
		//1*time.Second。协程执行完毕。。3 * time.Second。超时
		fmt.Println("超时")
		return
	}
}
