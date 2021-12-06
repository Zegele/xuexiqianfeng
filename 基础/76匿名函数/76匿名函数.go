package main

import "fmt"

func main() {
	/*
		匿名：没有名字
			匿名函数：没有名字的函数。

		定义一个匿名函数，直接进行调用，通常只能使用一次。
			也可以使用匿名函数赋值给某个函数变量，那么就可以调用多次了。

		匿名函数的作用：
			Go语言是支持函数式编程：
				1. 将匿名函数作为另一个函数的参数，回调函数。
				2. 将匿名函数作为另一个函数的返回值，可以形成闭包结构。

	*/

	fun1()
	fun2 := fun1
	fun2()

	//匿名函数
	func() {
		fmt.Println("我是一个匿名函数")
	}()

	fun3 := func() { //fun3就是这个函数体整体
		fmt.Println("我fun3也是一个匿名函数")
	}

	fun3() // 加括号就是调用这个函数
	//这样这个匿名函数就可以调用多次

	//定义带参数的匿名函数
	func(a, b int) {
		fmt.Println(a + b)
	}(1, 10)

	//定义带返回值的匿名函数
	res1 := func(a, b int) int {
		return a + b
	}(10, 30) //有这个括号，就是调用了函数，把返回值赋值给res1
	fmt.Println(res1)

	res2 := func(a, b int) int {
		return a + b
	} //没有这个括号，res2 就表示这个函数整体

	fmt.Println(res2(10, 100))

}

func fun1() {
	fmt.Println("我是fun1()函数")
}
