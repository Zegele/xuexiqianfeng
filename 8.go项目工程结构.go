1.go项目工程结构
1.1 gopath目录
gopath目录就是我们所编写源代码的目录。
该目录下有3个子目录：
src：里面每一个子目录，就是一个包，包内是go的源码文件
bin：编译后生产的，包的目标文件（go build就会产生）
pkg：生产可执行文件（好像go install就会生成）

方式一：使用go run命令

方式二：使用go build命令
	step1：打开终端：在任意文件路径下，运行：go install hello
			也可以进入项目的路径（文件夹下），然后运行：go instal

注意：在编译生产go程序的时候，go实际上回去两个地方找程序包：
GOROOT下的scr文件夹下，以及GOPATH下的src文件夹下。

在程序包里，自动找main包的main函数作为程序入口，然后进行编译。

	step2：运行go程序
	在.../go/bin/下，会发现出现了一个可执行文件，
	用如下命令运行，如：
	./hello.exe (windows下) 
	./hello (ubuntu下)

注意：src包下，如果同一个包下有多个源码文件,可以go run,
但不可编译，不可使用go install 和go build，

如src-hello-hello.go
			hello2.go
同在hello下的两个main文件（hello.go,hello2.go）,不可使用go build 和 go install。

