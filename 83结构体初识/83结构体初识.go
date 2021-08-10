package main

import "fmt"

// 1.定义一个结构体
type Person struct {
	name    string
	age     int
	sex     string
	address string
}

func main() {
	/*
		结构体：是由一系列具有相同类型或不同类型的数据构成的数据集合。
			结构体成员是由一系列的成员变量构成，这些成员变量也被称为“字段”
	*/

	//2.初始化结构体的方法：
	//1. 方法一：
	var p1 Person
	fmt.Println(p1) //各字段的零值
	p1.name = "小李"  //p1.name 就可以访问结构体中的字段了。
	p1.age = 30
	p1.sex = "男"
	p1.address = "北京市"
	fmt.Printf("姓名: %s,年龄：%d,性别：%s,地址：%s\n", p1.name, p1.age, p1.sex, p1.address)
	fmt.Println(p1)

	//2. 方法二：
	p2 := Person{}
	p2.name = "小王"
	p2.age = 32
	p2.sex = "男"
	p2.address = "上海市"
	fmt.Println(p2)

	//3. 方法三：
	p3 := Person{name: "如花", age: 20, sex: "女", address: "成都市"}
	fmt.Println(p3)
	p4 := Person{
		name:    "小刘", //注意不能少了 , 号
		age:     18,
		sex:     "男",
		address: "东京市",
	}
	fmt.Println(p4)

	//4. 方法四：
	p5 := Person{"小张", 25, "女", "台北市"} //顺序必须和结构体定义的顺序一致。
	fmt.Println(p5)
}
