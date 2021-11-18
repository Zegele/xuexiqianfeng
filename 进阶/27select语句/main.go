package main

import (
	"fmt"
)

func main() {
	/*
			select是Go中的一个控制结构。select语句类似于switch语句，但是select会随机执行一个可运行的case。
			如果没有case可运行，它将阻塞，直到有case可运行。
			一、语法结构
			select语句的语法结构和switch语句很相似，也有case语句和default语句：
			slect{
				case communication clause :
					statement(s);
				case communication clause :
					statement(s);
				default://可选
					statement(s);
			}
			注意：
			1. 每个case都必须是一个通信（就是必须给通道发数据，或从通道获取数据）
			2. 所有channel表达式都会被求值。
			3. 所有被发送的表达式都会被求值。
			4. 如果有多个case都可以运行，select会随机公平地选出一个执行。其他不会执行。
			5. 否则：
				如果有default子句，则执行该语句。
				如果没有default子句，select将阻塞，直到某个通信可以运行；Go不会重新对channel或值进行求值。

		分支语句：if, switch, select
		select语句类似于switch语句，但是select语句会随机执行一个可运行的case。
		如果没有case可以运行，要看是否有default，如果有就执行default，否则就进入阻塞，直到有case可以运行。
	*/
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		//time.Sleep(3 * time.Second)
		ch2 <- 200
	}()
	go func() {
		//time.Sleep(3 * time.Second)
		ch1 <- 100
	}()
	select {
	case num1 := <-ch1:
		fmt.Println("ch1中获取的数据：", num1)
	case num2, ok := <-ch2:
		if ok {
			fmt.Println("ch2中获取的数据：", num2)
		} else {
			fmt.Println("ch2通道已经关闭")
		}
		//default:
		//	fmt.Println("default语句")
	}

	fmt.Println("main over")
}
