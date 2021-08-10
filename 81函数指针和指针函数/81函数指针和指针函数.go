package main

import "fmt"

func main() {
	/*
		函数指针：函数的指针。 一个指针，指向了一个函数的指针。
			因为go语言中，function，默认看作一个指针，没有*。 如打印函数，显示出来的就是一个内存地址。

			slice， map， function

		指针函数：一个函数，该函数的返回值是一个指针。
		注意是：返回值是指针的函数。
	*/

	var a func()
	a = fun1
	fmt.Println(fun1)
	fmt.Println(a)
	a() //调用函数

	arr1 := fun2()
	fmt.Printf("arr1的类型：%T，地址：%p，数值：%v\n", arr1, &arr1, arr1)

	arr2 := fun3()
	fmt.Printf("arr2的类型：%T，地址：%p，数值：%v，数值的地址：%p\n", arr2, &arr2, arr2, arr2) //Printf的时候会调用函数fun3,所以会打印arr的地址

}

func fun1() {
	fmt.Println("fun1()...")
}

func fun2() [4]int {
	arr := [4]int{1, 2, 3, 4}
	return arr
}

func fun3() *[4]int {
	arr := [4]int{5, 6, 7, 8}
	fmt.Printf("arr的地址：%p\n", &arr)
	return &arr //看80课，28行 p1 = &arr1 --> fmt.Println(p1)         //&[4]{1,2,3,4} 这样并不会打印出数组的内存地址
}
