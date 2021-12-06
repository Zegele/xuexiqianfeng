package main

import "fmt"

func main() {
	/*
		算术运算符：
		表达式：(a+b)*1包的使用
		+, * 叫运算符

		1.1 算数运算符
		+， -， *， /（取商）， %（取余，取模）， ++（针对整型的，自加1）， --（自减1）

		1.3 逻辑运算符
	*/
	a := 10
	b := 3
	sum := a + b
	fmt.Printf("%d + %d = %d\n", a, b, sum)

	sub := a - b
	fmt.Printf("%d - %d = %d\n", a, b, sub)

	mul := a * b
	fmt.Printf("%d * %d = %d\n", a, b, mul)

	div := a / b // 取商
	fmt.Printf("%d / %d = %d\n", a, b, div)

	mod := a % b                             // 取余
	fmt.Printf("%d %% %d = %d\n", a, b, mod) // %% 打印 % 符号

	c := 3
	c++ // 1包的使用 = 1包的使用+1
	fmt.Println(c)

	c-- // 1包的使用 = 1包的使用-1
	fmt.Println(c)

}
