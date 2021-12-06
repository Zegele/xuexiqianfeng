package main

import "fmt"

//全局变量的定义：
// num3 := 1000 //全局变量的定义不支持简单定义的写法。
var num3 = 1000

func main() {
	/*
		作用域： 变量可以使用的范围。
			局部变量： 函数内部定义的变量，就叫做局部变量。
				变量在哪里定义，就只能在哪个范围使用，超出这个范围，我们认为变量就被销毁了。


			全局变量： 函数外部定义的变量，就叫做全局变量。
				所有的函数都可以使用，而且共享这一份数据。

	*/

	//定义在main函数中，所以n的作用域就是main函数的范围内。
	n := 10
	fmt.Println(n)

	if a := 1; a <= 10 {
		fmt.Println(a) // 1
		fmt.Println(n) // 10
	}

	//fmt.Println(a) //不能访问a，已经出了作用域
	fmt.Println(n) // 10

	if b := 1; b <= 10 {
		n := 20
		fmt.Println(b)
		fmt.Println(n) //20 n是if中定义的n。 ctrl+b 或 ctrl+左键 后，可以查看光标停在那个n前，则说明是哪个n
	}

	fun1()
	fun2()
	fmt.Println(num3)
}

func fun1() {
	//fmt.Println(n) // 访问不到n
	num1 := 100
	fmt.Println("fun1()函数中： num1:", num1)
	fmt.Println("fun1()函数，访问全局变量：", num3)
}

func fun2() {
	num1 := 200
	fmt.Println("fun2()函数中： num1:", num1)
	num3 = 2000
	fmt.Println("fun2()函数，访问全局变量：", num3)
}
