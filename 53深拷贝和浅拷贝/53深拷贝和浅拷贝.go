package main

import "fmt"

func main() {
	/*
		深拷贝：拷贝的是数据本身。
			值类型的数据，默认都是深拷贝：array, int, string, float, bool

		浅拷贝：拷贝的是数据的地址。
			导致多个变量指向同一块内存。
			引用类型的数据，默认都是浅拷贝：slice, map

			因为切片是引用类型的数据，直接拷贝的是地址。

		copy()
		func copy(dst, src []TYpe) int // copy(目标（slice）, copy源（slice）) ---> 返回int型
	*/

	s1 := []int{1, 2, 3, 4}
	s2 := make([]int, 0) // len:0, cap:0
	//for i := 0; i < len(s1); i++ {
	//		s2 = append(s2, s1[i])
	//}
	s2 = append(s2, s1...)
	fmt.Println(s1)
	fmt.Println(s2)

	s1[0] = 100
	fmt.Println(s1)
	fmt.Println(s2)

	//copy()
	s3 := []int{7, 8, 9}
	fmt.Println(".........", s2)
	fmt.Println(s3)

	copy(s2, s3) //把s3的数据，依次拷贝给s2
	fmt.Println(".........", s2)
	fmt.Println(s3)

	copy(s3, s2[2:]) //可以拷贝一部分，目标切片有多长，就只能拷贝几个
	fmt.Println(".........", s2)
	fmt.Println(s3)

	copy(s3[1:], s2[2:]) //可以从某个元素开始进行copy
	fmt.Println(".........", s2)
	fmt.Println(s3)

	copy(s3[1:], s2[2:]) //可以从某个元素开始进行copy
	fmt.Println(".........", s2)
	fmt.Println(s3)
}
