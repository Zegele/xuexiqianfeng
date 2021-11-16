package main

import "fmt"

func main() {
	/*
		定向通道（单向通道）
			一、双向通道
			channel是用于goroutine之间的通信的。一个goroutine可以向通道中发送数据，另一个goroutine可以从该通道中获取数据。把既可以读取数据，又可以发送数据的通道，叫双向通道。
			chan T
				chan <- data, 发送数据，写出
				data <- chan, 获取数据，读取

			二、单向通道
			单向通道，只能发送或接收数据，是定向的。
				chan <- T, 只支持写
				<- chan T， 只读

	*/

	//双向通道实例：（注意这个done通道，就是为了不要让main结束太快）
	/*
			ch1 := make(chan string)
			done := make(chan bool)
			go sendData(ch1, done)
			data := <-ch1
			fmt.Println("子goroutine传来：", data)

			ch1 <- "我是mian"
			//time.Sleep(1 * time.Second)
			<-done
			fmt.Println("main over..")
		}
		func sendData(ch1 chan string, done chan bool) {
			ch1 <- "Hello"

			data := <-ch1
			fmt.Println("main goroutine传来：", data)
			done <- true
	*/

	// 单向通道
	ch1 := make(chan int) //双向 //一般是将双向的，某种情况下为了做数据保护，只用读取或写入其中之一。
	//ch2 := make(chan<- int) //单向，只能写，不能读 //设置这种通道意义不大
	//ch3 := make(<-chan int) //单向，只能读，不能写 //设置这种通道意义不大

	//ch1 <- 100
	//data := <- ch1

	//ch2 <- 100

	//data := <-ch3

	go fun1(ch1) //可读，可写 ，但这里只用了可写
	data := <-ch1
	fmt.Println("fun1函数中写出的数据：", data)

	go fun2(ch1) //但这里只用了可读
	ch1 <- 200
	fmt.Println("main over...")
}

//怎样是使用单向通道？实例：

//设置一个函数，该函数，只能操作只写的通道
func fun1(ch chan<- int) { //单向通道往往用于函数的参数
	//在函数内部，对于ch1通道，只能写数据，不能读取数据
	ch <- 100
	fmt.Println("fun1函数结束。。。")
}

//设置一个函数，该函数，只能操作只读的通道
func fun2(ch <-chan int) {
	data := <-ch
	fmt.Println("fun2函数，从ch中读取的数据是：", data)
}
