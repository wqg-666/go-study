package main

import "fmt"

func main() {
	var nameList [3]string = [3]string{"张三", "李四", "王五"}
	//索引，下标从0开始
	fmt.Println(nameList)                  //输出整个数组的内容。[张三 李四 王五]
	fmt.Println(nameList[0])               //枫枫
	fmt.Println(nameList[len(nameList)-1]) //王五，最后以为
	//修改
	nameList[0] = "wqg"
	fmt.Println(nameList[0])
}
