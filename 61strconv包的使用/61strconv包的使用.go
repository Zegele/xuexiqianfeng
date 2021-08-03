package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		strconv包：字符串和基本类型之间的转换
			string convert
	*/
	// 1. bool类型
	s1 := "true"
	b1, err := strconv.ParseBool(s1) // 将“字符串”转化为“布尔”类型
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%t\n", b1, b1) //%t 布尔类型

	ss1 := strconv.FormatBool(b1) // 将“布尔”转化为“字符串”类型
	fmt.Printf("%T,%s\n", ss1, ss1)

	//2.整数
	s2 := "100"
	i2, err := strconv.ParseInt(s2, 10, 64) //参数1：要转化的对象；参数2：进制(如100按照10进制)；参数3：位数
	//ParseInt 该函数是把字符串转成10进制的整数
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("%T,%d\n", i2, i2)

	ss2 := strconv.FormatInt(i2, 10)
	fmt.Printf("%T, %s\n", ss2, ss2)

	//itoa(), atoi()
	i3, err := strconv.Atoi("-42") //string转为int类型
	fmt.Printf("%T,%d\n", i3, i3)

	ss3 := strconv.Itoa(-42)
	fmt.Printf("%T,%s\n", ss3, ss3)
}
