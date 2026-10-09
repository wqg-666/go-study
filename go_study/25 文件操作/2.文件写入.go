package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	////不存在，创建一个文件并写入
	////三个参数：文件名、打开标志、文件权限。
	////os.O_CREATE：如果文件不存在，则新建这个文件；文件存在就直接打开
	////os.O_WRONLY：只写模式，打开后只能写，不能读
	////os.O_RDONLY：打开文件之后，只能读取文件内容，不能写入
	////os.O_RDWR：可读可写
	////文件权限 perm FileMode ，重点：这个权限参数在 Windows 系统下基本无效！只对 Linux/macOS 生效
	////由三个数字组成，代表所有者，组，其他用户的权限
	////R,读，八进制为4 W，写，八进制为2，X,执行，八进制为1
	////0444.三者均为只读
	////0666，均为读写，（4+2=6）
	////0777，均为读写执行（4+2+1=7）
	////0764，同理
	//file, err := os.OpenFile("E:\\go\\study1\\go_study\\25 文件操作\\w.txt", os.O_CREATE|os.O_RDONLY, 0644)
	//if err != nil {
	//	panic(err)
	//}
	//defer file.Close()
	//file.Write([]byte("你好")) //，写os.O_WRONLY
	////读
	//byteDate, err := io.ReadAll(file)
	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//fmt.Println(string(byteDate))
	//err := os.WriteFile("E:\\go\\study1\\go_study\\25 文件操作\\w1.txt", []byte("你好"), 0666)
	//
	////os.WriteFile
	////┌────────────────────────────┐
	////│ 打开/创建 w1.txt           │
	////│                            │
	////│ 写入 []byte("你好")        │
	////│                            │
	////│ 成功 → err = nil            │
	//fmt.Println(err) //输出nil
	//文件复制
	//读
	rFile, err := os.Open("C:\\Users\\HD\\Desktop\\11111111111111111111111111.jpg")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer rFile.Close()
	//写
	wFile, err := os.OpenFile("E:\\go\\study1\\go_study\\25 文件操作\\11.jpg", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer wFile.Close()
	//多个 defer：后进先出（栈），先执行 wFile.Close()，再执行 rFile.Close()
	io.Copy(wFile, rFile)
	//目录操作
	dir, err := os.ReadDir("E:\\go\\study1\\go_study\\25 文件操作")
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, entry := range dir {
		//entry.Info() 方法：获取这个文件 / 目录完整的文件信息对象
		//entry（DirEntry）：os.ReadDir 一次性读取目录时拿到的轻量信息 自带： entry.Name() entry.IsDir()
		//这两个不需要额外调用系统 API，性能好
		//但是 entry 没有 Size ()、ModTime ()！
		//想要拿到文件大小、修改时间，就必须调用 entry.Info()， 它会单独再向操作系统发起一次查询获取详细信息。
		info, _ := entry.Info()
		fmt.Println(entry.IsDir(), entry.Name(), info.Size())
	}
	//false 1.文件读取.go 2420
	//false 11.jpg 3145728
	//false 2.文件写入.go 2287
	//false hello 0
	//false hello.txt 130
	//false w.txt 6
	//false w1.txt 6
	//true → 是文件夹
	//false → 是普通文件，不是文件夹
}
