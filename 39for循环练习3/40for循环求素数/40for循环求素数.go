package main

import (
	"fmt"
	"math"
)

func main() {
	/*
			打印2-100内的素数
		 	2/2 2/1
			3/3 3/2 3/1
			4/4 4/3 4/2 4/1
			5/5 5/4 5/3 5/2 5/1
	*/

	for i := 2; i <= 10; i++ {
		flag := true             //记录i是否是素数 //这里2是素数，整体设计里已经先入为主的判断了2是素数。
		for j := 2; j < i; j++ { //这个写法已经把1和数本身去除了，所以只要有一个可以整除，就不是素数。
			if i%j == 0 {
				flag = false //不是素数了
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

	fmt.Println("我的方法.................")
	for i := 2; i <= 10; i++ {
		count := 0 //count表示能被几个数整除，做计数
		//如果能被2个数整除，也就是1或本身，那这个数就是素数。
		for j := 1; j <= i; j++ {
			if i%j == 0 {
				count++
			}
		}
		if count == 2 { //这一步借鉴了老师的。
			fmt.Println(i)
		}
	}

	fmt.Println("我的方法.................")
	for i := 2; i <= 10; i++ {
		count := 0 //count表示能被几个数整除，做计数
		//如果能被2个数整除，也就是1或本身，那这个数就是素数。
		for j := 2; j < i; j++ {
			if i%j == 0 {
				count++
				break
			}
		}
		if count == 0 {
			fmt.Println(i)
		}
	}

	fmt.Println("老师的另一种方法.................")
	for i := 2; i <= 10; i++ {
		flag := true
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ { //判断到根号i，就可以了，不需要判断到i的前一个数。 缩减了计算量
			if i%j == 0 {
				flag = false //不是素数了
				break
			}
		}
		if flag {
			fmt.Println(i)
		}
	}

}
