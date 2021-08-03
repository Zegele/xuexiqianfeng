package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		strings包下的关于字符串的函数

	*/
	s1 := "helloworld"
	//1. 是否包含指定的内容 --> bool
	fmt.Println(strings.Contains(s1, "h"))
	//strings.Contains 这个函数表示在 s1 中有没有 h 这个元素。

	//2. 是否包含chars中任意一个字符，即可 --> bool
	fmt.Println(strings.ContainsAny(s1, "abcd"))

	//3. 统计substr在s中出现的次数
	fmt.Println(strings.Count(s1, "l"))

	s2 := "20190525课堂笔记.txt"
	//4. 以xxx前缀开头，以xxx后缀结尾
	if strings.HasPrefix(s2, "2019") {
		fmt.Println("19年的文件")
	}
	if strings.HasSuffix(s2, ".txt") {
		fmt.Println("文本文档")
	}

	//5.某元素第一次出现的下标位置。 如果不存在返回 -1
	fmt.Println(strings.Index(s1, "l"))

	fmt.Println(strings.IndexAny(s1, "abcdh"))
	//任意元素在s1中第一次出现的位置。

	fmt.Println(strings.LastIndex(s1, "l"))
	// s1 中最后一次出现 l 的位置

	//6. 字符串的拼接
	ss1 := []string{"abc", "world", "hello"}
	s3 := strings.Join(ss1, "-")
	//将字符串之间以 - 符号 进行拼接。
	fmt.Println(s3)

	//7. 切分字符串
	s4 := "123,456,aaa,495,45"
	ss2 := strings.Split(s4, ",") //返回的值是切片类型
	fmt.Println(ss2)
	for i := 0; i < len(ss2); i++ {
		fmt.Println(ss2[i])
	}

	//8. 重复拼接
	s5 := strings.Repeat("hello", 5)
	fmt.Println(s5) //把hello重复拼接了5次

	//9. 替换
	s6 := strings.Replace(s1, "l", "*", 2)
	//把 l 替换成 * ，替换 2 个
	fmt.Println(s6)

	s6 = strings.Replace(s1, "l", "*", -1) //负数则是全部替换
	//把 l 替换成 * ，全都替换
	fmt.Println(s6)

	//10. 转大写或小写
	s7 := "hello WodLdj**123.。你好"
	fmt.Println(strings.ToLower(s7))
	fmt.Println(strings.ToUpper(s7))
	//只是转换英语字母的大小写，对其他字符无作用。

	/*
		截取子串：
		substing(start, end) --> substr
		str[start:end]  --> substr
			包含start，不包含end下标
	*/
	fmt.Println(s1)
	s8 := s1[0:5]
	fmt.Println(s8)
}
