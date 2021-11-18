package main

func main() {
	/*
				go语言的最大两个亮点，一个式goroutine，一个就是chan了。二者合体的典型应用CSP，基本就是大家认可的并发开发神器，简化了并行程序的开发难度。
				一、CSP是什么
				CSP是Communicating Sequential Process的简称，通信顺序进程，是一种并发编程模型，是一个很强大的并发数据模型。用于描述两个独立的并发实体通过共享的通讯channel进行通信的并发模型。
				CSP中channel是第一类对象，它不关注发送消息的实体，而关注与发送消息时使用的channel。

				严格来说，CSP是一门形式语言，用于描述并发系统中的互动模式，也因此成为一众面向并发的编程语言的理论源头，衍生出go...
				go语言，其实只用到了CSP的很小一部分，即理论中的Process/Channel（对应到语言中的goroutine/channel）：这两个并发原语之间没有从属关系，
				Process可以订阅任意个channel，channel也并不关心是哪个process在利用它进行通信；process围绕Channel进行读写，形成一套有序阻塞和可预测的并发模型。

				二、Golang CSP
				与主流语言通过共享内存来进行并发控制方式不同。Go语言采用CSP模式。这是一种用于描述两个独立的并发实体通过共享的通讯（channel）进行通信的并发模型。
				Go就是借用CSP模型的一些概念为之实现并发进行理论支持，其实从实际上出发，go语言并没有，完全实现了CSP模型的所有理论，仅仅是借用了process和channel这两个概念。
				process在go语言上的表现就是goroutine是实际并发执行的实体，每个实体之间是通过channel通信来实现数据共享。

				Go语言的CSP模型是由协程Goroutine与通道Channel实现：
				1. Go协程goroutine：是一种轻量线程，它不是操作系统的线程，而是将一个操作系统线程分段使用，通过调度器实现协作式调度。、
				是一种绿色线程，微线程，它与Coroutine协程有区别，能够在发现堵塞后启动新的微线程。
				2. 通道channel：类似Unix的Pipe，用于协程之间通讯和同步。协程之间虽然解耦，但是他们和Channel有着耦合。
		·
	*/

}
