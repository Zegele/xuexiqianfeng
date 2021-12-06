package main

import "fmt"

func main() {
	/*
			方法：method
				 一个方法就是一个包含了接受者（调用者）的函数，接受者（调用者）可以是命名类型或者结构体类型的一个值或者是一个指针。
				所有给定类型的方法属于该类型的方法集。

			语法：
				func (接受者) 方法名(参数列表)(返回值列表){

				}

			总结：method，同函数类似，区别是需要有接受者。（也就是调用者）

		跟函数对比：
			A:意义：
				方法：某个类型的行为能力，需要指定的接受者调用。（只有这个类型的才可以调用）
				函数：一段独立功能的代码，可以直接调用。（任何想用这段代码的都可以调用）

			B：语法：
				方法：方法名可以相同，只要接受者不同即可。
				函数：函数名不能相同。


	*/
	w1 := Worker{name: "王二狗", age: 30, sex: "男"}
	w1.work()

	w2 := &Worker{name: "Rub", age: 10, sex: "女"}
	fmt.Printf("%T,%p\n", w2, w2)
	fmt.Println(w2)

	w2.work()

	w2.rest()
	w1.rest()

	c1 := &Cat{color: "白色", age: 2}
	c1.printInfo() //方法名可以相同，但收受者不能相同
	w2.printInfo() //方法名可以相同，但收受者不能相同
}

//1. 定义一个工人结构体
type Worker struct {
	//字段
	name string
	age  int
	sex  string
}

type Cat struct {
	color string
	age   int
}

//2.定义行为方法

func (w Worker) work() { //类型是结构体，所以是值传递
	//如果当结构体的指针作为接受者，依然可以执行。 因为地址对于的数据，就是符合需要的类型和数据。
	fmt.Println(w.name, "，在搬砖。。")
}

func (p *Worker) rest() {
	//如果是结构体作为接受者，也可以执行。其实自动转为该结构体的地址为接受者

	//fmt.Println((*p).name, ",在摸鱼。。") //结构体类型可以简写为下面的形式
	fmt.Println(p.name, ",在摸鱼。。") //直接简写，省略*号
}

func (p *Worker) printInfo() {
	fmt.Printf("工人姓名：%s，工人年龄：%d,工人性别：%s\n", p.name, p.age, p.sex)
}

func (p *Cat) printInfo() {
	fmt.Printf("猫的颜色：%s，年龄：%d\n", p.color, p.age)
}
