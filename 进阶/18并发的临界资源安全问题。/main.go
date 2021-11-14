package main

import (
	"fmt"
	"math/rand"
	"time"
)

//全局变量，表示票数
var ticket = 10 //10张票

func main() {
	/*
		一、临界资源
		临界资源：并发环境中多个进程、线程、协程共享的资源。
		但是再并发编程中对临界资源的处理不当，往往会导致数据不一致的问题。

		三、临界资源安全问题的解决
		要想解决临界资源安全的问题，很多编程语言的解决方案都是同步。通过上锁的方式，某一时间段，只能允许一个goroutine来访问这个共享数据，当前goroutine访问完毕，解锁后，其他的goroutine才能来访问。

		我们可以借助sync包下的锁操作。

		四、写在最后
		在Go的并发编程中有一句很经典的话：不要以共享内存的方式去通信，而要以通信的方式去共享内存。
		在Go语言中并不鼓励用锁保护共享状态的方式在不同的Goroutine中分享信息（以共享内存的方式去通信）。
		而是鼓励通过Channel将共享状态或共享状态的变化在各个Goroutine之间传递（以通信的方式去共享内存），这样同样能像用锁一样保证在同一时间只有一个Goroutine访问共享状态。

		当然在主流的编程语言中为了保证多线程之间共享数据安全性和一致性，都会提供一套基本的同步工具集，如锁，条件变量，原子操作等等。
		GO语言标志库也毫不意外地提供了这些同步机制，使用方式也和其他语言差不多。


	*/

	//4个goroutine，模拟4个售票口
	go saleTickets("售票口1")
	go saleTickets("售票口2")
	go saleTickets("售票口3")
	go saleTickets("售票口4")
	time.Sleep(5 * time.Second)
	// 结果会出现-2票，就是因为临界资源的不安全问题。

	//
}

func saleTickets(name string) {
	rand.Seed(time.Now().UnixNano())
	for {
		if ticket > 0 {
			time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
			fmt.Println(name, "售出：第", ticket)
			ticket--
		} else {
			fmt.Println(name, "售罄，没有票了。。")
			break
		}
	}
}
