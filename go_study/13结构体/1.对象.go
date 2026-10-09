package main

import "fmt"

type Class struct {
	Name      string
	ClassName string
}

type Student struct {
	Class //Student对象组合了Class对象的东西
	Name  string
	Age   int
}

// Student类对应的方法
func (s Student) Study() {
	fmt.Printf("%s 正在学习\n", s.Name)
}
func (s Student) Info() {
	//Class和Student的名字都为Name,所以这里使用s.Class.Name
	fmt.Printf("名字：%s 班级：%s\n", s.Name, s.Class.Name)
	//Class和Student的名字不一样,所以这里直接使用s.ClassName。也可以用s.Class.ClassName
	fmt.Printf("名字：%s 班级：%s\n", s.Name, s.ClassName)
}

// 注意引用传递，不要值传递
func (s *Student) SetName(name string) {
	s.Name = name
	fmt.Printf("method in: %p\n", s)
}

// 值传递，改的是副本
func (s Student) SetName1(name string) {
	s.Name = name
	fmt.Printf("method in1: %p\n", &s)
}

func SetAge(info Student, age int) {
	info.Age = age
}

func SetAge1(info *Student, age int) {
	info.Age = age
}

func main() {
	c1 := Class{Name: "三年一班"}
	s1 := Student{Class: c1, Name: "qq"} //组合了Class
	s1.Study()                           //输出：qq 正在学习
	s1.Info()
	//名字：qq 班级：三年一班
	//名字：qq 班级：没有调用ClassName
	s2 := Student{Name: "张三"}
	s2.Study() //输出：张三 正在学习
	s2.Info()
	//名字：张三 班级：
	//名字：张三 班级：s2没有组合Class
	fmt.Printf("method out: %p\n", &s2)
	s2.SetName("李四")
	s2.Study()
	s2.SetName1("王五") //不能改名
	s2.Study()
	s := Student{
		Name: "枫枫",
		Age:  21,
	}
	fmt.Println(s.Age)
	SetAge(s, 18)//值拷贝，只改副本，原s不变
	fmt.Println(s.Age)//21
	SetAge1(&s, 17)//传指针，修改原结构体
	fmt.Println(s.Age) //17
}
