//sync包 synchronization 同步
package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup //创建同步等待组的对象

func main() {
	/*
		WaitGroup: 同步等待组
			Add() 设置等待组中要执行的子 goroutine 的数量

			Wait() 让主goroutine处于等待状态

			Done() 让等待组中的counter计数器的值减1， 同Add(-1)一样
	*/
	wg.Add(2) //设置同步等待的goroutine的个数 即counter的数值是 2
	//设置这个计数器的值，一定要在调用goroutine程序之前设置，或增加Add的值，减少Add的值可以在任何地方。

	go fun1()
	go fun2()
	fmt.Println("main 进入阻塞状态。。等待wg中的子goroutine结束。。")
	wg.Wait() // 表示main goroutine进入等待，意味着阻塞
	fmt.Println("main..接触阻塞。。")
}
func fun1() {
	for i := 1; i < 10; i++ {
		fmt.Println("fun1()函数中打印。。A", i)
	}

	wg.Done() //给wg等待组中的counter数值减1，同  wg.Add(-1)

}

func fun2() {
	defer wg.Done()
	for j := 1; j < 10; j++ {
		fmt.Println("\tfun2()函数打印。。", j)
	}
}
