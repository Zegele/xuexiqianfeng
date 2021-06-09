package main

import "fmt"

func main() {
	/*
		数组的遍历：依次访问数组中的元素
		方法一：arr[0],arr[1]...（不演示）

		方法二：通过for循环，配合下标
		for i := 0; i<cap(arr1); i++{
			arr[i]
		}

		方法三： 是用range
			range，范围
		不需要操作数组的下标，到达数组的末尾，自动结束 for range循环。
			每次都取一对数据，就是下标和对应的数值。
	*/
	//方法二
	arr1 := [4]int{1, 2, 3, 4}

	for i := 0; i < cap(arr1); i++ {
		fmt.Println(arr1[i])
	}

	fmt.Println("..................")
	for index, value := range arr1 {
		fmt.Printf("下标是：%d， 值是：%d。\n", index, value)
	}

	sum := 0
	for _, v := range arr1 { // _ 下划线舍弃不需要的量。
		sum += v
	}
	fmt.Println(sum)
}
