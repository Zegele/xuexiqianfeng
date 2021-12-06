package main

import (
	"fmt"
)

func main() {
	/*
		变量：variable
		概念：一小块内存，用于存储数据，在程序运行过程中数值可以改变

		使用：
			step1: 变量的声明，也叫定义
				第一种： var 变量名 数据类型
						变量名 = 赋值
					或	var 变量名 数据类型 = 赋值

				第二种：类型推断，省略数据类型
						var 变量名 = 赋值

				第三种： 简短声明，省略var
						变量名 := 赋值

			step2: 变量的访问，对变量赋值和取值
					直接根据变量名访问


		go的特性：
			静态语言：强类型语言（赋值的类型必须和声明的类型一致才能通过）
				go, java, 1包的使用++, 1包的使用# ..
			动态语言：弱类型语言（不知道是什么类型的语言，赋值后是什么类型的就是什么类型的）
				JavaScript， php，python， ruby。。
	*/
	// 第一种：定义变量，然后进行赋值
	var num1 int
	num1 = 30
	fmt.Printf("num1的数值是： %d\n", num1)
	//Printf是格式化打印，打印引号里的字符，%d是占位符，代表的是整数，打印的时候是num1替换%d的位置。
	//或 写在一行、
	var num2 int = 15
	fmt.Printf("num2的数值是： %d\n", num2)

	//第二种类型推断
	var name = "二狗" //类型推断
	fmt.Printf("类型是：%T，数值是：%s\n", name, name)

	//第三种，简短定义，也叫简短声明 简短声明其实也是类型推断。
	sum := 100
	fmt.Printf("%T\n", sum)
	fmt.Println(sum)

	//多个变量同时定义
	var a, b, c int
	a = 1
	b = 2
	c = 3
	fmt.Println(a, b, c)

	var m, n int = 100, 200
	fmt.Println(m, n)

	var n1, n2, n3 = 100, 3.14, "SHIME" //类型推断
	fmt.Println(n1, n2, n3)

	var (
		studentName = "狗蛋"
		age         = 18
		sex         = "男"
	)
	fmt.Printf("学生姓名：%s, 年龄：%d, 性别：%s\n", studentName, age, sex)
}
