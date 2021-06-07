package main

import "fmt"

func main() {
	/*
			if语句的其他写法：
			if 初始化语句； 条件语句{
				执行语句
			}

		//注意：初始化语句的作用域只在if语句中。
	*/

	if num := 0; num > 0 {
		fmt.Println("正数")
	} else if num < 0 {
		fmt.Println("负数")
	} else {
		fmt.Println(num)
	}
}
