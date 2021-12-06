package main

import "fmt"

func main() {
	/*
		函数的返回值：
			一个函数的执行结果，返回给函数的调用处，执行结果就叫做函数的返回值。

		return 语句：
			一个函数的定义上有返回值，那么函数中必须使用return语句，将结果返回给调用处。

	*/

	//1. 设计一个函数，用于求1-10的和，将结果在主函数中打印输出。

	res := getSum()
	fmt.Println("1-10的和： ", res)

	res2 := getSum2()
	fmt.Println(res2)
}

func getSum() int {
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	return sum
}

//return 数据的另一种写法
func getSum2() (sum int) { //返回值的名字是sum
	for i := 1; i <= 100; i++ {
		sum += i
	}
	return //这里可以不用写，因为上面已经写了，返回值是sum
	// return 1 // 也可以这样写， 相当于 返回的 sum 值是 1
}
