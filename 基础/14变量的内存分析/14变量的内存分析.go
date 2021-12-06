package main

import "fmt"

/*
	注意点：
	1. 变量必须先定义才能使用
	//定义就是开辟内存，一定是先定义，然后赋值。
	2. 变量的类型和赋值的类型必须一致
	3. 变量的名字在同一个作用域中不能冲突
	4. 简短定义方式， 左边的变量至少有一个是新的。
	5. 简短定义不能定义全局变量。
	6. 变量的零值，也叫默认值。
	7. 变量定义了就要使用，否则无法通过编译。

*/
var a = 1000 //放在函数外声明的变量，就是全局变量。
var b = 2000

// 1包的使用 := 3000 //语法错误，在函数外不能使用简短定义变量的方法。

func main() {
	var num int
	num = 100
	fmt.Printf("num的数值是：%d, 地址是：%p\n", num, &num)

	num = 200
	fmt.Printf("num的数值是：%d, 地址是：%p\n", num, &num)

	num, name, sex := 1000, "狗蛋", "男" //简短定义，必须至少有一个新变量。
	fmt.Println(num, name, sex)

	fmt.Println(a)

	fmt.Println(".......................")

	var m int
	fmt.Println(m)
	var n float64
	fmt.Println(n)
	var s string
	fmt.Println(s)
	var s2 []int //切片
	fmt.Println(s2)
	fmt.Println(s2 == nil)
}
