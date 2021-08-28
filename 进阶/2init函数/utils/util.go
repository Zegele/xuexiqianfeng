package utils

import "fmt"

func Count() {
	fmt.Println("utils包下的Count()函数。。。")
}

func init() {
	fmt.Println("utils包的init()函数，用于初始化一些信息。")
}
func init() {
	fmt.Println("utils包的另一个init()函数，我是第二个么？")
}

/*
在同一个包中，按照先后顺序调用init()函数。
*/
