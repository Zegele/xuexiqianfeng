package main

import "fmt"

func main() {
	/*
		接口的嵌套：

	*/
	var cat Cat = Cat{}
	cat.test1()
	cat.test2()
	cat.test3()

	var a A
	a = cat //把cat作为A接口的类型，就只能实现test1方法。
	// 简写成： var a A = cat

	a.test1() //A接口的test1方法

	var b B
	b = cat   //把cat作为B接口的类型，就只能实现test2方法。
	b.test2() //B接口的test2方法

	var c C
	c = cat   //把cat作为C接口的类型，就能实现test1，test2，test3方法。
	c.test1() //C接口继承了A接口的test1，B接口的test2
	c.test2() //C接口继承了A接口的test1，B接口的test2
	c.test3() //C接口自己的方法test3

}

type A interface {
	test1()
}
type B interface {
	test2()
}
type C interface { //C接口继承了A、 B 接口
	//如果想实现接口C，那不止要实现接口C中test3的方法，还要实现接口A，B中的方法。
	A
	B
	test3()
}

type Cat struct { //实现类
	//如果想实现接口C，那不止要实现接口C中test3的方法，还要实现接口A，B中的方法。
}

func (c Cat) test1() {
	fmt.Println("运行test1。。。A接口中的方法")
}
func (c Cat) test2() {
	fmt.Println("运行test2。。。B接口中的方法")
}
func (c Cat) test3() {
	fmt.Println("运行test3。。。C接口中的方法")
}
