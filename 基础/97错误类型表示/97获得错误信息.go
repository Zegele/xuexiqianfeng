package main

import (
	"fmt"
	"os"
)

func main() {
	/*
			想获得错误信息，得先知道错误信息的类型。

		怎样获得错误的类型？ 查看代码看该错误的函数是什么，返回值是什么。
			如os.Open的错误返回是

	*/
	f, err := os.Open("test.txt")
	if err != nil {
		fmt.Println(err)
		if ins, ok := err.(*os.PathError); ok {
			fmt.Println("1.0p:", ins.Op)
			fmt.Println("2.Path:", ins.Path)
			fmt.Println("3.Err:", ins.Err)
		}
		return
	}
	fmt.Println(f.Name(), "打开文件成功。。")
}
