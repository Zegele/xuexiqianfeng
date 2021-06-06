package main

import "fmt"

func main() {
	/*
		字符串：
		1. 概念： 多个byte的集合，理解位一个字符序列
		2. 语法：使用双引号
			"abc", "hello", "A"
			也可以使用 `abc`
		3. 编码问题：
			计算机本质只识别0和1
			A:65， B:66 ...
			ASCII(美国标准信息交换码)

			中国的编码表：gbk, 兼容ASCII

			Unicode编码表：号称统一了全世界
			UTF-8， UTF-16， UTF-32
			GO语言使用UTF-8
		4. 转义字符
			4.1 有一些字符有特殊的作用，可以转义为普通的字符。
				\'xxx'\
				\" xxx "\
			4.2 有一些字符，就是一个普通的字符，转义后有特殊作用。
				\n 换行
				\t 制表符
	*/
	//1. 定义字符串
	var s1 string
	s1 = "狗蛋"
	fmt.Printf("%T,%s\n", s1, s1)

	s2 := `Hello World!`
	fmt.Printf("%T,%s\n", s2, s2)

	// 2. 双引号和单引号的区别 "A" 'A'
	v1 := 'A'
	v2 := "A"
	fmt.Printf("%T,%s, %d\n", v1, v1, v1)
	fmt.Printf("%T,%s\n", v2, v2)

	v3 := '中'
	fmt.Printf("%T, %d, %c, %q\n", v3, v3, v3, v3) //%c %q 可以显示单引号的内容

	// 3. 转义字符
	fmt.Println("\"Hello World\"") //打印带双引号的hello world
	// 在引号前加反斜杠

	fmt.Println("He`llowor'ld")
	fmt.Println(`he"llowor"ld`)
}
