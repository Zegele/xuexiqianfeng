package main

import "fmt"

func main() {
	/*
		写法1. 标准写法：33课

		写法2. 同时省略表达式1，表达式3
			for 表达式2 {

			}
		相当于while（条件）

		3. 同时省略3个表达式
			for {

			}
		相当于while(true)
		注意点：当for循环中，省略了表达式2，就相当于直接作用在了true上。

		4. 其他的写法： for循环中同时省略几个表达式都可以
			如果省略表达式1：往往需要把初始化，写在for循环外
			如果省略表达式2：循环永远成立-->死循环
			如果省略表达式3：往往把条件的变化写在循环内部。
	*/

	//写法2
	i := 0 //作用域在for循环外
	for i < 5 {
		fmt.Println("Hello!")
		i++
	}
	fmt.Println("-->", i)
}
