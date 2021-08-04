package main

import "fmt"

func main() {
	/*
		函数调用：
			1. 函数名：声明的函数名和调用的函数名要统一
			2. 实参必须严格匹配形参：顺序，个数，类型，一一对应。
	*/
	getAdd(1, 100)
	fun1(1, 1.2, "100")
}

func getAdd(a, b int) {
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)
}

func fun1(a, b float64, c string) {
	fmt.Printf("a:%.2f, b:%.2f, c:%s\n", a, b, c)
}
