package main

import "fmt"

func main() {
	/*
		结构体嵌套：一个结构体中的字段，是另一个结构体类型。
			什么情况使用嵌套结构体
			1. 	has a
			如一个学生，“有”一个本书
			2. is a
			模拟继承性
			学生（student）是一个人（person）

	*/
	b1 := Book{}
	b1.bookName = "西游记"
	b1.price = 50.13

	s1 := Student{}
	s1.name = "小王"
	s1.age = 18
	s1.book = b1 //结构体是值传递

	fmt.Println(b1) //{西游记 50.13}
	fmt.Println(s1) //{小王 18 {西游记 50.13}}
	fmt.Printf("学生姓名：%s, 学生年龄：%d,看的书是：《%s》,书的价格是：%.2f\n", s1.name, s1.age, s1.book.bookName, s1.book.price)
	//学生姓名：小王, 学生年龄：18,看的书是：《西游记》,书的价格是：50.13

	s2 := Student{name: "小红", age: 19, book: Book{bookName: "Go语言", price: 90.1}}
	fmt.Println(s2)
	fmt.Println(s2.name, s2.age)
	fmt.Println("\t", s2.book.bookName, s2.book.price)

	s3 := Student{
		name: "Jerry",
		age:  17,
		book: Book{
			bookName: "为什么",
			price:    55.9,
		},
	}
	fmt.Println(s3.name, s3.age)
	fmt.Println("\t", s3.book.bookName, s3.book.price)

	s1.book.bookName = "红楼梦"
	fmt.Println(s1)
	fmt.Println(b1)

	b4 := Book{bookName: "财富自由", price: 88.88}
	s4 := Student2{name: "RRR", age: 33, book: &b4}
	fmt.Println(s4)
	fmt.Println(s4.name, s4.age)
	fmt.Println("\t", s4.book.bookName, s4.book.price)

	s4.book.bookName = "财富自由*2"
	fmt.Println(b4)
	fmt.Println(s4)
	fmt.Println("\t", s4.book.bookName, s4.book.price)

}

// 1.定义一个书的结构体
type Book struct {
	bookName string
	price    float64
}

type Student struct {
	name string
	age  int
	book Book
}

type Student2 struct {
	name string
	age  int
	book *Book //结构体的指针
}
