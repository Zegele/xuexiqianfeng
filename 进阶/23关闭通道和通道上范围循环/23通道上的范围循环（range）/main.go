package main

import (
	"fmt"
	"time"
)

/*
	二、通道上的范围循环
	我们可以从循环通道上获取数据，知道通道关闭。for循环的for range形式可用于从通道接收值，直到它关闭为止。

*/
func main() {
	/*
		通过range访问通道
	*/
	ch1 := make(chan int)
	go sendData(ch1)
	//for循环的for range, 来访问通道
	for v := range ch1 { //range就直接从ch1中读取数据，与 v <- ch1 同样
		fmt.Println("读取数据", v)
	}
	fmt.Println("main over..")
}

func sendData(ch1 chan int) {
	for i := 0; i < 10; i++ {
		time.Sleep(1 * time.Second)
		ch1 <- i
	}
	close(ch1) //通知读取方，通道关闭
}
