package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片长度可变
	var nameList []string
	var nameList2 []string = []string{}
	nameList1 := []string{}
	nameList3 := make([]string, 0)
	var nameList4 []string
	nameList4 = make([]string, 0) //初始长度为0
	//声明之后要赋值，否则会报错
	nameList = append(nameList, "qq")
	nameList = append(nameList, "张三")
	//同一个nameList，只能接同类型的
	nameList1 = append(nameList1, "qw")
	nameList2 = append(nameList2, "e")
	nameList3 = append(nameList3, "i")
	nameList4 = append(nameList4, "op")
	fmt.Println(nameList[0]) //输出第一个元素
	//创建长度为3的切片。默认值为0
	ageList := make([]int, 3)
	fmt.Println(ageList)
	//定义一个长为3的数组
	array := [3]int{1, 2, 3}
	slices := array[:]
	fmt.Println(slices) //输出[1 2 3]
	slices1 := array[1:2]
	fmt.Println(slices1)    //输出[2]
	fmt.Println(array[0:1]) //输出[1]
	//切片排序
	var ints = []int{4, 8, 3, 2} //切片
	//var arr = [4]int{4, 8, 3, 2}数组
	sort.Ints(ints) //[2 3 4 8]升序
	fmt.Println(ints)
	sort.Sort(sort.Reverse(sort.IntSlice(ints))) //[8 4 3 2]降序
	fmt.Println(ints)

}
