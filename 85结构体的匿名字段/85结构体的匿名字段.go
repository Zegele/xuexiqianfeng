package main

import "fmt"

func main() {
	/*
		匿名结构体和匿名字段：

		匿名结构体：没有名字的结构体。
			在创建弥明结构体时，同时创建对象。

			变量名 := struct{
				定义字段Field
			}{
				对字段进行复制
			}

		匿名字段：一个结构体的字段没有名字。

		匿名函数：

	*/
	s1 := Student{name: "小张", age: 18}
	fmt.Println(s1.name, s1.age)
	fmt.Println(s1)

	func() {
		fmt.Println("hello world...")
	}()

	s2 := struct { //匿名结构体 这种只能使用一次，一般这样写在代码中意义不大
		name string
		age  int
	}{
		name: "小王",
		age:  19,
	}

	fmt.Println(s2.name, s2.age)
	fmt.Println(s2)

	w2 := Worker{"小李", 30}
	fmt.Println(w2)
	fmt.Println(w2.string, w2.int) //匿名字段，默认使用数据类型作为名字。

}

type Student struct {
	name string
	age  int
}

type Worker struct {
	string //匿名字段
	//string //字段类型不能重复，否则会冲突
	int //匿名字段
}
