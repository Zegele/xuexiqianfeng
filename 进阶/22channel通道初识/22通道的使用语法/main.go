package main

import "fmt"

func main() {
	/*
			发送和接收语法：
			data := <- a // read from channel a
			a <- data // write to channel a
			在通道上箭头的方向指定数据是发送还是接收
			另外：
			v, ok := <- a //从一个channel中读取

		1.4 通道的注意点
		Channel通道在使用的时候，有以下几点注意点：
		1. 用于goroutine，传递消息
		2. 通道，每个都有相关联的数据类型,
			nil chan, 不能使用，类似nil map，不能直接存储键值对
		3. 使用通道传递数据：<-
			chan <- data, 发送数据到通道。向通道中写数据
			data <- chan， 从通道中获取数据。从通道中读数据
		4. 阻塞：
			发送数据：chan <- data,阻塞的，直到另一条goroutine，读取数据来解除阻塞。
			读取数据：data <- chan, 也是阻塞的。一道另一条goroutine，写出数据解除阻塞。
		5. 本身channel就是同步的，意味着同一时间，只能有一条goroutine来操作。

		最后：通道是goroutine之间的连接，所以通道的发送和接收必须处在不同的goroutine中。

	*/
	var ch1 chan bool
	ch1 = make(chan bool)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println("子goroutine中，i: ", i)
		}
		//循环结束后，向通道中写数据，表示要结束了。。。
		ch1 <- true //向通道中写数据
		fmt.Println("结束。。")
	}()

	data := <-ch1 //从通道中读取数据
	fmt.Println("main ... data -->", data)
	fmt.Println("main..over..")
}
