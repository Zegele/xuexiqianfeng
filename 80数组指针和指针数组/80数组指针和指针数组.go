package main

import "fmt"

func main() {

	/*
		数组指针：数组的指针。 首先是一个指针，一个数组的地址。
			*[4]Type  *[4]int

		指针数组：指针的数组。 首先是一个数组，存储的数据类型是指针。
			[4]*Type [4]*int

		*[5]float64， 指针，一个存储了5个浮点类型数据的数组的指针。
		*[3]string， 指针，数组的指针，数组存储了3个字符串
		[3]*string， 数组，存储了3个字符串的指针地址的数组
		[5]*float64， 数组，存储了5个浮点数据指针地址的数组
		*[5]*float64，指针，一个数组的指针，该数组存储了5个float64类型的数据的指针地址
		*[3]*string， 指针，存储了3个字符串指针地址数组的指针。
		**[4]string，指针，存储了4个字符串类型的数组的指针，的指针
		**[4]*string 指针 存储了4个字符串指针的数组，的指针，的指针。
	*/
	//1. 创建一个普通的数组
	arr1 := [4]int{1, 2, 3, 4}
	fmt.Println(arr1)
	//2. 创建一个指针，存储数组的地址 -->数组指针
	var p1 *[4]int
	p1 = &arr1
	fmt.Println(p1)         //&[4]{1,2,3,4} 这样并不会打印出数组的内存地址
	fmt.Printf("%p\n", p1)  //必须%p才能打印出数组的内存地址
	fmt.Printf("%p\n", &p1) //&p1是p1指针自己的内存地址

	//3. 根据数组指针，操作数组
	(*p1)[0] = 100 //上节课我们学到的写发，但是数组指针有另一种简写方法。
	fmt.Println(arr1)

	p1[1] = 200 //数组指针可以简化写法
	fmt.Println(arr1)

	//4.指针数组
	a := 1
	b := 2
	c := 3
	d := 4
	arr2 := [4]int{a, b, c, d}      //这是复制，以后修改arr2，跟a，b，c，d无关
	arr3 := [4]*int{&a, &b, &c, &d} //这是对应地址，如果修改就修改到根儿上了，涉及的都得跟着变。
	fmt.Println(arr2)               //[1,2,3,4]
	fmt.Println(arr3)               //[0xc00000a100 0xc00000a108 0xc00000a110 0xc00000a118] abcd对应的地址

	arr2[0] = 100
	fmt.Println("arr2的值是：", arr2)
	fmt.Println("a的值是：", a)

	*arr3[0] = 1000
	fmt.Println("arr3的值是：", arr3)
	fmt.Println("a的值是：", a)

	b = 2000
	fmt.Println(arr2)
	fmt.Println(arr3)
	for i := 0; i < len(arr3); i++ {
		fmt.Println(*arr3[i])
	}

}
