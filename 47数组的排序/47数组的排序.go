package main

import "fmt"

func main() {
	/*
		数组的排序：
			让数组中的元素具有一定的顺序。

			arr := [5]int{15,23,8,10,7}
				升序：[7,8,10,15,23]
				降序：[23,15,10,8,7]

		排序算法：
			冒泡排序，插入排序，选择排序，希尔排序，堆排序，快速排序

		冒泡排序：（Bubble Sort）
			依次比较两个相邻的元素，如果他们的顺序（如从大到小），就把他们交换过来。
	*/

	arr := [3]int{1, 2, 3}

	for i := cap(arr); i > 1; i-- {
		for j := 0; j < cap(arr)-1; j++ {
			if arr[j] < arr[j+1] {
				temp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = temp
			}
		}
	}
	fmt.Println(arr)

	//老师的方法：
	for i := 1; i < len(arr); i++ {
		for j := 0; j < len(arr)-i; j++ { // j<len(arr)-i 这样可以减少对比次数。没必要重复比较。
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j] // 两个数互换用这种方法
			}
		}
	}
	fmt.Println(arr)

	//结合老师的，改造下我的
	for i := 1; i < cap(arr); i++ { //初始化减少了运算量
		for j := 0; j < cap(arr)-i; j++ { //第一轮对比2次，第二论，对比1次
			if arr[j] < arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	fmt.Println(arr)
}
