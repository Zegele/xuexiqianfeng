package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		水仙花数：是三位数，所以值在[100, 999]中
		每个位上的数字的立方和，放好等于该数字本身，那么这个数就叫水仙花数。
		例如： 153 -->  1*1*1+5*5*5+3*3*3=153
	*/
	a := 0
	b := 0
	c := 0
	for i := 100; i <= 999; i++ {
		a = i / 100        //百位
		b = (i % 100) / 10 //十位
		c = (i % 100) % 10 //个位  同样可以这样写：i % 10
		if a*a*a+b*b*b+c*c*c == i {
			fmt.Println(i)
		}
	}

	fmt.Println("..............")
	for i := 100; i < 1000; i++ {
		x := i / 100     //百位
		y := i / 10 % 10 //十位
		z := i % 10      //个位

		if math.Pow(float64(x), 3)+math.Pow(float64(y), 3)+math.Pow(float64(z), 3) == float64(i) {
			//math.Pow（float64，x） 表示给这个float64数字x次方
			fmt.Println(i)
		}
	}

	fmt.Println("..............")
	for x := 1; x < 10; x++ {
		for y := 0; y < 10; y++ {
			for z := 0; z < 10; z++ {
				if x*x*x+y*y*y+z*z*z == x*100+y*10+z {
					fmt.Println(x*100 + y*10 + z)
				}
			}
		}
	}
}
