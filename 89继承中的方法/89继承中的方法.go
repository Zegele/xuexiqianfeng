package main

import "fmt"

func main() {
	/*
		OOP（面向对象）中的继承性：
			如果两个类（class）存在就继承关系，其中一个是子类，另一个作为父类，那么：

			1. 子类可以直接访问父类的属性和方法。
			2. 子类可以新增自己的属性和方法。
			3. 子类可以重写父类的方法（orverride,就是将父类已有的方法，重新实现）


		Go语言的结构体嵌套：
			1. 模拟继承性： is a
				type A struct{
					field
				}

			2. 模拟聚类关系： has a
				type C struct{
					field
				}
				type D struct{
					c C //聚合关系
				}

	*/

	//1.创建Person类型
	p1 := Person{name: "老谭", age: 36}
	fmt.Println(p1.name, p1.age) //父类对象，访问父类的字段属性
	p1.eat()                     //父类对象，访问父类的方法。

	//2. 创建子类对象
	s1 := Student{Person: p1, school: "Goxuexi"}
	fmt.Println(s1.name) //s1.Person.name 结构体作为匿名字段，其内部的字段，叫提升字段。
	// 提升字段 详见87课

	s1.eat()   //子类对象，访问父类方法。 子类重写父类方法后，子类将不再访问父类方法。
	s1.study() //子类对象,访问自己新增的方法
	s1.eat()   //子类重写父类的方法后，调用的是重写的方法。
}

//1.定义一个“父类”
type Person struct {
	name string
	age  int
}

//2. 定义一个“子类”
type Student struct {
	Person //结构体嵌套，模拟继承性
	school string
}

//3. 方法
func (p Person) eat() { //父类的方法
	fmt.Println("父类的方法，吃窝窝头。。")
}

func (s Student) study() { //子类的方法
	fmt.Println("子类新增的方法，学习了")
}

func (s Student) eat() { //父类的方法
	fmt.Println("子类重写父类的方法，吃米饭。。")
}
