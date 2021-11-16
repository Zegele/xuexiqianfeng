package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		一、关闭通道
		发送者可以通过关闭信道，来通知接收方不会有更多的数据被发送到channel上。
		close(ch)

		接收者可以在接收来自通道的数据时使用额外的变量来检查通道是否已经关闭。
		语法结构：
		v, ok := <- ch
		类似map操作，存储key，value键值对
		v, ok := map[key] //根据key从map中获取value，如果key存在，v就是对应的数据，如果key不存在，v时默认值。

	*/
	ch1 := make(chan int)
	go sendData(ch1)

	//读取通道的数据
	for { //注意这是死循环，慎重！
		time.Sleep(1 * time.Second)
		v, ok := <-ch1 //主goroutine中读取数值
		//由于这个读取通道在前，所以它等着要读取数据，如果没有读到数据（没有写入通道，或通道没有关闭），可能会死锁，因为在一直等着。
		if !ok { //表示：如果是false
			fmt.Println("已经读取了所有的数据。。", v, ok)
			break
		}
		fmt.Println("读取的数据: ", v, ok)
	}
	fmt.Println("main over..")

}
func sendData(ch1 chan int) {
	//发送方：10条数据
	for i := 0; i < 10; i++ {
		ch1 <- i //将i写入到通道中
		//0 1 .. 9
		fmt.Println(i)
	}
	//close(ch1)
}
