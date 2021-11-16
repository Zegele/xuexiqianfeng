package main

import (
	"fmt"
	"time"
)

/*
	2.3死锁
	使用通道时要考虑的一个重要因素是死锁。
	如果Goroutine在一个通道上发送数据，那么预计其他的goroutine应该接收数据。
	如果这种情况不发生，那么程序将在运行时出现死锁。

	类似地，如果goroutine正在等待从通道接收数据，那么另一些goroutine将会在该通道上写入数据，否则程序将会死锁。
*/
func main() {
	ch1 := make(chan int)
	go func() {
		fmt.Println("子goroutine开始执行。。。")

		data := <-ch1 //从ch1中读取数据
		fmt.Println("data:", data)
	}()
	time.Sleep(3 * time.Second)
	ch1 <- 10
	time.Sleep(3 * time.Second)
	fmt.Println("main over..")

	//死锁
	ch := make(chan int)
	ch <- 100 //阻塞
}
