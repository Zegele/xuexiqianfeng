package main

import "fmt"

func main() {
	/*
		go语言支持函数式编程：
			1. 支持将一个函数作为另一个函数的参数。
			2. 支持将一个函数作为另一个函数的返回值。

		闭包（closure）：
			一个外层函数中，有内层函数，该内层函数中，会操作外层函数的局部变量（外层函数中的参数，或者外层函数中直接定义的变量），并且该外层函数的返回值就是这个内层函数。

			这个内层函数和外层函数的局部变量，统称为闭包结构。

			该局部变量的生命周期会发生改变，正常的局部变量随着函数调用而创建，随着函数的结束而销毁。
			但是！但是！但是！闭包结构中的外层函数的局部变量并不会随着外层函数的结束而销毁，因为内层函数还要继续使用。


	*/
	res1 := increment() //res1 = fun
	//res1调用了increment() ，所以res1是increment的返回值，而increment的返回值是一个函数体，所以res1 = fun函数整体（increment的返回值）
	fmt.Printf("%T\n", res1)
	fmt.Println("函数的值是:", res1)

	v1 := res1()
	fmt.Println(v1) // 1
	v2 := res1()
	fmt.Println(v2)     // 2
	fmt.Println(res1()) //每调用一次就会运算执行一次
	fmt.Println(res1())
	fmt.Println(res1())
	fmt.Println(res1())

	res2 := increment()
	fmt.Println("函数的值是:", res2) //函数的地址一样，但是res2创建的i的地址与res1的i的地址不一样了.
	v3 := res2()                //i不一样了，但是函数是一样的。 理解为子弹不一样了，枪是一样的。
	fmt.Println(v3)

}

func increment() func() int { //外层函数
	//1. 定义了一个局部变量
	i := 0

	//2. 定义了一个匿名函数，给变量自增并返回
	fun := func() int { //内层函数
		i++
		fmt.Printf("i的地址是：%p\n", &i) //取i的地址，查看i的地址
		return i
	}
	//3.返回该匿名函数
	return fun //把函数作为了返回值
}
