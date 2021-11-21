package main

import (
	"fmt"
	"reflect"
)

/*
	反射的概念中，编译时就知道变量类型的是静态类型；运行时才知道一个变量类型的叫动态类型。
	静态类型：
	type Myint int // int 就是静态类型

	var i *int // *int就是静态类型

	type A struct{
		Name string // string就是静态类型
	}

	动态类型：
	运行时这个变量赋值时，这个值的类型（如果值为nil的时候没有动态类型）。一个变量的动态类型在运行时可能改变，这主要依赖于他的赋值（前提是这个变量是接口类型）。
	var A interface{} //静态类型interface{}
	A = 10 //静态类型为interface，动态为int
	A = "string" //静态类型为interface{}， 动态为string
	var M *int
	A = M // A的值可以改变

	四、反射的规则
	反射的操作步骤非常的简单，就是通过实体对象获取反射对象（Value， Type），然后操作相应的方法即可。

	反射API的分类总结如下：
	1. 从实例到Value
	通过实例获取Value对象，直接使用reflect.ValueOf()。例如：
	func ValueOf(i interface{}) Value

	2. 从实例到Type
	通过实例获取反射对象的Type，直接使用reflect.TypeOf()函数。例如：
	func TypeOf(i interface{}) Type

	3. 从Type到Value
	Type里面只有类型信息，所以直接从一个Type接口变量里面是无法获得实例的Value的，但可以通过该Type构建一个新实例的Value。reflect包提供了两种方法，示例如下：
	//New 返回的是一个Value，该Value的type为PtrTo(typ),即Value的Type是指定同一批的指针类型
	func New(typ Type) Value
	//Zero返回的是一个typ类型的零值，注意返回的Value不能寻址，位不可改变
	func Zero(typ Type) Value

	如果知道一个类型的底层存放地址，则还有一个函数是可以依据type和该地址值恢复出Value的。例如：
	func NewAt(typ Type, p unsafe.Pointer) Value

	4. 从Value到Type
	从反射对象Value到Type可以直接调用Value的方法，因为Value内部存放着Type类型的指针。例如：
	func(v Value)Type() Type

	5. 从Value到实例
	Value本身就包含类型和值信息，reflect提供了丰富的方法来实现从Value到实例的转换。例如：
	//该方法最通用，用来将Value转化为空接口，该空接口内部存放距离类型实例
	//可以使用接口类型查询去还原为具体的类型
	func (v Value) Interface() (i interface{})

	//Value 自身也提供丰富的方法，直接将Value转换为简单类型实例，如果类型不匹配，则直接引起panic
	func (v Value) Bool() bool
	func (v Value) Float() float64
	func (v Value) int() int64
	func (v Value) Uint() uint64

	6. 从Value的指针到值
	从一个指针类型的Value获得值类型Value有两种方法，例如：
	//如果 v 类型是接口，则Elem()返回接口绑定的实例的 Value，如采 v 类型是指针，则返回指针值的 Value，否则引起panic
	func (v Value) Elem() Value
	//如果v 是指针，则返回指针值的value，否则返回v自身，该函数不会引起panic
	func Indirect(v Value) Value

	7. Type 指针和值的相互转换
	指针类型Type到值类型Type。例如：
	//t 必须是Array, Chan, Map, Ptr, Slice, 否则会引起panic
	//Elem 返回的是其内部元素的 Type
	t.Elem() Type

	值类型Type到指针类型Type。例如：
	//PtrTo 返回的是指向 t 的指针型 Type
	func PtroTo(t Type) Type

	8. Value值的可修改性
	value值的修改涉及如下两个方法：
	//通过 CanSet 判断是否能修改
	func (v Value) CanSet() bool
	//通过Set进行修改
	func (v Value) Set(x Value)

	Value值在什么情况下可以修改？我们知道实例对象传递给接口的是一个完全的值拷贝，如果调用反射的方法reflect.ValueOf()传进去的是一个值类型变量，则获得的Value实际上是原对象的一个副本，这个Value是无论如何也不能被修改的。
	也就是如果是Array类型的，是修改不了原Value值的。

	Go关于反射的三大定律：
	第一条：是最基本的，反射可以从接口值得到反射对象。
		反射是一种检测存储在interface{}中的类型和值机制。这可以通过TypeOf函数和ValueOf函数得到。
	第二条：和上一条相反的机制，反射可以从反射对象获得接口值
		他将ValueOf的返回值通过interface()函数反向转变成interface变量。
	第三条：如果需要操作一个反射变量，则其值必须可以修改。
		反射变量可设置的本质是存储了原变量本身，这样对反射变量的操作，就会反映到原变量本身；
		反之，如果反射变量不能代表原变量，那么操作了反射变量，不会对原变量产生任何影响，这会给使用者带来疑惑。
		所以第二种情况在语言层面是不被允许的。

*/
func main() {
	var num float64 = 1.23 //由于空接口可以是任意类型，所以在演示的时候，把它认为是个空接口的实体。
	//“接口类型变量”--> "反射类型对象"
	value := reflect.ValueOf(num) //获得了接口的Value值
	fmt.Println(value)

	//"反射类型对象" --> “接口类型变量”
	convertValue := value.Interface().(float64) //通过已知Value的类型，获取Value值的真实数值。
	fmt.Println(convertValue)
	/*
	   反射类型对象 --> 接口类型变量，理解为“强制转换”
	   Golang对类型要求非常严格，类型一点更要完全符合
	   一个是*float64，一个float64,如果混淆，直接panic
	*/
	pointer := reflect.ValueOf(&num) //获得了接口的Value值的地址
	fmt.Println(pointer)
	convertPointer := pointer.Interface().(*float64)
	fmt.Println(convertPointer)
}
