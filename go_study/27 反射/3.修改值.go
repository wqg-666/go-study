package main

import (
	"fmt"
	"reflect"
)

func setValue(obj any, value any) {
	v1 := reflect.ValueOf(obj) //获取值
	v2 := reflect.ValueOf(value)
	//v1是指针，要解引用，v1.Elem(),先取值，.Kind()再得到值的类型
	if v1.Elem().Kind() != v2.Kind() {
		return
	}
	switch v1.Elem().Kind() { //获取的类型是哪个，string,int,还是。。。
	case reflect.String:
		v1.Elem().SetString(value.(string))
	case reflect.Int:
		v1.Elem().SetInt(v2.Int())

	}
}
func main() {
	var name = "qq"
	var age = 24
	setValue(&name, "wqw")
	setValue(&age, 25)
	fmt.Println(name, age)
}
