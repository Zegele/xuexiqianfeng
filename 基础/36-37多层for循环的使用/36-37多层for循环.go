package main

import "fmt"

func main() {
	/*
		循环嵌套： 多层循环嵌套在一起

		题目一：
		*****
		*****
		*****
		*****
		*****

		Print()
		Printf()
		Println()

		题目二：
		1*1 = 1
		2*1 = 2 2*2= 4
		...
	*/
	//for i := 0; i < 5; i++ {
	//	fmt.Println("*****")
	//}
	for j := 0; j < 5; j++ {
		for i := 0; i < 5; i++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	for m := 1; m < 10; m++ {
		for n := 1; n < 10; n++ {
			if n > m {
				fmt.Println()
				break
			}
			fmt.Printf("%d * %d = %d\t", n, m, m*n)
		}
	}

	fmt.Println("................") //改良上
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", j, i, j*i)
		}
		fmt.Println()
	}
}
