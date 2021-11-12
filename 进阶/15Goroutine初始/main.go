package main

import (
	"fmt"
	"time"
)

func main() {
	/*
				一个goroutine打印数字，另外一个goroutine打印字母，观察运行结果

			Goroutine的规则：
			1. 当新的Goroutine开始时，Goroutine调用立即。与函数不同，go不等待Goroutine执行结束。当Goroutine调用，并且Goroutine的任何返回值被忽略之后，go立即执行到下一行代码。
			2. main的Goroutine应该为其他的Goroutine执行。如果main的Goroutine终止了，程序将被终止了，而其他Goroutine将不会运行。

		主goroutine
		封装main函数的goroutine称为主goroutine。
		主goroutine所做的事情并不是执行main函数那么简单。
		它首先要做的是：设定每个goroutine所能申请的栈空间的最大尺寸。
		在32位的计算机系统中此最大尺寸为250MB，而在64位的计算机系统中此尺寸位1GB。
		如果有某个goroutine的栈空间尺寸大于这个限制，那么运行时系统就会引发一个栈溢出（stack overflow）的运行时恐慌。
		随后，这个go程序的运行也会终止。
		此后，主goroutine会进行一系列的初始化工作，设计的工作内容大致如下：
		1. 创建一个特殊的defer语句，用于在主goroutine退出时做必要的善后处理。因为主goroutine也可能非正常的结束。
		2. 启动专用于在后台清扫内存垃圾的goroutine，并设置GC可用的标志。（GC是什么？）
		3. 执行main包中的init函数（如果没有就不执行）
		4. 执行main函数
		执行完main函数后，他还会检查主goroutine是否引发了运行时恐慌，并进行必要的处理。
		最后主goroutine会结束自己以及当前进程的运行。（这也是为什么主函数都结束了，怎么还有其他goroutine还在执行一小会儿的原因吧）

	*/

	//1. 先创建并启动子goroutine，执行printNumber()
	go printNum() //printNum()函数没有返回值，即使有返回值，在goroutine中也会被舍弃。

	//2. main中打印字母
	for i := 1; i <= 20; i++ {
		fmt.Printf("\t主goroutine中打印字母： A %d\n", i)
	}
	time.Sleep(1 * time.Second)
	fmt.Println("main over") //当main（主程序）结束了，即使子goroutine没有执行结束，也得结束。

}

func printNum() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("子goroutine中打印数字： %d\n", i)
	}
}
