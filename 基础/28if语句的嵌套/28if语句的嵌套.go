package main

import "fmt"

func main() {
	/*
		if语句的嵌套
		if 	布尔表达式1 {
			// 条件表达式1为true时，执行
		}else if 布尔表达式2 {
			// 条件表达式2为true时，执行
			}else{
			// 条件表达式2为false时，执行
			}
		}
	*/
	sex := "泰国"

	if sex == "男" {
		fmt.Println("去男厕所")

	} else if sex == "女" {
		fmt.Println("去女测试")

	} else {
		fmt.Println("我也不知道了")
	}

}
