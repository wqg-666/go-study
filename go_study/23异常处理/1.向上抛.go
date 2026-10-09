package main

import (
	"errors"
	"fmt"
)

// div() 发现错误 → 把错误返回给 server() →
// server() 发现错误 → 再把错误返回给 main() → main() 负责最终处理。
// 整个过程就是一个错误向上传递。
func div(a, b int) (res int, err error) {
	if b == 0 {
		err = errors.New("除数不能为0")
		return
	}
	res = a / b
	return
	//res → 计算结果
	//err → 错误信息
	//return 为什么没有写东西，func div(a, b int) (res int, err error) 这叫：命名返回值
	//也就是说函数已经提前给返回值起名字了：

}
func server() (res int, err error) {
	res, err = div(7, 5)
	if err != nil { //有没有错误，nil 可以理解成： 没有值 / 没有错误，err == nil 表示： 没有错误。
		//向上抛，server() 自己不处理这个错误，而是把错误继续返回给调用它的人。main函数
		return
	}

	return
}
func main() {
	res, err := server()
	if err != nil { //有错误
		fmt.Println(err)
		return
	}
	fmt.Println(res)
}
