package main

import "fmt"

func main() {
	/*
		面向对象：OOP
		Go语言的结构体嵌套：
			1.模拟继承性： is a   如: 学生(studient)是个人(person)
				type A struct{
					field
				}
				type B struct{
					A //匿名字段
				}
		提升字段：当一个结构体中有一个匿名结构体，那么该匿名结构体中的字段，就叫做提升字段。


			2. 模拟聚合关系: has a  //86节课
				type C struct{
					field
				}
				type D struct{
					1包的使用 C //聚合关系
				}

	*/
	//模拟继承
	//1.创建父类的对象
	p1 := Person{name: "小张", age: 19}
	fmt.Println(p1)

	//2.创建子类的对象
	s1 := Student{Person: Person{name: "小王", age: 30}, school: "jiali"}
	//简写： s2 := Student{Person{"小王", 30}, "jiali"} 这样写，貌似看起来不方便。
	fmt.Println(s1)

	s2 := Student{}
	s2.Person.name = "小李" //s2.name = "小李"  因为是提升字段，所以可以这样写。
	s2.Person.age = 19    //s2.age = 19  因为是提升字段，所以可以这样写。
	s2.school = "hao"
	fmt.Println(s2)

	var s3 Student
	s3.Person.name = "小混"
	s3.Person.age = 5
	s3.school = "幼儿园"
	fmt.Println(s3)

	s3.name = "ru" //提升字段的用法  s3.Person.name = "ru" 由于Person在Strudent结构体中是匿名结构体，所以Person中的字段叫提升字段，可以直接写。
	s3.age = 10
	fmt.Println(s3)

	/*
		s3.Person.name ---> s3.name
		Student结构体将Person结构体作为一个匿名结构体字段了
		那么Person结构体中的字段，对于Student来讲，就是提升字段。
		所以Student对象可以直接访问Person中的字段。
	*/

}

//模拟继承
//1.定义父类
type Person struct {
	name string
	age  int
}

//2.定义子类
type Student struct {
	Person
	school string
}
