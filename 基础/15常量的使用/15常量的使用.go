package main

import "fmt"

func main() {
	/*
			常量：
			1. 概念： 常量是一个简单值的标识符，在程序运行时，不会被修改的量。
			2. 语法：
				显式类型定义
				隐式类型定义
			3. 常数：
				固定的数值： 100， "abc"

		注意：
		1. 常量只能是布尔型，数字型（整数型，浮点型和复数）和字符串型。
		2. 不曾使用的常量，在编译时，不会报错。
		3. 显式指定常量类型时（显式和隐式中的显式），需要保证赋值和指定的类型是一致的。

	*/
	fmt.Println(100)   //常数
	fmt.Println("abc") //常数

	// 1. 定义常量
	const PATH string = "www.baidu.com" //显示定义 有声明类型
	const PI = 3.14                     //隐式定义 没有声明类型
	fmt.Println(PATH)
	fmt.Println(PI)

	// 2. 定义一组常量
	//PATH = "www.mixin.com" // cannot assign to  PATH

	// 3. 定义一组常量
	const C1, C2, C3 = 100, 3.14, "haha"
	const (
		MALE   = 0
		FEMALE = 1
		UNKNOW = 3
	)
	fmt.Println(C1, C2, C3, MALE, FEMALE, UNKNOW)
	// 4.一组常量中，如果某个常量没有初始值，默认和上一行一致。
	const (
		a int    = 100
		b        // b 没有初始值，所以跟上一行的常量一致。
		c string = "Go"
	)
	fmt.Printf("%T,%d\n", a, a)
	fmt.Printf("%T,%d\n", b, b) //b和a一致
	fmt.Printf("%T,%s\n", c, c)

	// 5. 枚举类型： 使用常量组作为枚举类型。 一组相关数值的数据。
	const (
		SPRING = 0
		SUMMER = 1
		AUTUMN = 2
		WINTER = 3
	)
}
