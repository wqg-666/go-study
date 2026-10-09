package main

import (
	"fmt"
	"os"
	"testing"
)

// 测试前执行
func setup() {
	fmt.Println("Before all tests")
}

// 测试后执行
func teardown() {
	fmt.Println("After all tests")
}
func TestAdd2(t *testing.T) {
	fmt.Println("测试中...")
	fmt.Println("你好")
}

// TestMain 是 Go 单元测试里包级别的测试入口函数。
func TestMain(m *testing.M) {
	fmt.Println("TestMain")
	setup()
	code := m.Run()
	//执行当前包下所有的测试用例（所有 TestXXX (t testing.T)），比如calc_test里的TestAdd, m.Run() 返回一个状态码
	teardown()
	os.Exit(code) //把测试的退出码交给操作系统。
}
