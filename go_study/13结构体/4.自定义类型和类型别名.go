package main

import "fmt"

type MyCode int        //自定义类型
type MyAliasCode = int //类型别名,不能绑定方法
func (m MyCode) Code() {

}

// func(m MyAliasCode) Code(){
//
// }报错
const myCode MyCode = 1
const myAliasCode MyAliasCode = 1

func main() {
	fmt.Printf("%v,%T\n", myCode, myCode)           //1,main.MyCode
	fmt.Printf("%v,%T\n", myAliasCode, myAliasCode) //1,int
	var age=1
	//if myCode==age{//myCode不是int类型，不可以与int类型的age比较
	//	fmt.Println("ok")
	//}
	if int(myCode)==age{//进行类型转换
		fmt.Println("ok")
	}
	if myCode==MyCode(age){//进行类型转换
		fmt.Println("ok")
	}
	if myAliasCode==age{
		fmt.Println("ok1")
	}
}
