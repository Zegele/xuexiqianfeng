package main

import (
	"fmt"
	"math"
)

func main() {
	/*
		接口断言：貌似就是判断接口中的实际参数，到底是什么，是什么类型的数据。如：你看到的就是t1，这个t1实际是什么类型的数据呢？就需要接口断言
			方式一：
				1. instance := 接口对象.(实际类型) //不安全，会panic（）
				2. instance , ok := 接口对象.(实际类型)//安全

			方式二：
				switch instance := 接口对象.(type){
				case 实际类型1：
					。。。
				case 实际类型2：
					。。。
				。。。
				}
	*/
	t1 := Triangle{a: 4, b: 5, c: 6}
	fmt.Println(t1.peri())
	fmt.Println(t1.area())
	fmt.Println(t1.a, t1.b, t1.c)

	var c1 Circle = Circle{radius: 4}
	fmt.Println(c1.radius)
	fmt.Println(c1.peri())
	fmt.Println(c1.area())

	var shape1 Shape
	shape1 = t1
	fmt.Println(shape1.peri())
	fmt.Println(shape1.area())
	//fmt.Println(shape1.a, shape1.b, shape1.c) 接口中是不能访问实现类中的属性字段的。因为接口中不包含该数据。

	var shape2 Shape
	shape2 = c1
	fmt.Println(shape2.peri())
	fmt.Println(shape2.area())
	//fmt.Println(shape2.radius)接口中是不能访问实现类中的属性字段的。因为接口中不包含该数据。

	testShape(t1)
	testShape(c1)
	testShape(shape1)
	testShape(shape2)

	//接口断言 本节课重点
	getType(t1)
	getType(c1)
	getType(shape1)
	getType(shape2)

	var t2 *Triangle = &Triangle{3, 4, 2}
	fmt.Println("t2:%Tm%p\n", t2, &t2)
	getType(t2)
	getType2(t2)
	getType(t1)
	getType2(t1)

}

//1.定义一个接口
type Shape interface {
	peri() float64 //求形状的周长  peri()函数，返回值是float64
	area() float64 // 求形状的面积
}

//2. 定义实现类：三角形
type Triangle struct {
	a, b, c float64
}

func (t Triangle) peri() float64 {
	return t.a + t.b + t.c
}
func (t Triangle) area() float64 {
	p := t.peri() / 2
	s := math.Sqrt(p * (p - t.a) * (p - t.b) * (p - t.c)) //海伦公式
	return s
}

type Circle struct {
	radius float64 //半径
}

func (c Circle) peri() float64 {
	return c.radius * 2 * math.Pi
}

func (c Circle) area() float64 {
	return math.Pow(c.radius, 2) * math.Pi
	//math.Pow(x,y) 该函数表示： x的y次方，或x的y次幂
}

func testShape(s Shape) {
	fmt.Printf("周长：%.2f,面积：%.2f\n", s.peri(), s.area())
}

//这节课的重点
func getType(s Shape) {
	//断言
	if ins, ok := s.(Triangle); ok { //1.初识化： ins, ok := s.(Triangle) 2. 判断：ok不ok？ 类似:初始化i:=0, 如果i<10,就怎么怎么的
		//这种写法的意思是：初始化ins, ok := s.(Triangle)，如果ok（表示数据存在），则执行内部，如果不ok（不存在），执行之外的（也就是else）
		fmt.Println("是三角形，三边是：", ins.a, ins.b, ins.c)
	} else if ins, ok := s.(Circle); ok {
		fmt.Println("是圆形，半径是：", ins.radius)
	} else if ins, ok := s.(*Triangle); ok {
		fmt.Printf("ins:%T,%p\n", ins, &ins)
		fmt.Printf("s:%T,%p\n", s, &s)
		fmt.Println("是！是！是三角形！三边是：", (*ins).a, ins.b, ins.c) // (*ins).a, ins.b 这两种写法都行。 结构体可以默认简写这种写法
	} else {
		fmt.Println("我也不知道是什么形状。")
	}
}

func getType2(s Shape) {
	inss := s
	fmt.Println(inss)
	switch ins := s.(type) { //注意这个type
	case Triangle:
		fmt.Println(ins)
		fmt.Println("三角形。。", ins.a, ins.b, ins.c)

	case Circle:
		fmt.Println("圆形：", ins.radius)
	case *Triangle:
		fmt.Println("三角形结构体指针：", ins.a, ins.b, ins.c)
	}
}
