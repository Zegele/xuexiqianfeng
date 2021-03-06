package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go func() {
		time.Sleep(1 * time.Second)
		ch1 <- 100
	}()
	select {
	case <-ch1:
		fmt.Println("case1可以执行")
	case <-ch2:
		fmt.Println("case2可以执行")
	case <-time.After(1 * time.Second):
		fmt.Println("case3可以执行")
		//default:
		//	fmt.Println("执行了default")
	}
}
