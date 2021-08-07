package main

import "fmt"

func main() {
	/*
		go语言的数据类型：
			基本类型：
				int， float， bool， string

			复合类型：
				array， slice， map， function， pointer， struct， interface...

		函数的类型：
			func(参数列表的数据类型)（返回值列表的数据类型）

	*/

	a := 10
	fmt.Printf("%T\n", a) // int
	b := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", b) // array [4]int
	/*
		[4]string
		[6]string
	*/
	c := []int{1, 2, 3, 4}
	fmt.Printf("%T\n", c) // slice []int

	d := make(map[int]string)
	fmt.Printf("%T\n", d) // map[int]string
	/*
		map[string]string
		map[string]map[int]string
	*/

	fmt.Printf("%T\n", fun1)   // fun1 函数本身
	fmt.Printf("%T\n", fun1()) // fun1() 函数调用 有括号是函数调用

	fmt.Printf("%T\n", fun2) // fun2(int)int
	fmt.Printf("%T\n", fun3) // fun3(float64, int, int)(int, int)
	fmt.Printf("%T\n", fun4) // fun4(string, string, int, int)(string, int, float64)
}

func fun1() int {
	return 1
}
func fun2(a int) int {
	return 0
}

func fun3(a float64, b, c int) (int, int) {
	return 1, 1
}
func fun4(a, b string, c, d int) (string, int, float64) {
	return "a", 0, 0
}
