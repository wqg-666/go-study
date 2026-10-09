package main

import (
	"fmt"
	"reflect"
)

func getType(obj any) {
	//var a int = 10
	t := reflect.TypeOf(obj) //获取类型int
	switch t.Kind() {        //获取的类型是哪个，string,int,还是。。。
	case reflect.String:
		fmt.Println("String")
	case reflect.Int:
		fmt.Println("Int")
	case reflect.Struct:
		fmt.Println("Struct")
	}
}
func main() {
	getType(1)
	getType("232")
	getType(struct {
		Name string
	}{})
	//Int
	//String
	//Struct
}
