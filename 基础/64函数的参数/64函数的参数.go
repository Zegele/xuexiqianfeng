package main

import "fmt"

func main() {
	/*
		函数的参数：
			形式参数：也叫形参，函数定义的时候，用于接收外部传入的数据的变量。
				函数中，某些变量的数值无法确定，需要由外部传入数据。

			实际参数：也叫实参，函数调用的时候，给形参复制的实际数据
	*/
	getSum(9, 10)
}

func getSum(x, y int) { //参数不用再声明
	sum := 0
	a := x
	for a <= y {
		sum += a
		a++
	}
	fmt.Printf("%d 到 %d 的和是： %d\n", x, y, sum)
}
