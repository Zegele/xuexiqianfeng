package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	/*
		bufio:高效io读写
			buffer缓存
			io: input/output

		将io包下的Reader，Write对象进行包装（封装），带缓存的包装，提高读写效率。

			ReadBytes()
			ReadString()
			ReadLine()

	*/

	fileName := "进阶/10bufio包 Read/应对.txt"
	file, err := os.Open(fileName) //os.Open 只读、
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	//创建bufio下的Reader对象

	//1.Read() ,高效读取
	//b1 := bufio.NewReader(file) // 这个b1就是缓存区？
	//p := make([]byte, 1024)
	//n1, err := b1.Read(p) //把缓冲区读到的，传给p切片（字节类型的切片）
	////n1是读到的数量
	//fmt.Println(n1)
	//fmt.Println(string(p[:n1]))

	//2.ReadLine() 读一行，此方法比较低级，不建议使用
	//data, flag, err := b1.ReadLine()
	//fmt.Println(flag)
	//fmt.Println(err)
	//fmt.Println(data)
	//fmt.Println(string(data))

	//3.ReadString(delim byte)(string, error) 指定读到哪个分隔符，返回值直接是个string类型
	//s1, err := b1.ReadString('\n') // 读取到换行符，也就是读取一行。注意： 使用单引号，单引号用于字节类型时？应该吧
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	//s1, err = b1.ReadString('\n') //继续使用，则可以继续读下一行（那么就可以使用for循环了）
	//fmt.Println(err)
	//fmt.Println(s1)
	//
	//for {
	//	s1, err := b1.ReadString('\n')
	//	if err == io.EOF {
	//		fmt.Println("读取完毕。。")
	//		break
	//	}
	//	fmt.Println(s1)
	//}

	//4. ReadByte() 读取一个字节  ReadBytes（） 读取多个字节
	//data, err := b1.ReadBytes('\n')
	//fmt.Println(err)
	//fmt.Println(string(data))

	//5. Scanner 读取键盘输入 (但只能读取到空格之前的内容)
	//s2 := ""
	//fmt.Scanln(&s2) //只读到键盘输入中的空格之前的内容。如键盘输入了Hello World，回车后，就只会读取到Hello
	//fmt.Println(s2)

	//os.Stdin（可以读取键盘输入的任何内容）
	b2 := bufio.NewReader(os.Stdin)
	s2, err := b2.ReadString('\n')
	fmt.Println(err)
	fmt.Println(s2)

}
