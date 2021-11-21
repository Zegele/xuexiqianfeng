package main

/*
	3.3 通过reflect.Value来进行方法的调用
	这算是一个reflect的高级用法了，前面我们说到对类型、变量的集中反射的用法，包括如何获取其值，其类型，以及如何重新设置新值。
	但是再项目应用中，另外一个常用并且属于高级的用法，就是通过reflect来进行方法（函数）的调用。
	比如我们要做框架工程的时候，需要可以随意扩展方法，或者说用户可以自定义方法，那么我们通过什么手动来扩展让用户能够自定义呢？
	关键点在于用户的自定义方法是未可知的，因此我们可以通过reflect来搞定。

	Call() 方法
	func (Value) Call
		func (v Value) Call(in []Value) []Value// 参数 in []Value 是切片类型，是设置函数的参数。如果函数没有参数，就使用nil，或空切片。

	通过反射，调用方法。
	先获取结构体对象，然后
*/
import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) {
	fmt.Println("hello,", msg)
}
func (p Person) PrintInfo() {
	fmt.Printf("姓名：%s, 年龄：%d，性别： %s\n", p.Name, p.Age, p.Sex)
}

func main() {
	/*
		通过反射来进行方法的调用
		思路：
		step1:接口变量->反射对象：Value  //通过反射获取对应的反射对象  通过reflect.ValueOf()实现
		step2：获取对应的方法对象 MethodByName()
		step3:将方法对象进行调用：Call()
	*/
	p1 := Person{"小花", 20, "女"}
	value := reflect.ValueOf(p1)
	fmt.Printf("kind: %s, type:%s\n", value.Kind(), value.Type()) //kind: struct, type:main.Person

	methodValue1 := value.MethodByName("PrintInfo")                           //获取PrintInfo方法 ，返回值是个Value
	fmt.Printf("kind:%s,type:%s\n", methodValue1.Kind(), methodValue1.Type()) //kind:func,type:func() //无参数

	//1. 对没有参数的方法或函数怎样调用？
	methodValue1.Call(nil) //没有参数的函数，这里使用nil，则可以调用

	args1 := make([]reflect.Value, 0) //空切片
	methodValue1.Call(args1)          //传入空切片，也能调用该函数。

	//2. 对有参数的方法或函数怎样调用？
	methodValue2 := value.MethodByName("Say")
	fmt.Printf("kind:%s,type:%s\n", methodValue2.Kind(), methodValue2.Type()) //kind:func,type:func(string) //有参数
	args2 := make([]reflect.Value, 1)                                         //先创建一个反射对象Value类型的切片
	args2[0] = reflect.ValueOf("yyy")                                         //要传入字符串类型，但是要先转成反射对象
	//args2 := []reflect.Value{reflect.ValueOf("hhhh")} //要传入字符串类型，但是要先转成反射对象

	methodValue2.Call(args2)
}
