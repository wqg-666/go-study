set CGO_ENABLED=0
//关闭 CGO，编译出来的程序不依赖系统 C 库，静态编译，放到服务器直接运行。
set GOOS=linux
//设置目标操作系统为 Linux（打包 linux 下可执行程序）
set GOARCH=amd64
//目标 CPU 架构 amd64（x86_64）
go build -o main main.go
//编译 main.go，输出名为main（linux 程序，不带 exe 后缀）
set GOOS=windows
//改回 windows，恢复环境变量，防止后续编译搞错系统。