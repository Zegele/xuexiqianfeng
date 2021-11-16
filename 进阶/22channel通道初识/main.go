package main

import "fmt"

/*
	channel通道
	通道可以被认为是Goroutines通信的通道。
	类似于管道中的水从一端到另一端的流动，数据可以从过一段发送到另一端，通过通道接收。

	Go语言强烈建议使用Channel通道来实现Goroutine之间的通信。

	”不要通过共享内存来通信，而应该通过通信来共享内存“

	Go语言中，要传递某个数据给另一个goroutine（协程），可以吧这个数据封装成 一个对象，然后吧这个对象的指针传入某个channel中，
	另外一个goroutine从这个channel中读出这个指针，并处理其指向的内存对象。
	Go从语言层面保证同一个时间只有一个goroutine能够访问channel里面的数据，为开发者提供了一种优雅简单的工具，所以Go的做法就是使用channel来通信，
	通过通信来传递内存数据，使得内存数据在不同的goroutine中传递，而不是使用共享内存来通信。

	一、什么是通道
	1.1通道的概念
	通道是什么，通道就是goroutine之间的通道。它可以让goroutine之间相互通信。
	每个通道都有与其相关的类型。
	该类型是通道允许传输的数据类型。（通道的零值为nil。nil通道没有任何用处，因此通道必须使用类似于map和切片的方法来定义。）
	1.2通道的声明
	声明 一个通道和定义一个变量的语法一样：
	//声明通道
	var 通道名 chan 数据类型 // var a chan int
	//创建通道：如果通道为nil（就是不存在），就需要先创建通道
	通道名 = make(chan 数据类型) //a = make(chan int)

	也可以简短的声明，如下：
	a := make(chan int)
*/
func main() {
	var a chan int //声明通道
	fmt.Printf("%T,%v\n", a, a)
	if a == nil {
		fmt.Println("channel是nil的，不能使用，需要先创建通道")
		a = make(chan int)
		fmt.Println(a)
	}
	test1(a)
}
func test1(ch chan int) {
	fmt.Printf("%T,%v\n", ch, ch) //chan存的值是个地址
}
