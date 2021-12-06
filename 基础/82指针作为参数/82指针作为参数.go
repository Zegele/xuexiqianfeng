package main

import "fmt"

func main() {
	/*
		指针作为参数：

		参数的传递：值传递，引用传递

	*/
	a := 10
	fmt.Println("fun1()函数调用前，a: ", a)
	fun1(a) //值传递：num = a = 10
	fmt.Println("fun1()函数调用后，a: ", a)

	fun2(&a) //传递的是a的地址，就是引用传递
	fmt.Println("fun2()函数调用后，a: ", a)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("fun3()函数调用前：", arr1)
	fun3(arr1)
	fmt.Println("fun3()函数调用后：", arr1)

	fmt.Println("fun4()函数调用前：", arr1)
	fun4(&arr1)
	fmt.Println("fun4()函数调用后：", arr1)

}

func fun1(num int) {
	fmt.Println("fun1()函数中，num的值：", num)
	num = 100
	fmt.Println("fun1()函数中修改num：", num)
}

func fun2(p1 *int) { //p1是指针，int类型的指针
	fmt.Println("fun2()函数中，p1的值：", *p1) //*p1, *指针，取指针对应的值。
	*p1 = 100
	fmt.Println("fun2()函数中修改p1：", *p1)
}

func fun3(arr [4]int) {
	fmt.Println("fun3()函数中数组：", arr)
	arr[0] = 100
	fmt.Println("fun3()函数中修改数组：", arr)
}

func fun4(arr *[4]int) {
	fmt.Println("fun4()函数中数组：", arr)  //&[1 2 3 4] 由于是数组的指针，可以不需要加*，也能显示数据。
	fmt.Println("fun4()函数中数组：", *arr) //[1 2 3 4]
	arr[0] = 100
	fmt.Println("fun4()函数中修改数组：", arr)
}
