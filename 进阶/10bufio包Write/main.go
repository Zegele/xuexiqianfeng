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

			func (b *Writer) Write(p []byte) (n int, err error)
			func (b *Writer) WriteByte(c byte) error
			func (b *Writer) WriteRune(r rune) (sizze int, err error)
			func (b *Writer) WriteString(s string) (int, error)

	*/
	fileName := "进阶/10bufio包 Read/cc.txt"
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	w1 := bufio.NewWriter(file)        // w1就是缓冲区。（这一步就是为file文件，创建了一个w1缓冲区）
	n, err := w1.WriteString("Ahello") //返回值时写出的数据量和返回的错误
	fmt.Println(err)
	fmt.Println(n) //写了5个数据，但是cc.txt还是空的。这就需要 Flush(),刷新缓冲区
	//那这5个数据写到哪里了？
	w1.Flush() //刷新缓冲区。 作用就是把写到缓冲区的内容写到目标文件中。

	//下面演示，目标文件比缓冲区的容量大一些。
	for i := 1; i <= 1000; i++ {
		w1.WriteString(fmt.Sprintf("%d:hello", i))
	}
	w1.Flush()
}
