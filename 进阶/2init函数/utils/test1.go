package utils

import "fmt"

func init() {
	fmt.Println("我是test1包下的init()函数，我是第几个？")
}

/*
在同一个package中（在同一个目录中（在同一个包文件夹下）），
utils 叫 包（package） utils包
test1.go 叫 包文件  test1.go包文件
打开文件 -> package utils 中的utils 叫 包名

如果多个包文件都有init()函数，则按照包文件名的顺序，从小到大，执行init()函数。
如：先执行 test1.go中的 init()，后执行util.go中的 init()函数。
*/
