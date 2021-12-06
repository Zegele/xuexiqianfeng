package main

import "fmt"

func main() {
	/*
		参数传递：
			A：值传递：传递的是数据的副本。与原数据无干扰
				修改数据，对原始的数据没有影响。
			值类型的数据，默认都是值传递：基础类型，array, struct

			B：引用传递：传递的是数据的地址，导致多个变量指向同一块内存地址
				引用类型的数据，默认都是引用传递：slice， map， chan
	*/
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println("函数调用前，数组arr1的数据： ", arr1)
	fun1(arr1)
	fmt.Println("函数调用后，数组arr1的数据： ", arr1)

	fmt.Println("________________________")

	s1 := []int{1, 2, 3, 4}
	fmt.Println("函数调用前，切片的数据： ", s1)
	fun2(s1)
	fmt.Println("函数调用后，切片的数据： ", s1)
}

func fun1(arr2 [4]int) {
	fmt.Println("函数中，数组的数据： ", arr2)
	arr2[0] = 100
	fmt.Println("函数中，数组的数据修改后：", arr2)
}

func fun2(s2 []int) {
	fmt.Println("函数中，切片的数据：", s2)
	s2[0] = 100
	fmt.Println("函数中，切片的数据修改后： ", s2)
}
