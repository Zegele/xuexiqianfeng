package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	/*
		写出数据：

	*/
	fileName := "D:/gogo/src/xuexiqianfeng/进阶/7io写操作/ab.txt"
	//step1: 打开文件
	//file, err := os.Open(fileName)//Open 是只读权限
	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, os.ModePerm) //Open 是只读权限
	//可读，可写 | 创建

	if err != nil {
		fmt.Println("err:", err)
		return
	}

	//step2: defer 关闭文件
	defer file.Close()

	//step3： 写出数据
	bs := []byte{65, 66, 67, 68, 69, 70} //A, B, C, D, E, F
	n, err := file.Write(bs)             //返回两个参数，第一个参数是写出了几个字符
	//Write() 是从文件的开头写。如果原来有内容，就会被覆盖。 在文件的打开方式上加append 则可以在文件原内容后增加数据。例如：os.O_RDWR|os.O_CREATE加上os.O_APPEND即可
	n, err = file.Write(bs[:2]) //把bs中的前两个字符写出去。
	fmt.Println(n)
	HandleErr(err)

	// 直接写出字符串
	m, err := file.WriteString("Hello!") //WriteString("字符串") 是直接写出字符串。 不像Write()只是写字节切片。
	fmt.Println(m)
	HandleErr(err)

	file.WriteString("\n")
	n, err = file.Write([]byte("today")) //把字符串类型，转为字节切片类型
	fmt.Println(n)
	HandleErr(err)

}

func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

/*拓展：
StringWriter
WriterAt
*/
