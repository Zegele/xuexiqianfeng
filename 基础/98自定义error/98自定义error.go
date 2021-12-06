package main

import (
	"fmt"
	"math"
)

func main() {
	/*

	 */
	radius := -3.0

	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err, "....") //这两个都调用的是系统的error接口，所以效果是一样的。err是err.Error()的简写？？？
		fmt.Println(err.Error(), "....")
		if err, ok := err.(*areaError); ok { //通过断言，可判断类型，以及获得结构体中的数据。
			fmt.Printf("半径是：%.2f\n", err.radius)
		}
		//	fmt.Println(err.fun1())
		return
	}
	fmt.Println("圆形的面积是：", area)

	fmt.Println("-----------------")

	area1, err1 := circleArea1(radius)
	if err1 != nil {
		fmt.Println(err1, "....") //这个直接用的系统error接口里的Error()方法 我们这里创建了方法，就用了我们的。
		fmt.Println(err1.Error(), "....")
		fmt.Println(err1.fun1())
		return
	}
	fmt.Println("圆形的面积是：", area1)

}

type error1 interface {
	Error() string
	fun1() float64
}

//1. 定义一个结构体，表示错误的类型
type areaError struct { //这个结构体如： 半径有误：-100
	msg    string  //错误信息  如：描述XXX出错了
	radius float64 // 错误的元素， 如：该运算中，只依靠半径，如果半径错了，那后面的计算都有问题。
}

//2. 实现error接口，就是实现Error()方法：
func (e *areaError) Error() string {
	return fmt.Sprintf("\"自己的\"error: 半径， %.2f, %s", e.radius, e.msg)
}

func (e *areaError) fun1() float64 {
	r := e.radius + 1
	return r
}

func circleArea(radius float64) (float64, error) { //这个error是error接口？  type error interface { Error() string }
	if radius < 0 {
		return -1, &areaError{msg: "半径是非法的", radius: radius}
	}
	return math.Pi * radius * radius, nil
}
func circleArea1(radius float64) (float64, error1) { //我创建了error1接口.并且给该接口增加了一个方法 fun1()float64
	if radius < 0 {
		return -1, &areaError{msg: "半径是非法的", radius: radius}
	}
	return math.Pi * radius * radius, nil
}
