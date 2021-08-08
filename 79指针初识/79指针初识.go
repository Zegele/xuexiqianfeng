package main

import "fmt"

func main() {
	/*
		指针：pointer
			存储另一个变量的内容地址的变量
	*/
	//1.定义一个int类型的变量
	a := 10
	fmt.Println("a的数值是：", a)     //10
	fmt.Printf("%T\n", a)        //int
	fmt.Printf("a的地址是：%p\n", &a) //0xc00000a0a8 地址是系统自动分配的

	//如何定义指针类型
	//声明指针，*T是指针变量的类型，它指向T类型的值。 如下：
	// var ip *int
	// var fp *float32

	// 2. 创建一个指针变量，用于存储变量a的地址。
	var p1 *int     //p1就是指针
	fmt.Println(p1) //<nil> 空指针 还没有存地址
	p1 = &a
	fmt.Println("p1的数值：", p1)                 //p1中存储的是a的地址。
	fmt.Printf("p1自己的地址：%p\n", &p1)           //指针p1自己的地址
	fmt.Println("p1的数值，是a的地址，该地址存储的数据：", *p1) //获取指针指向的变量的数值。

	//3. 操作变量，更改数值，并不会改变地址
	a = 100
	fmt.Println(a)
	fmt.Printf("%p\n", &a)
	fmt.Println(p1)

	//4. 通过指针，改变变量的数值
	*p1 = 200
	fmt.Println(a)

	//5. 指针的指针 套娃
	var p2 **int
	fmt.Println(p2)
	p2 = &p1
	fmt.Printf("%T,%T,%T\n", a, p1, p2)               //int,*int,**int
	fmt.Println("p2存储的数值：", p2)                       //其实是p1的地址
	fmt.Println("p2指向地址，对应的存储的数据：", *p2)              //其实就是p1存储的值，p1存储的就是a的地址
	fmt.Println("p2指向地址对应存储的数据，继续对应的存储的数据，套娃:", **p2) //就是a存储的值

}
