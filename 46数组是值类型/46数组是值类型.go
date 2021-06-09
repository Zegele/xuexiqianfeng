package main

import "fmt"

func main() {
	/*
		数据类型：
			基本类型： int, float, string, bool
			复合类型： array, slice, map, function, pointer, channel...

		数组的数据类型：
			[size]type

		值类型：理解为存储的数值本身
			将数据传递给其他变量，传递的是数据的副本（备份）
			int float bool string array(数组)
		引用类型：理解为存储的数据的内存地址
			slice map...
	*/

	// 1.数据类型
	num := 10
	fmt.Printf("%T\n", num)

	arr1 := [4]int{1, 2, 3, 4}
	fmt.Printf("%T\n", arr1)

	arr2 := [3]float64{1.11, 2.22, 3.33}
	fmt.Printf("%T\n", arr2)

	arr3 := [2]string{"1", "2"}
	fmt.Printf("%T\n", arr3)

	//2. 赋值
	num2 := num            //值传递
	fmt.Println(num, num2) // 10 10
	num2 = 20
	fmt.Println(num, num2) // 10 20

	//数组呢？
	arr4 := arr1
	fmt.Println(arr1)
	fmt.Println(arr4)

	arr4[0] = 100
	fmt.Println(arr1)
	fmt.Println(arr4)

	f := arr4 == arr1 //数组类型一致的情况下，比较数字对应下标位置的数值是否相等。
	fmt.Println(f)
}
