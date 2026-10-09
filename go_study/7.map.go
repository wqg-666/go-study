package main

import "fmt"

func main() {
	//创建后要初始化
	var userMap map[int]string = map[int]string{
		1: "wqg",
		2: "张三",
		4: "李四",
		5: "",
	}
	fmt.Println(userMap)            //输出map[1:wqg 2:张三]
	fmt.Println(userMap[1])         //wqg
	fmt.Println(userMap[4])         //李四
	fmt.Println(userMap[3])         //输出空值，默认值
	fmt.Printf("%#v\n", userMap[3]) //输出""
	value := userMap[5]
	value, ok := userMap[5]
	fmt.Println(value, ok)                     //输出 true, 为空字符串，true表示5对应""，而不是默认为""
	fmt.Printf("value=%v, ok=%v\n", value, ok) //value=, ok=true
	//修改变量
	userMap[1] = "nihao"
	fmt.Println(userMap[1])
	//删除
	delete(userMap, 1)   //删除1对应的字符串
	fmt.Println(userMap) //map[2:张三 4:李四 5:]
	//会报错
	//var aMap map[string]string
	//aMap["name"] = "nihao"
	//fmt.Println(aMap)
	//初始化
	var aMap = map[string]string{}
	aMap["b"] = "a"   //赋值
	fmt.Println(aMap) //map[b:a]
	var aMap1 = make(map[string]string)
	aMap1["a"] = "c"
	fmt.Println(aMap1) //map[a:c]

}
