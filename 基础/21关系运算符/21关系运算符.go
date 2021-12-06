package main

import "fmt"

func main() {
	/*
		关系运算符：
		1.2 关系运算符
			==, !=, >, <, >=, <=
			结果总是布尔类型的，true,false

	*/
	a := 3
	b := 5
	c := 3
	res1 := a > b
	res2 := b > c
	fmt.Printf("%T,%t\n", res1, res1) // %t是布尔类型的占位符
	fmt.Printf("%T,%t\n", res2, res2)

	res3 := a == b
	fmt.Println(res3)

	res4 := a == c
	fmt.Println(res4)

	fmt.Println(a != b, a != c)
}
