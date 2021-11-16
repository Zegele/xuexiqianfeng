package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*
		二、缓冲通道
		缓冲通道就是指一个通道，带有一个缓冲区。发送到一个缓冲通道只有在缓冲区满时才被阻塞。
		类似地，从缓冲通道接收的信息只有在缓冲区为空时才会被阻塞。
		可以通过将额外的容量参数传递给make函数来创建缓冲通道，该函数指定缓冲区的大小。

		语法：
		ch := make(chan type, capacity)
		上述语法的容量应该大于0，以便通道具有缓冲区。默认情况下，无缓冲通道的容量为0，因此在之前创建通道时省略了容量参数。

		非缓冲通道：make(chan T)
			一次发送，一次接收，都是堵塞的。

		缓冲通道：make(chan T capacity)
			发送：缓冲区的数据满了，才会阻塞。
			接收：缓冲区的数据空了，才会阻塞。
	*/

	ch1 := make(chan int)           //非缓冲通道
	fmt.Println(len(ch1), cap(ch1)) // 0 0
	//ch1 <- 100                      //阻塞式的，需要有其他的goroutine解除阻塞，否则deadlock（死锁）。

	ch2 := make(chan int, 5) //缓冲通道，缓冲区大小是5
	fmt.Println(len(ch2), cap(ch2))

	ch2 <- 100 // 1 5
	fmt.Println(len(ch2), cap(ch2))
	ch2 <- 200                      // 2 5
	ch2 <- 300                      // 3 5
	ch2 <- 400                      // 4 5
	ch2 <- 500                      // 5 5
	fmt.Println(len(ch2), cap(ch2)) // 5 5
	//ch2 <- 600                      //这里死锁，没有读取通道中的数据，死锁了。

	fmt.Println("-------------------")
	ch3 := make(chan string, 4)

	go sendData(ch3) //貌似要把写入数据放在前面

	for {
		v, ok := <-ch3
		if !ok {
			fmt.Println("读完了。。", ok)
			break
		}
		fmt.Println("\t读取的数据是：", v) //是按顺序的么？
	}

	fmt.Println("main over ..")
}
func sendData(ch chan string) {
	for i := 0; i < 10; i++ {
		ch <- "数据" + strconv.Itoa(i)
		fmt.Printf("子goroutine中写出第 %d 个数据\n", i)
	}
	close(ch)
}
