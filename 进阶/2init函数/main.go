package main

import (
	"fmt"
	"xuexiqianfeng/进阶/2init函数/pk1"
)

func main() {
	fmt.Println("我是不是第一个？") //由于utils包中有init()函数，就先执行init()函数，所以该打印并不是第一个被执行。
	//utils.Count() //5. 如果不同包下的init函数，且相互倚赖，调用顺序为最后被倚赖的最先被初始化，
	pk1.MyTest1()
}

/*
1. 导入的包内如果有init函数，就先执行该包内的init函数。

2. 如果导入的包内有多个init函数，看util.go包文件

3. 如果同一个包下的，多个包文件都有init()函数，看util.go和test1.go两个文件，他们都是utils包下的。

4. 如果不同包下的init()函数，且不相互依赖，看utils包和pk1包。

5. 如果不同包下的init函数，且相互倚赖，调用顺序为最后被倚赖的最先被初始化，
倚赖最深的，要最先被初始化。
main -> A -> B -> C main倚赖A，A倚赖B，B倚赖C
则初始化顺序：C -> B -> A -> main
注意：
1. 要避免出现循环import，例如 A -> B ->C ->A
2. 一个包被其他多个包import，但只能被初始化一次。
*/
