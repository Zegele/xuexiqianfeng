package main

import "fmt"

func main() {
	/*
			处理异常：panic() recover()
			处理错误：error
		panic:
		1. 内建函数
		2. 如果函数F中书写了panic语句，会终止其后要执行的代码，在panic所在函数F内如果存在要执行的defer函数列表，按照defer的逆序执行。
		3. 返回函数F的调用者G，在G中，调用函数F语句之后的代码不会执行，假如函数G中存在要执行的defer函数列表，按照defer的逆序执行，这里的defer有点 try-catch-finally中的finally。
		4. 直到goroutine整个退出，并报告错误。

		recover：
		1. 内建函数
		2. 用来控制一个goroutine的panicking行为，捕获panic，从而影响应用的行为
		3. 一般的调用建议
			a 在defer函数中，通过recover来终止一个goroutine的panicking过程，从而恢复正常代码的执行。
			b 可以获取通过panic传递的error

		简单来讲：go中可以抛出一个panic的异常，然后在defer中通过recover捕获这个异常，然后正常处理。


		以下场景适合panic（异常）：
		1. 空指针引用
		2. 下标越界
		3. 除数为0
		4. 不应该出现的分支，比如default
		5. 输入不应该引起函数错误。
	*/
	funA()
	defer myprint("defer main: 3....")
	funB() //如果funB()中有recover，则主程序会继续执行。panic只会影响到funB()
	defer myprint("defer mian: 4....")

	fmt.Println("main over")
}

func myprint(s string) {
	fmt.Println(s)
}

func funA() {
	fmt.Println("我是一个函数funA()...")
}

func funB() { //外围函数：函数内有defer，该函数对于defer来说，就是外围函数。
	//recover 看看：
	defer func() { //一个匿名函数
		if msg := recover(); msg != nil { //recover() 不需要传入参数,但是有返回值，返回值就是pannic的参数。
			//该例子中 recover()的值 就是 "funB函数，恐慌了"
			fmt.Println(msg, "--> 变成程序恢复了。。。")
		}
	}()
	//recover 看看：
	fmt.Println("我是函数funB()...")
	defer myprint("defer funB(): 1....")
	for i := 1; i <= 10; i++ {
		fmt.Println("i:", i)
		if i == 5 {
			//让程序中断
			panic("funB函数，恐慌了") //可以传入任何类型的数据。我们传了一点string，做为提示信息。
			//panic之后的不再执行，但已经defer过的要执行。 最后在调用处，返回panic的内容。（没有recover的情况下）
			//所以记得把recover放在panic（恐慌）之前的defer里，
		}
	}
	defer myprint("defer fun():2....")

}
