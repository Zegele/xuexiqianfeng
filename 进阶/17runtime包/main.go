package main

import (
	"fmt"
	"runtime"
	"time"
)

func init() {
	//获取逻辑cpu的数量
	fmt.Println("逻辑CPU的数量-->", runtime.NumCPU())

	//设置go程序执行的最大cpu数量：[0,256] (注意不要乱设置！)
	n := runtime.GOMAXPROCS(runtime.NumCPU()) //一般就是设置计算机本身的逻辑cpu的数量。且应该再Go程序执行前设置，而不是执行中再设置。所以这里放在init函数中。
	fmt.Println(n)
	//另外，go1.8版之前，还是要设置一下，可以更高效的利用cpu。
	//go1.8之后，默认让程序运行再多个核上，可以不用设置了。
}
func main() {
	/*

	 */
	//获取goroot目录
	fmt.Println("GOROOT-->", runtime.GOROOT())
	//获取操作系统
	fmt.Println("os/platform-->", runtime.GOOS)
	//获取逻辑cpu的数量
	//fmt.Println("逻辑CPU的数量-->", runtime.NumCPU())

	//Gosched: 让当前线程让出cpu，让其他线程运行，他不会挂起当前线程，因此当前线程未来会继续执行。
	/*
		这个函数的作用是让当然goroutine让出cpu,当一个goroutine发生阻塞，Go会自动地把与该goroutine处于同一系统线程的其他goroutine转移到另一个系统线程上去，以使这些goroutine不阻塞。
	*/
	/*
		go func() {
				for i := 0; i <= 5; i++ {
					fmt.Println("goroutine...")
				}
			}()

			for i := 0; i < 4; i++ {
				//让出时间片，先让别的goroutine执行//但是它本身还是会抢资源的。所以每次的运行的结果会可能不一样。
				runtime.Gosched()
				fmt.Println("main...")
			}
	*/

	//创建goroutine
	go func() {
		fmt.Println("goroutine开始。。")
		//调用fun
		fun() //fun()函数中，Goexit()使goroutine结束后,后面的不会再被打印。

		fmt.Println("goroutine结束。。")
	}()

	//睡一会儿
	time.Sleep(3 * time.Second)

}

func fun() {
	defer fmt.Println("defer...")
	//return //终止了函数，但defer还是执行的。
	runtime.Goexit() //终止当前的goroutine ，但defer还是执行的 （对比下和return的区别）
	fmt.Println("fun函数。。。")
}
