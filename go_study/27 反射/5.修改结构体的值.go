package main

import (
	"fmt"
	"reflect"
	"strings"
)

type User struct {
	Name1 string `big:"-"`
	Name2 string // 没有写big标签，big的值是空字符串 ""
}

func SetStruct(obj any) {
	// obj传进来是 &s，*User指针。 Elem()解指针
	v := reflect.ValueOf(obj).Elem()
	t := reflect.TypeOf(obj).Elem()
	for i := 0; i < v.NumField(); i++ {
		//NumField() = 结构体字段的数量，有两个，Name1,Name2
		value := v.Field(i)
		big := t.Field(i).Tag.Get("big") //tag 原始字符串：big:"-",获取这个字段 struct tag 里 big 对应的值"-"
		if big == "" {                   // 如果big标签为空字符串，continue，跳过处理。对应Name2的情况
			continue
		}
		// big不为空，把字段的值转大写，再设置回去
		value.SetString(strings.ToUpper(value.String()))
	}

}

func main() {
	s := User{Name1: "name1", Name2: "name2"}
	SetStruct(&s)
	fmt.Println(s) //{NAME1 name2}
}
