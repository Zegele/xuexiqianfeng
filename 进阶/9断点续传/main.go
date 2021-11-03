package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	/*
			断点续传：
				文件传递：文件复制
		思路：边复制，边记录复制的总量（临时文件，复制结束，即可删除临时文件）
	*/

	srcFile := "D:/gogo/src/xuexiqianfeng/进阶/9断点续传/123.txt"
	destFile := srcFile[strings.LastIndex(srcFile, "/")+1:] // +1 就不包含 '/'
	tempFile := destFile + "temp.txt"
	fmt.Println(tempFile)

	file1, err := os.Open(srcFile) //源文件
	HandleErr(err)

	file2, err := os.OpenFile(destFile, os.O_CREATE|os.O_WRONLY, os.ModePerm) //创建目标文件
	HandleErr(err)

	file3, err := os.OpenFile(tempFile, os.O_CREATE|os.O_RDWR, os.ModePerm) //创建临时文件
	HandleErr(err)

	defer file1.Close()
	defer file2.Close()

	//step1: 先读取临时文件中的数据，再seek 注意：该临时文件的内容只有数字
	file3.Seek(0, io.SeekStart)
	bs := make([]byte, 100, 100)
	n1, err := file3.Read(bs) //把 file3 读到 bs切片, n1直接表示读取的个数
	//HandleErr(err)
	countStr := string(bs[:n1])                      //转成字符串类型
	count, err := strconv.ParseInt(countStr, 10, 64) //转成数字类型，表示把countStr转成 10进制，64位数字
	//HandleErr(err)
	fmt.Println(count)

	//step2: 根据数值，设置读写的位置：
	file1.Seek(count, io.SeekStart)
	file2.Seek(count, io.SeekStart)
	data := make([]byte, 100, 100)
	n2 := -1            //读取的数据量
	n3 := -1            //写出的数据量
	total := int(count) // 读取的总量

	//step3: 复制文件
	for {
		n2, err = file1.Read(data)
		if err == io.EOF || n2 == 0 {
			fmt.Println("文件复制完毕。。", total)
			file3.Close()       //先关闭，再删除
			os.Remove(tempFile) //删除该临时文件
			break
		}
		n3, err = file2.Write(data[:n2])
		total += n3

		//将复制的总量，存储到临时文件中(要擦除临时文件原来的数据，教程貌似没这一步)
		//由于初始的临时文件是个空文件，所以教程把擦除数据这一步省略了？
		file3.Seek(0, io.SeekStart)
		file3.WriteString(strconv.Itoa(total)) //数字转成字符串

		fmt.Printf("total:%d\n", total)

		//假装断电
		//if total > 300 {
		//	panic("假装断电")
		//}
	}

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//可以处理的不足之处
//没有考虑错误处理
//复制的文件和临时文件生成的地址没有被指定。
