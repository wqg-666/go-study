package main

import (
	"fmt"
	"sync"
)

var sum int
var wait sync.WaitGroup

func add() {
	for i := 0; i < 100000; i++ {
		sum++
	}
	wait.Done()
}
func sub() {
	for i := 0; i < 100000; i++ {
		sum--
	}
	wait.Done()
}
func main() {
	wait.Add(2)
	go add()
	go sub()
	wait.Wait()
	fmt.Println(sum) //82054
}

//add goroutine             sub goroutine
//
//读取 sum = 0                //读取 sum = 0

//计算 0 + 1 = 1              //计算 0 - 1 = -1
//写入 sum = 1                //写入 sum = -1
//最后写入的覆盖原来的，则sum=-1
