package main

import "fmt"

func main() {
	/*
		iota: 特殊的常量，可以被编译器自动修改的常量
			每当定义一个const， iota的初始值为0
			每当定义一个常量，就会自动累加1
			知道下一个const出现，清零
	*/
	const (
		a = iota // 0
		b = iota // 1
		c = iota // 2
	)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	const (
		d = iota //清零，从 0 开始
		e        //默认和上面类型一样，iota值加1
	)
	fmt.Println(d)
	fmt.Println(e)

	//枚举中
	const (
		MALE = iota // 0
		FEMALE
		UNKNOW
	)
	fmt.Println(MALE, FEMALE, UNKNOW)
}
