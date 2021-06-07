package main

import "fmt"

func main() {
	/*
		if 布尔表达式{
			//在表达式为true时，执行
			A段
		}else{
			//在表达式为false时，执行
			B段
		}
	*/
	//给定一个成绩，如果大于等于60，就是及格，否认则就是不及格。

	var score int
	fmt.Println(score)
	fmt.Println("请输入您的成绩：")
	fmt.Scanf("%d\n", &score)
	fmt.Println(score)
	if score >= 60 {
		fmt.Println("您的成绩是:", score, "，及格")
	} else {
		fmt.Printf("您的成绩是:%d， 不及格\n", score)

	}

	//给定性别，如果是男，就去男厕所，否则去女厕所。
	sex := "男" // bool, int, string自己设置

	if sex == "男" {
		fmt.Println("去男厕所")
	} else {
		fmt.Println("去女测所")
	}
	fmt.Println("Main over.")
}
