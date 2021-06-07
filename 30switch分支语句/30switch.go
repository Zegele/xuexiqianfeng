package main

import "fmt"

func main() {
	/*
		switch语句：
		语法结构：
			switch 变量名{
			case 数值1： 分支1
			case 数值2： 分支2
			case 数值3： 分支3
			...
			default:
				最后一个分支
		}

		注意：
		1. switch可以作用在其他类型上，case后的数值必须和switch作用的变量类型一致。
		2. case是无序的。哪个匹配执行哪个。甚至default也可以放在前面，只是习惯上default放最后写。
		3. case后的数值是唯一的。
		4. default语句是可选的操作。可有可无，看具体是否需要。
	*/
	num := 5
	switch num {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
	case 3:
		fmt.Println("第三季度")
	case 4:
		fmt.Println("第四季度")
	default:
		fmt.Println("数据有误")
	}

	//模拟计算器
	num1 := 0
	num2 := 0

	oper := ""

	fmt.Println("请输入一个整数：")
	fmt.Scanln(&num1)
	fmt.Println("请输再入一个整数：")
	fmt.Scanln(&num2)
	fmt.Println("请输入一个操作： +， -， *， /")
	fmt.Scanln(&oper)

	switch oper {
	case "+":
		fmt.Printf("%d + %d = %d", num1, num2, num1+num2)
	case "-":
		fmt.Printf("%d - %d = %d", num1, num2, num1-num2)
	case "*":
		fmt.Printf("%d * %d = %d", num1, num2, num1*num2)
	case "/":
		fmt.Printf("%d / %d = %d", num1, num2, num1/num2)
	default:
		fmt.Println("我不会了。。。")
	}

}
