package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"` //转json后的内容
	Age  int    `json:"age"`
	//Age2     int    `json:"age2"`
	Age2     int    `json:"age2,omitempty"` //忽略0值
	Password string `json:"password"`
	//Password string `json:"-"` //-表示忽略该字段
}

func main() {
	user := User{Name: "qq", Age: 18, Password: "123456"}
	byteDate, _ := json.Marshal(user)
	fmt.Println(string(byteDate)) //{"Name":"qq","Age":18,"Password":"123456"}
	//{"name":"qq","age":18,"password":"123456"}，有`json:"name"`之后
	//有Age2,{"name":"qq","age":18,"age2":0,"password":"123456"}.user：=里没有写Age,默认为0
	//有`json:"age2,omitempty"`,{"name":"qq","age":18,"password":"123456"}
}
