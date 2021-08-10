package main

import "fmt"

func main() {
	/*
		数据类型：
			值类型：int, float, bool, string, array， struct

			引用类型：slice, map, function, pointer

		通过指针操作值类型的数据。
			new() //该函数可以创建任意类型的指针。
				//不是nil, 而是存储的是零值。
			//指向了新分配的类型的内存地址（内存空间），里面存储了零值。
	*/
	p1 := Person{"小王", 20, "男", "北京市"}
	fmt.Println(p1)
	fmt.Printf("%p,%T\n", &p1, p1) //0xc000060040,main.Person //main包下的结构体类型

	p2 := p1
	fmt.Println(p2)
	fmt.Printf("p2地址：%p,%T\n", &p2, p2)

	p2.name = "小李"
	fmt.Println(p1)
	fmt.Println(p2)

	//2. 定义结构体指针
	var pp1 *Person
	pp1 = &p2
	fmt.Printf("pp1中存储的地址：%p，%T\n", pp1, pp1) //pp1中存储的地址就是p2的地址，p2的地址对应的就是p2结构体的数据。
	fmt.Printf("pp1自身的地址：%p\n", &pp1)
	fmt.Println(pp1)  //&{小李 20 男 北京市}存储的地址对应的值
	fmt.Println(*pp1) //{小李 20 男 北京市}取存储的地址对应的值。

	//(*pp1).name = "..."
	pp1.name = "???" //可以不写 *
	fmt.Println(pp1)
	fmt.Println(p2)

	p3 := new(Person) //创建了一个Person结构体类型的指针，返回值是指针
	//也就是，p3是Person结构体类型的指针。
	p3 = &p1
	fmt.Println(p3)
	p3.name = "***"
	fmt.Println(p3)
	fmt.Println(p1)

	//使用内置函数new()，go语言中专门用于创建某种类型的指针的函数。
	pp2 := new(Person)
	fmt.Printf("%T\n", pp2)
	fmt.Println(pp2) //&{ 0  } 各项的零值
	//(*pp2).name = "Jerry" 一般用简写
	pp2.name = "Jerry"
	pp2.age = 20
	pp2.sex = "女"
	pp2.address = "南京"
	fmt.Println(pp2)

}

type Person struct {
	name    string
	age     int
	sex     string
	address string
}

/*
//关于new和make听视频 7：06
make用于内建类型（map slice channel）的内存分配。
new用于各种类型的内存分配。
内建函数new本质上说跟其他语言中的同名函数功能一样：new（T）分配了零值填充的T类型的内存空间，
并且返回其地址，即一个*T类型的值。
用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。
有一点非常重要：new返回指针。

内建函数make(T, args)与new()有着不同的功能，
make只能创建slice、map和channel，
并且返回一个有初始值(非零)的T类型，而不是*T。
本质来讲，导致这三个类型有所不同的原因是指向数据结构的引用在使用前必须被初始化。（可以参考54课map的nil，以及初始化）
例如：一个slice，是一个包含指向数据（内部array）的指针、长度和容量的三项描述符；在这些项目被初始化之前，slice为nil，不可直接使用。
对于slice，map，channel来说，make初始化了内部的数据结构，填充适当的值。 之后，便可以使用了。

make返回初始化后的值 ，不再是nil。
*/
