package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		标准输入（键盘输入）和标准输出（显示器输出）
			fmt包：输入，输出

			输出：
				Print() //打印
				Printf() //格式化打印
					占位符：
			一般		%v 原样输出
					%T 打印类型
					%% 打印%

			bool	%t 打印布尔类型

			整型		%b 2进制
					%c 打印该数字对应的字符 // 如 65， 打印出来就是 A。 fmt.Printf("%c", 65) --> A
					%d 10进制
					%o 8进制
					%x 16进制，小写的
					%X 16进制，大写的
					%p 打印内存地址
					%q 单引号打印对应的字符		fmt.Printf("%q", 65) --> 'A'
					//对字符串，打印出来是双引号 //如 "world",打印出来是 world  fmt.Printf("%q", "world") --> "world"

			浮点数
					%f 打印浮点是
					%.3f 打印小数点后三位

			字符串
					%s 打印字符串或切片
					%q 双引号打印对应的字符串

				Println() // 打印之后换行

		标准输入：
			Scanln()  func Scanln
			Scanln is similar to Scan, but stops scanning aat a newline and after the final item there must be a newline or EOF.

			Scanf()

			bufio包

	*/
	a := 100            // int
	b := 3.14           // float64
	c := true           // bool
	d := "Hello World!" // string
	e := `你好，世界！`       //string
	f := 'A'            //字符

	fmt.Printf("%T,%b\n", a, a)
	fmt.Printf("%T,%.1f\n", b, b)
	fmt.Printf("%T,%t\n", c, c)
	fmt.Printf("%T,%s，%q\n", d, d, d)
	fmt.Printf("%T,%s,%q\n", e, e, e)
	//%q对字符串，打印是双引号的字符串 --> "Hello World!"
	fmt.Printf("%T,%d,%c，%q\n", f, f, f, f) //本质是个整数，可以当成数值，也可以当字符输出。
	// %d打出了10进制的数值--65，%c打印出了对应的字符-- A，%q对应整数，打印出由单引号包的该值-- 'A'
	fmt.Printf("%q\n", 65)
	fmt.Println(".....................")

	fmt.Printf("%v\n", a)
	fmt.Printf("%v\n", b)
	fmt.Printf("%v\n", c)
	fmt.Printf("%v\n", d)
	fmt.Printf("%v\n", e)
	fmt.Printf("%v\n", f)

	fmt.Println("...................")
	//Scanln 和 Scanf
	var x int
	var y float64
	fmt.Println("请输入一个整数，一个浮点数：")
	fmt.Scanln(&x, &y) //读取键盘的输入，通过操作地址，赋值给x和y   阻塞式的:等输入了程序才能继续。
	fmt.Printf("x的数值：%d, y的数值：%f\n", x, y)

	fmt.Scanf("%d,%f\n", &x, &y) //要按照这个格式输入 如： 1，1.33  格式里有英文逗号，一定要带上。
	fmt.Printf("x : %d, y : %f\n", x, y)

	//bufio
	fmt.Println("请输入一个字符串")
	reader := bufio.NewReader(os.Stdin) //Stdin 标准输入
	s1, _ := reader.ReadString('\n')    // _ 忽略了一个err
	//'\n'读到换行符就停止
	fmt.Println("读到的数据：", s1)

	fmt.Println("请输入一个字符串")
	reader2 := bufio.NewReader(os.Stdin) //Stdin 标准输入
	s2, _ := reader2.ReadString('e')     // _ 忽略了一个err
	//'e'读到e就停止
	fmt.Println("读到的数据：", s2)

}
