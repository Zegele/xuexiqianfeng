package main

import "fmt"

func main() {
	/*
		defer的词义：“延迟”，”推迟“
		在go语言中，使用defer关键字来延迟一个函数或者方法的执行。

		1. defer函数或方法：一个函数或方法的执行被延迟了。

		2. defer的用法：
			A：对象.close() ，临时文件的删除 等等
				文件.open()
				defer close()
				读或写

			B： go语言中关于异常的处理，使用panic() 和recover()
				panic函数用于引发恐慌，导致程序中断执行
				recover函数用于恢复程序的执行，recover()语法上要求必须在defer中执行。

		3. 如果多个defer函数：先延迟的后执行，后延迟的先执行。

		4. defer函数传递参数的时候：defer函数调用时，就已经传递了阐述数据，只是暂时不执行函数中的代码而已。

		5. defer函数注意点：
			当外围函数中的语句正常执行完毕时，只有其中所有的延迟函数都执行完毕，外围函数才会真正的结束执行。
			当执行外围函数中的return语句时，只有其中所有的延迟函数都执行完毕后，外围函数才会真正返回。
			当外围函数中的代码引发运行恐慌时，只有其中所有的延迟函数的执行完毕后，该运行时恐慌才会正真被扩展值调用函数。
	*/
	defer fun1("hello") //后进先出
	fmt.Println("12345")
	defer fun1("world") //该fun1的外围函数调用结束后，才开始执行defer的fun1函数。
	// 该fun1的”外围函数“就是main函数。
	fmt.Println("什么")

	a := 2
	fmt.Println(a)
	defer fun2(a)
	a++ //相当于a = a+1
	fmt.Println("main中：", a)

	fmt.Println(fun3())
}

func fun1(s string) {
	fmt.Println(s)
}

func fun2(a int) {
	fmt.Println("fun2()函数中打印a：", a)
}

func fun3() int {
	fmt.Println("fun3()函数的执行。。。")
	defer fun1("haha")
	return 0
}
