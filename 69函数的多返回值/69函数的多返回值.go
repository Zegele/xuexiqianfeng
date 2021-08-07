package main

import "fmt"

func main() {
	/*
		函数返回的结果，必须和函数定义的一致：类型，个数，顺序，一一对应。

		1. 将函数的结果返回给调用处。
		2. 同时结束了该函数的执行。

		空白标识符，专门用于舍弃数据： "_"
	*/
	peri1, area1 := rectangle(20.2, 10)
	fmt.Println("周长：", peri1, "面积：", area1)
}

//函数，用于求矩形的周长和面积
//return 方法1  这种好点，可以看到形参和返回值代表的东西。不用看整个函数体
func rectangle(len, wid float64) (perimeter, area float64) { //形参和返回参数，不用在函数体内再次声明
	perimeter = (len + wid) * 2
	area = len * wid
	return
}

//return 方法2
func rectangle2(len, wid float64) (float64, float64) {
	perimeter := (len * wid) * 2
	area := len * wid
	return perimeter, area
}
