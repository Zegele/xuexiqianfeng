package main

import (
	"fmt"
	rand2 "math/rand"
	"time"
)

func main() {
	/*
		生成随机数random:
		伪随机数，根据一定的算法公式算出来的。
		math/rand
	*/
	num1 := rand2.Int() //rand有个默认的种子（seed），且种子数不会变,所以会生成固定的随机数
	fmt.Println(num1)

	for i := 0; i < 10; i++ {
		num := rand2.Intn(10) //随机数在0-9中选取，不包括10.
		fmt.Println("__:", num)
	}
	//同样由于种子不变，每次生成的随机数都一样的。

	rand2.Seed(1) //种子数是整数类型，设置为1,改变种子数会使随机数改变。
	num2 := rand2.Intn(10)
	fmt.Println("-->:", num2)

	t1 := time.Now()
	fmt.Println(t1)
	fmt.Printf("%T\n", t1)

	//时间戳：指定时间，距离1970年1月1日0时0分0秒，之间的插值：秒，纳秒
	timeStamp1 := t1.Unix() // 秒
	fmt.Println(timeStamp1)

	timeStamp2 := t1.UnixNano() //纳秒
	fmt.Println(timeStamp2)

	//step1: 设置种子数，可以设置成时间戳
	rand2.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		//step2: 调用生成随机数的函数
		fmt.Println("-->", rand2.Intn(100))
	}

	/*
		如何获取[15,76]之间的随机数？
				[0,61]+15
	*/

	fmt.Println("---->", rand2.Intn(62)+15)
}
