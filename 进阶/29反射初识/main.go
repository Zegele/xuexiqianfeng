package main

/*
	反射reflect
	在计算机科学中，反射是指计算机程序在运行时（run time）可以访问、检测和修改它本身状态或行为的一种能力。
	用比喻来说，反射就是程序在运行的时候能够“观察”并且修改自己的行为。
	go创始人之一对反射定义：反射是一种让程序————主要是通过类型————理解其自身结构的一种能力。他是元编程的组成之一，同时它也是一大引入困惑的难题。
	《GO 语言圣经》中这样定义反射：Go语言提供了一种机制在运行时更新变量和检查他们的值、调用他们的方法，但是在编译时并不知道这些变量的具体类型，这称为反射机制。
	为什么要反射
	需要反射的2个场景
	1. 有时你需要编写一个函数，但是并不知道传给你的参数类型是什么，可能没有约定好；也可能是传入的类型很多，这些类型并不能同一表示。这是反射就会用得上了。
	2. 有时候需要根据某些条件决定该调用哪个函数，比如根据用户的输入来决定。这时就需要对函数和函数的参数进行反射，在运行期间动态地执行函数。

	但是对于反射，还是有几点不太建议使用反射的理由：
	1. 与反射相关的代码，经常是难以阅读的。在软件工程中，代码可读性也是一个非常重要的指标。
	2. Go语言作为一门静态语言，编码过程中，编译器能提前发现一些类型错误，但是对于反射代码是无能为力的。所以包含反射相关的代码，很可能会运行很久，才会出错，这时候经常是直接panic，可能会造成严重的后果。
	3. 反射对性能影响还是比较大的，比正常代码运行速度慢一到两个数量级。
	所以，对于一个项目中处于运行效率关键位置的代码，尽量避免使用反射特性。

	二、相关基础
	反射是如何实现的？我们以前学习过interface， 它是Go语言实现抽象的一个非常强大的工具。
	当向接口变量赋予一个实体类型的时候，接口会存储实体的类型信息，反射就是通过接口的类型信息实现的，反射建立在类型的基础上。
	Go语言在reflect包里定义了各种类型，实现了反射的各种函数，通过他们可以在运行时检测类型的信息，改变类型的值。
	重新温习下Go语言相关的特性
	go语言是静态类型语言：编译时类型已经确定，比如对以基本数据类型的再定义后的类型，反射时候需要确认返回的时何种类型。
	空接口interface{}： go的反射机制是要通过接口来进行的，而类似于Java的Object的空接口可以和任何类型进行交互，因此对基本数据类型等的反射也直接利用了而这一特点。

	go语言的类型：
	1. 变量包括（type、 value）两部分
		理解这一点就知道为什么nil != nil了
	2. type包括static type(静态类型)和concrete type（具体类型，系统或程序运行中获得的类型），简单来说静态类型是在编码中看见的类型（如int，string）
		concrete type是run time时系统看见的类型。
	3. 类型断言能否成功过，取决于变量的concrete type，而不是static type。
		因此，一个reader变量如果它的concrete type也实现了write方法的话，它也可以被类型断言为writer。

	反射主要于interface类型相关（它的type时concrete type），只有interface类型才有反射一说。

	每个interface变量都有一个对应pair，pair中记录了实际变量的值和类型：（value， type）
	value是实际变量值，type是实际变量的类型。一个interface{}类型的变量包含了2个指针，一个指针指向值的类型（对应concrete type），另一个指针指向实际的值（对应value）
	反射就是用来检测存储在接口变量内部（值value；类型concrete type）pair对的一种机制。

	三、反射的使用。
	3.1 reflect的基本功能TypeOf和ValueOf
	func ValueOf(i interface{}) Value{...}
	ValueOf用来获取输入参数接口中的数据的值，如果接口为空则返回0

	func TypeOf(i interface{}) Type{...}
	TypeOf用来动态获取输入参数接口中的值的类型，如果接口为空则返回nil

	reflect.TypeOf()是获取pair中的type，reflect.ValueOf()获取pair中的value

	使用reflect一般分成三步：
	1. 首先得有个接口类型的变量
	2. 把它转成reflect对象（reflect.Type或reflect.Value对象）
	3. 根据不同的情况调用不同的函数。
*/
import (
	"fmt"
	"reflect"
)

func main() {

	//反射操作：通过反射，可以获取一个接口类型变量的 类型和数值
	var x float64 = 3.4
	fmt.Println("type: ", reflect.TypeOf(x))   //type: float64
	fmt.Println("value: ", reflect.ValueOf(x)) // value: 3.4

	fmt.Println("--------------")
	//根据反射的值，来获取对应的类型和数值
	v := reflect.ValueOf(x)
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64)
	fmt.Println("type: ", v.Type())
	fmt.Println("value:", v.Float())
}
