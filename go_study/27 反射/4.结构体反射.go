package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name"` //输出的内容中，用name代替Name
	Age   int
	IsMan bool
}

func ParseJson(obj any) {
	v := reflect.ValueOf(obj)
	t := reflect.TypeOf(obj)
	for i := 0; i < v.NumField(); i++ {
		tf := t.Field(i)
		jsonTag := tf.Tag.Get("json")
		if jsonTag == "" {
			jsonTag = tf.Name
		}
		fmt.Println(jsonTag, v.Field(i))
		//name qq
		//Age 18
		//IsMan true
	}

}
func main() {
	s := student{Name: "qq", Age: 18, IsMan: true}
	ParseJson(s)

}
