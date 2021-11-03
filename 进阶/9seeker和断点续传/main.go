package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	/*
			Seek(offseet int64, whence int) (int64, error), 设置指针光标的位置
				第一个参数：偏移量（距离初始位置的偏移量）
				第二个参数：设置初始位置
					0：SeekStart, 表示相对文件开始位置
					1: SeekCurrent, 表示光标当前位置
					2: SeekEnd， 表示文件末尾位置

			// Seek whence values
		const(
			SeekStart = 0 // seek relative to the origin of the file
			SeekCurrent = 1 // seek relative to the current offest
			SeekEnd = 2 // seek relative to the end
		)
	*/

	fileName := "D:/gogo/src/xuexiqianfeng/进阶/9seeker和断点续传/aa.txt"
	file, err := os.OpenFile(fileName, os.O_RDWR, os.ModePerm)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	//读写
	bs := []byte{1} //创建了一个字节，该字节只有一个字符 (字节类型的切片)
	file.Read(bs)
	fmt.Println(string(bs)) //a 	fmt.Println(bs[0]) // 97  fmt.Printf("%q\n", bs[0]) // 'a'

	file.Seek(4, io.SeekStart)
	//这时候光标在第四个字符处
	file.Read(bs)
	//再向后读一个字符（因为bs只有一个字符的容量），也就是读取了第五个字符
	fmt.Println(string(bs)) // e  所以打印出的就是第五个字符

	file.Seek(2, 0) // 0 相当于 io.SeekStart
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(3, io.SeekCurrent) // io.SeekCurrent 相当于 1
	file.Read(bs)
	fmt.Println(string(bs))

	file.Seek(0, io.SeekEnd) //表示光标在文件的结尾出，向后再移动0个位置
	file.WriteString("ABC")  //在光标处写

	file.Seek(0, io.SeekStart) //表示光标在文件的结尾出，向后再移动0个位置
	file.WriteString("ABC")    //在光标处写

}
func HandleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
