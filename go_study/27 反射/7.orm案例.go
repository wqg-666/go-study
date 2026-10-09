package main

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// orm,对象关系映射，把结构体映射成数据库表，自动生成SQL语句，不用手写SQL语句，只负责造SQL
// qq-orm，用来标记结构体字段对应数据库的列名。
type ClassModel struct {
	Name string `qq-orm:"name"` // Name字段对应数据库列 name
	Id   int    `qq-orm:"id"`   // Id字段对应数据库列 id
}

// obj any：传入结构体对象（用来反射获取字段、标签、结构体名字）
// query ...any：可变参数，用来写 where 条件
func Find(obj any, query ...any) (sql string, err error) {
	//obj必须是结构体,reflect.TypeOf 获取传入对象的类型信息。
	t := reflect.TypeOf(obj)
	if t.Kind() != reflect.Struct {
		err = errors.New("obj必须是结构体")
		return
	}
	//处理查询条件 query（where 子句），where name = '三年一班';
	var where string
	//验证条件,?表示一个条件,>0说明有条件，要验证一下
	if len(query) > 0 {
		//Find(ClassModel{}, "id = ? and name = ?", 1, "三年一班")
		q := query[0]
		//验证第一个参数的类型，必须是字符串
		qs, ok := q.(string)
		//qs = "id = ? and name = ?"
		if !ok {
			err = errors.New("查询的第一个参数必须是字符串")
			return
		}
		//query 这个可变参数，第一个元素永远是 where 模板字符串，剩下所有元素都是用来填充 ? 的值。
		//算它问号的个数，值的个数 = len(query) - 1（减掉第一个模板字符串）
		if strings.Count(qs, "?")+1 != len(query) {
			err = errors.New("查询参数个数不匹配")
			return
		}
		//拼接where
		//query[1:] 是所有要填充到 ? 占位符的值，循环逐个取出；
		//query 是可变参数切片：第 0 号元素是模板字符串，query[1:] 就是从下标 1 开始往后所有元素，也就是填充值列表
		//query[1:] = [1, "三年一班"]
		for _, a := range query[1:] {
			switch s := a.(type) {
			case int:
				//最后的 1 代表：只替换找到的第 1 个问号，不是全部替换。
				//果 a 是 int 类型：
				//fmt.Sprintf("%d", s)：把数字转成字符串，比如1 → "1"
				qs = strings.Replace(qs, "?", fmt.Sprintf("%d", s), 1)
			case string:
				//fmt.Sprintf("'%s'", s)：给字符串加上单引号，数据库字符串常量需要单引号，比如 "三年一班" → '三年一班'
				qs = strings.Replace(qs, "?", fmt.Sprintf("'%s'", s), 1)
			}
		}
		//qs为id = 1 and name = '三年一班'"
		where = "where " + qs
	}
	//去拼接所有的有qq-orm的字段
	var columns []string //字符串切片
	//t 是结构体的反射类型对象。NumField() 返回这个结构体一共有多少个字段。
	//例子 ClassModel 有 2 个字段：Name、Id → t.NumField() = 2。
	for i := 0; i < t.NumField(); i++ {
		//.Tag.Get("qq-orm")：读取这个字段标签里 qq-orm 对应的值
		orm := t.Field(i).Tag.Get("qq-orm")
		//如果这个字段没有写 qq-orm 标签，orm就是空字符串。
		//continue：跳过本次循环，这个字段不加入查询列。
		if orm == "" {
			continue
		}
		//append 往切片columns追加元素。
		//循环结束后：columns = ["name", "id"]
		columns = append(columns, orm)
	}
	//生成表名，小写结构体的名字+s
	//t.Name()：获取结构体的名字。ClassModel 的名字就是 "ClassModel"
	//strings.ToLower(...)：全部转小写 → "classmodel" ，+ "s"：末尾加 s → "classmodels"
	name := strings.ToLower(t.Name()) + "s"
	sql = fmt.Sprintf("select %s from %s %s;",
		//strings.Join(columns, ",")：把切片 ["name","id"] 用逗号拼接成字符串 "name,id"
		strings.Join(columns, ","),
		name,  //classmodels
		where, //where id = 1 and name = '三年一班'
	)
	return
}
func main() {
	sql, err := Find(ClassModel{}, "name = ?", "三年一班")
	fmt.Println(sql, err)
	//sql语句，select name,id from class_models where name = '三年一班';
	sql, err = Find(ClassModel{}, "id = ? and name = ?", 1, "三年一班")
	fmt.Println(sql, err)
	//sql语句，select name,id from class_models where id = 1 and name = '三年一班';
	sql, err = Find(ClassModel{})
	fmt.Println(sql, err)
	//sql语句，select name,id from class_models;
	//select name,id from classmodels where name = '三年一班'; <nil>
	//select name,id from classmodels where id = 1 and name = '三年一班'; <nil>
	//select name,id from classmodels ; <nil>
}
