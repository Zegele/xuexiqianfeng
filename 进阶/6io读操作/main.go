package main

import (
	"fmt"
	"io"
	"os"
)

/*
I/O操作也叫输入输出操作。 其中I是指input，O是指output，用于读或写数据的。
有些语言中也叫流操作，是指数据通信的通道。

Go语言标准库对IO的抽象非常精巧，各个组件可以随意组合，可以作为接口设计的典范。
*/

// 读取本地a.txt 文件中的数据
// step1: 打开文件
// step2: defer 关闭文件
// step3: 读取文件

func main() {
	//1.打开文件
	fileName := "D:/gogo/src/xuexiqianfeng/进阶/6io读操作/a.txt"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	//2. 关闭文件
	defer file.Close()

	//3. 关闭文件
	bs := make([]byte, 4, 4) //创建一个切片。 其作用是，将读取文件中的内容，存入该切片。

	/*
		// 第一次读取
		n, err := file.Read(bs) //file.Read(bs) Read函数
		fmt.Println(err)        //简单写错误。
		fmt.Println(n)          // 4 读取的个数
		fmt.Println(bs)         //[97 98 99 100]
		fmt.Printf("%c\n", bs)  //[a b c d]
		fmt.Println(string(bs)) //abcd

		// 第二次读取
		n, err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)          // 4
		fmt.Println(bs)         //[101 102 103 104]
		fmt.Println(string(bs)) //efgh

		// 第三次读取
		n, err = file.Read(bs)
		fmt.Println(err)
		fmt.Println(n)          // 2
		fmt.Println(bs)         //[105 106 103 104]
		fmt.Println(string(bs)) //ijgh

		// 第四次读取
		n, err = file.Read(bs)
		fmt.Println(err) //EOF
		fmt.Println(n)   // 0
		fmt.Println(bs)  //[105 106 103 104]

	*/
	n := -1 //用于表示返回的数量，先初始化为 -1 这样就没啥影响。
	// 如果初始值设置成100，就会困惑，是读取了100个字节，还是就只是初始值？
	for { //容易造成死循环
		n, err = file.Read(bs)
		if n == 0 || err == io.EOF {
			fmt.Println("读取到了文件的末尾，结束读取操作")
			break
		}
		fmt.Println(string(bs[:n]))
	}
}
