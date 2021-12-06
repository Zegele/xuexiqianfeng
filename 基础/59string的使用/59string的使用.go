package main

import "fmt"

func main() {
	/*
		Go中的字符串是一个字节的切片。
			可以通过将其内容封装在 "" 中来创建字符串。 Go中的字符串是Unicode兼容的，并且是UTF—8编码的。

		字符串是一些字节的集合。
			理解为一个字符的序列。
			每个字符都有固定的位置（索引，下标，index：从0开始，到长度减1）

		语法： "", ``(esc键下面的键)
			""
			"a", "b", "中"
			"abc", "hello"

		字符： --> 字符对应编码表中的编码值
			A --> 65
			B --> 66
			a --> 97

		这些数值叫做编码数，类型上叫字节。
		字节：byte --> uint8
	*/

	//1. 定义字符串
	s1 := "hello中国"
	s2 := `hello world`
	fmt.Println(s1)
	fmt.Println(s2)

	//2. 字符串的长度：返回的是字节的个数。一个英文字母是一个字节。一个中文一般是3个字节。
	fmt.Println(len(s1))
	fmt.Println(len(s2))

	//3. 获取某个字节
	fmt.Println(s2[0]) //获取字符串中的第一个字节
	a := 'h'
	fmt.Println(a) //获取字符串中的第一个字节
	b := 104
	fmt.Printf("%1包的使用, %1包的使用, %1包的使用\n", s2[0], a, b) //%c是字符的占位符

	for i := 0; i < len(s2); i++ {
		//fmt.Println(s2[i]) //输出的都是字节
		fmt.Printf("%1包的使用\t", s2[i])

	}

	//for range遍历字符串
	for i, v := range s2 {
		fmt.Println(i, v)
	}

	//5. 字符串是字节的集合。
	slice1 := []byte{65, 66, 67, 68, 69} //定义一个字节类型的切片。[]byte 类似 []int
	s3 := string(slice1)                 //把slice1转化为string类型。
	fmt.Println(s3)

	s4 := "abcde"
	slice2 := []byte(s4) //把s4转化为 []byte 字节类型的切片。
	fmt.Println(slice2)

	//6. 字符串是不能修改的
	fmt.Println(s4)
	//s4[2] = '8' //strings are immutable 字符串是不可变的。

}
