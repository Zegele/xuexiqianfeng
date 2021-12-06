package main

import "fmt"

func main() {
	/*
		for循环：某些代码会被多次的执行
		语法：
			for 表达式1（只执行一次，惯用于初始化）; 表达式2（循环的条件语句，bool类型）; 表达式3（循环体之后执行，惯用于条件的变化）{
				循环体
			}

	*/
	for i := 0; i < 5; i++ {
		fmt.Println("Hello World.")
	}

}
