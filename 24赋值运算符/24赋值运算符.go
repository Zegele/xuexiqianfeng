package main

import "fmt"

func main() {
	/*
			赋值运算符：
			=， +=， -=， *=， /=， %=， <<=, >>=, &=, |=, ^= ...

			= 把=右边的数值，赋值给左边的变量
			+=, a += b, 相当于 a = a + b

		运算优先级
		7： ~(什么？), !， ++， --
		6： *, /, %, <<, >>, &, &^
		5: +, -, ^
		4: ==, !=, <, <=, >, >=
		3: <- (通道操作)
		2: &&
		1: ||


	*/
	var a int
	a = 3
	fmt.Println(a)

	a += 4         //a = a + 4
	fmt.Println(a) //7
	a -= 3
	fmt.Println(a) //4
	a *= 2
	fmt.Println(a) //8
	a /= 3
	fmt.Println(a) //2
	a %= 1
	fmt.Println(a) //0
}
