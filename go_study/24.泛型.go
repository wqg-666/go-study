package main

import (
	"encoding/json"
	"fmt"
)

//	func plus1(n1,n2 int) int {
//		return n1 + n2
//	}
//
//	func plusUint(n1,n2 uint) uint {
//		return n1 + n2
//	}
type Number interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64
}

func plus[T Number](n1, n2 T) T {
	return n1 + n2
}
func myPrint[T int, K string | int](n1 T, n2 K) {

}

type Response[T any] struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data T      `json:"data"`
}

// 可以理解成一个统一的返回数据格式
func main() {
	//plus1(1, 2)
	plus(1, 2)
	var u1, u2 = uint(2), uint(3)
	//plusUint(u1,u2)
	//plus1(int(u1), int(u2))
	fmt.Println(plus(u1, u2))
	type User struct { //用户信息
		Name string `json:"name"`
	}
	type UserInfo struct { //用户详情
		Name string `json:"name"`
		Age  int    `json:"age"`
	}
	//user := Response{
	//	Code: 0,
	//	Msg:  "ok",
	//	Data: User{
	//		Name: "qq",
	//	},
	//}
	//userInfo := Response{
	//	Code: 0,
	//	Msg:  "ok",
	//	Data: UserInfo{
	//		Name: "qq",
	//		Age:  1,
	//	},
	//}
	//byteDate1, _ := json.Marshal(userInfo)
	//byteDate, _ := json.Marshal(user)
	////json.Marshal()返回的是：[]byte也就是：字节切片。
	//fmt.Println(string(byteDate))
	////{"code":0,"msg":"ok","data":{"name":"qq"}}
	//fmt.Println(string(byteDate1))
	////{"code":0,"msg":"ok","data":{"name":"qq","age":1}}
	//反解
	var userResponse Response[User]
	json.Unmarshal([]byte(`{"code":0,"msg":"ok","data":{"name":"qq"}}`), &userResponse)
	fmt.Println(userResponse)           //{0 ok {qq}}
	fmt.Println(userResponse.Data.Name) //qq
	var userInfoResponse Response[UserInfo]
	json.Unmarshal([]byte(`{"code":0,"msg":"ok","data":{"name":"qq","age":1}}`), &userInfoResponse)
	fmt.Println(userInfoResponse.Data.Name, userInfoResponse.Data.Age) //qq 1
	//泛型切片
	type MySlice[T int | string] []T //切片里面的类型是T,底层是 []T → 切片！不是数组
	var mySlice = MySlice[int]{1, 2, 3}
	fmt.Println(mySlice[0] + 1) //2
	//泛型map.map的key只能是基本类型
	type MyMap[T string | int, K any] map[T]K
	var myMap = MyMap[string, string]{
		"name": "12",
	}
	fmt.Println(myMap) //map[name:12]
}
