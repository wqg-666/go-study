package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//绝对路径。一次性读取
	//byteData, err := os.ReadFile("E:\\go\\study1\\go_study\\25 文件操作\\hello.txt")
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteData))
	////分片读，文件内容很多
	//file, err := os.Open("E:\\go\\study1\\go_study\\25 文件操作\\hello.txt")
	//if err != nil { //有错误
	//	//  fmt.Println(err)
	//	//  return
	//	panic(err) //如果错误不为空，直接抛出这个错误，终止程序运行。
	//}
	//defer file.Close() //延迟执行，函数退出前才会执行这行代码,作用：关闭文件，释放操作系统资源
	//for {
	//	var byteData = make([]byte, 13) // 创建一个字节切片，长度 13，用来存放本次读到的文件内容
	//	n, err := file.Read(byteData)   //n表示读取的长度
	//	if err == io.EOF {              //读完了
	//		break //退出循环
	//	}
	//	fmt.Println(string(byteData[:n]))
	//	//byteData[:n]：截取切片，只保留本次读到的有效字节
	//}
	//带缓冲读,按行读
	//file1, err := os.Open("E:\\go\\study1\\go_study\\25 文件操作\\hello.txt")
	//if err != nil { //有错误
	//	//  fmt.Println(err)
	//	//  return
	//	panic(err) //如果错误不为空，直接抛出这个错误，终止程序运行。
	//}
	//defer file1.Close() //延迟执行，函数退出前才会执行这行代码,作用：关闭文件，释放操作系统资源
	//buf := bufio.NewReader(file1)
	////file1 是直接伸手去文件柜子拿东西；
	////bufio.Reader 相当于先拿一个托盘（缓冲区），一次性搬一堆放到托盘，之后从托盘取，不用反复开柜子
	////减少磁盘系统调用，提升读取速度。
	//for {
	//	line, _, err := buf.ReadLine()
	//	if err == io.EOF { //读完了
	//		break
	//	}
	//	fmt.Println(string(line))
	//}
	//指定分隔符
	file1, err := os.Open("E:\\go\\study1\\go_study\\25 文件操作\\hello.txt")
	if err != nil { //有错误
		//  fmt.Println(err)
		//  return
		panic(err) //如果错误不为空，直接抛出这个错误，终止程序运行。
	}
	defer file1.Close() //延迟执行，函数退出前才会执行这行代码,作用：关闭文件，释放操作系统资源
	buf := bufio.NewScanner(file1)
	buf.Split(bufio.ScanWords) //按照单词读
	var index int
	for buf.Scan() {
		index++
		fmt.Println(index, buf.Text())
	}
}
