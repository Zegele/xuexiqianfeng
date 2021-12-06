package main

import "fmt"

func main() {
	/*
		for循环练习：
		1. 打印58-23数字
		2. 求1-100的和
		3. 打印1-100内，能够被3整除，但是不能被5整除的数字，统计被打印的数字的个数。每行打印5个。
	*/
	for i := 58; i >= 23; i-- {
		fmt.Printf("%d,", i)
	}

	fmt.Println("")
	sum := 0
	for i := 1; i <= 100; i++ {
		sum = i + sum
	}
	fmt.Println(sum)

	fmt.Println("....................")
	count := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			count++
			if count%5 == 0 {
				fmt.Printf("%d\n", i)
			} else {
				fmt.Printf("%d,", i)
			}
		}
	}
	fmt.Println()
	fmt.Println(count)

	fmt.Println("....................")
	//改进第3题
	count1 := 0
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 != 0 {
			fmt.Print(i, "\t") // \t插入制表符
			count1++
			if count1%5 == 0 {
				fmt.Println("")
			}
		}
	}
	fmt.Println()
	fmt.Println(count)

}
