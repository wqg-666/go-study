package main

import (
	"fmt"
	"sync"
)

var sum1 int
var wait1 sync.WaitGroup
var lock sync.Mutex

// add1() 和 sub1() 各自只调用了一次 Lock()，所以各自只“进去”一次。
func add1() {
	lock.Lock() //加锁。其他尝试 Lock() 的 goroutine 会等待
	for i := 0; i < 100000; i++ {
		sum1++
	}
	lock.Unlock() //解锁。
	wait1.Done()
}
func sub1() {
	lock.Lock()
	for i := 0; i < 100000; i++ {
		sum1--
	}
	lock.Unlock() //解锁
	wait1.Done()
}
func main() {
	wait1.Add(2)
	go add1()
	go sub1()
	wait1.Wait()
	fmt.Println(sum1)
}
