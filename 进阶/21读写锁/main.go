package main

import (
	"fmt"
	"sync"
	"time"
)

var rwMutex *sync.RWMutex //创建锁的类型
var wg *sync.WaitGroup    //创建同步等待组类型， 也可以是指针类型

func main() {
	/*
		通过对互斥锁的学习，我们已经知道了锁的概念以及用途。主要是用于处理并发中的临界资源问题。
		Go语言众阿哥的sync包提供了两种锁类型：sync.Mutex和sync.RWMutex。 其中RWMutex是基于Mutex实现的，只读锁的实现使用类型引用计数器的功能。

		RWMutex是读/写互斥锁。锁可以任意数量的读取器或单个编写器持有。RWMutex的零值是未锁定的mutex。

		如果一个goroutine持有一个rRWMutex进行读取，而另一个goroutine可能调用lock，那么在释放初始读取锁之前，任何goroutine都不应该期望能够获取读取锁。
		特别是，这禁止递归读取锁定。这是为了确保锁最终可用；被阻止的锁调用会将新的读卡器排除在获取锁之外。

		读写锁的写锁在解锁前只能上锁一次（一个），不能多个写锁同时上锁。（相对比，多个读锁是可以同时读上锁的）

		我们怎样理解读写锁呢？
		当有一个goroutine获得写锁定，其他无论是读锁定还是写锁定都将阻塞知道写解锁；
		当有一个goroutine获得读锁定，其他读锁定任然可以继续；
		当有一个或任意多个读锁定，写锁将等待所有读锁解锁之后才能进行写锁定。
		所以说这里的读锁定（RLock）目的其实是告诉写锁定：有很多人正在读取数据，你给我站一边去，等他们读（读解锁）完你再来写（写锁定）。可以总结为如下三条：
		1. 同时只能有一个goroutine能够获得写锁定。
		2. 同时可以有任意多个goroutine获得读锁定。
		3. 同时只能存在写锁定或读锁定（读和写互斥）。
		所以，RWMutex这个读写锁，该锁可以加多个读锁或者一个写锁，其经常用于读次数远远多于写次数的场景。
		读写锁的写锁只能锁定一次，解锁前不能多次锁定，读锁可以多次，但读解锁次数最多只能比读锁次数多一次（注意：只多一次），一般情况下我们不建议读解锁次数多于读锁次数。

		基本遵循两大原则：
		1. 可以随便读，多i个goroutine同时读
		2. 写的时候，啥也不能干。其他共routine不能读，也不能写，都只能等它解锁。

		读写锁即是针对于读写操作的互斥锁。它于普通的互斥锁最大的不同就是，它可以分别针对读操作和写操作进行锁定和解锁操作。
		读写锁遵循的访问控制规则与互斥锁有所不同。在读写锁管辖范围内，它允许任意个读操作的同时进行。但是在同一时刻，它只允许有一个写操作在进行。

		并且在某一个写操作被进行的过程中，读操作的进行也是不被允许的。
		也就是说读写锁控制下的多个写操作之间都是互斥的，并且写操作与读操作之间也都是互斥的。
		但是，多个读操作之间却不存在互斥关系。

	*/
	rwMutex = new(sync.RWMutex) //创建锁
	wg = new(sync.WaitGroup)    //创建同步等待组对象
	/*
		wg.Add(2)

		//读操作是可以多个同时读取的
		go readData(1)
		go readData(2)
	*/
	wg.Add(3)
	go writeData(1)
	go readData(2)
	go writeData(3)

	wg.Wait()
	fmt.Println("main over..")
}

func readData(i int) {
	defer wg.Done()

	fmt.Println(i, "开始读：read start..")

	rwMutex.RLock() //读操作上锁
	fmt.Println(i, "正在读取数据：reading...")
	time.Sleep(3 * time.Second)
	rwMutex.RUnlock() //读操作解锁
	fmt.Println(i, "读结束：read over...")
}

func writeData(i int) {
	defer wg.Done()
	fmt.Println(i, "开始写：write start。。")
	rwMutex.Lock() //写操作上锁
	fmt.Println(i, "正在写：writing。。")
	time.Sleep(3 * time.Second)
	rwMutex.Unlock()
	fmt.Println(i, "写结束：write over")
}
