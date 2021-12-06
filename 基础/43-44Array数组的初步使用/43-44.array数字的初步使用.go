package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型：整数，浮点，布尔，字符串
			复合类型：array(数组)，slice（切片），map，struct，pointer，function，channel。。。

		1. 数组：
			概念：存储一组相同数据类型的数据结构
					理解为容器，存储一组数据
			语法：var 数组名 [长度] 数据类型
				 var 数组名 = [长度] 数据类型{元素1, 元素2...}
				 数组名 := [...]数据类型{元素。。。}

			通过下标访问：
				下标，也叫索引：index
				默认从0开始的整数，知道长度减一
				数组名[index]
					赋值
					取值

				不能越界：只能[0,长度-1]

			长度和容量：go语言的内置函数
				len(array\map\slice\string) 长度
				cap() 容量
	*/
	//step1: 创建数组
	var arr1 [4]int //创建了一个名为arr1的，类型为int型的，空数组。且有4个容量。
	fmt.Println(arr1)
	fmt.Printf("%p\n", &arr1)
	//step2: 数组的访问
	arr1[0] = 1
	arr1[1] = 2
	arr1[2] = 3
	arr1[3] = 4
	fmt.Println(arr1[0])
	fmt.Println(arr1[2])
	//fmt.Println(arr1[5])//invalid array index 5 (out of bounds for 4-element array)

	fmt.Println("数组的长度：", len(arr1)) //容器中实际存储的数据量
	fmt.Println("数组的长度：", cap(arr1)) //容器中能存储的最大数据量

	//一旦生成了数组，就不能修改数组的长度和容量，只能修改数组的数据。
	arr1[0] = 100
	fmt.Println(arr1)

	//数组的其他创建方式
	var a = [4]int{} // 同var a [4]int //有等号就要有初始化的值。哪怕只带括号，就是默认的0值。
	fmt.Println(a)

	var b = [4]int{1, 2, 3, 4}
	fmt.Println(b)

	var c = [5]int{1, 2, 4}
	fmt.Println(c)

	var d = [5]int{1: 1, 3: 2} //角标：数据
	fmt.Println(d)

	var e = [5]string{}
	fmt.Println(e)

	f := [...]int{1, 2, 3, 4, 5} //会根据后面的数据，自动判断容量
	fmt.Println(f)
	fmt.Println(len(f), cap(f))

	g := [...]int{1: 3, 6: 5}
	fmt.Println(g)
	fmt.Println(len(g), cap(g))
}
