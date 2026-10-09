package main

import (
	"fmt"
	"reflect"
)

func getValue(obj any) {
	//var a int = 10
	v := reflect.ValueOf(obj) //获取值10
	switch v.Kind() {         //获取的值的类型是哪个，string,int,还是。。。
	case reflect.String:
		fmt.Println("String", v.String())
	case reflect.Int:
		fmt.Println("Int", v.Int())
	case reflect.Struct:
		fmt.Println("Struct")
	}
}
func main() {
	getValue(1)
	getValue("232")
	getValue(struct {
		Name string
	}{Name: "qq"})
	//Int 1
	//String 232
	//Struct

}
