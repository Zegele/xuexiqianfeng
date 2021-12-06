package main

import "fmt"

func main() {
	/*
		切片Slice：
			1. 每一个切片引用了一个底层数组
			2. 切片本身不存储任何数据，都是这个底层数组存储，所以修改切片也就是修改这个底层数组中的数据。
			3. 当向切片中添加数据时，如果没有超过容量，直接添加，如果超过容量，自动扩容（成倍增长）
			4. 切片一旦扩容，就是重新指向一个新的底层数组。
	*/
	s1 := []int{1, 2, 3} //切片的地址指向了s1，所以s1其实就是一个内存地址
	fmt.Println(s1)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1)) //len:3, cap:3
	fmt.Printf("%p\n", s1)                           //这是切片地址
	//fmt.Printf("%p\n", &s1)                          //这是s1本身的地址

	s1 = append(s1, 4, 5)
	fmt.Println(s1)
	fmt.Printf("len:%d, cap:%d\n", len(s1), cap(s1)) //len:5, cap:6（成倍增加）
	fmt.Printf("%p\n", s1)

	s2 := make([]int, 6, 8)
	fmt.Println(s2)
	s2[0] = 1
	s2[1] = 2
	s2[2] = 3
	s2[3] = 4
	fmt.Println(s2)
	fmt.Printf("%p\n", s2)
	s2 = append(s2, 5, 6)
	fmt.Println(s2)
	fmt.Printf("%p\n", s2) //容量没有超，原地址没有变
}
