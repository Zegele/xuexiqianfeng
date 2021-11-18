package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		time包中与通道相关的函数
		主要就是定时器，标准库中的Timer让用户可以定义自己的超时逻辑，尤其是在应对select处理多个channel的超时、单channel读写的超时等情形时尤为方便。

		Timeer是一次性的时间触发事件，这点与Ticker不同，Ticker是按一定事件间隔持续触发事件。

		Timer常见的创建方式：
		t := time.NewTimer(d)
			//func NewTimer(d Duration) *Timer 创建一个计时器，d时间以后触发。
		t := time.AfterFunc(d, f)
		t := time.After(d)
			//func After(d Duration) <-chan Time // 返回一个通道：chan，存储的是d使时间间隔之后的当前时间
			//在等待持续时间之后，在返回的通道上发送当前时间。相当于NewTimer(d).C

		虽然说创建方式不同，但是原理是相同的。
		Timer有3个要素：
		定时时间：就是那个d
		触发动作：就是那个f
		时间channel：也就是t.C
	*/
	//1. time.NewTimer(d)
	//1.1创建一个计时器
	/*
		timer := time.NewTimer(3 * time.Second)
		fmt.Printf("%T\n", timer) //*time.Timer
		fmt.Println(time.Now())

		ch2 := timer.C
		fmt.Println(<-ch2)

	*/

	//1.2让计时器中间停止
	/*
		timer2 := time.NewTimer(5 * time.Second)
		//开始goroutine，来处理出发后的事件

			go func() {
				<-timer2.C
				fmt.Println("Timer 2 结束了。。。")
			}()

			time.Sleep(3 * time.Second)

			flag := timer2.Stop()
			if flag { //如果flag为true
				fmt.Println("Timer2被停止了")
			}

	*/
	//3. time.After(d)
	ch := time.After(3 * time.Second)
	fmt.Printf("%T\n", ch) //<-chan time.Time
	fmt.Println(time.Now())

	time2 := <-ch
	fmt.Println(time2)
}
