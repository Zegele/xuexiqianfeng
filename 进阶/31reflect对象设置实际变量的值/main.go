package main

import (
	"fmt"
	"reflect"
)

/*
	Kind和Type的区别：
		Kind有slice， map， pointer指针，struct，interface，string，Array，Function，int或者其他基本类型组成。
		例如:
		type Person struct{} //其中Kind是struct, Type 是 Person

	反射变量对应的Kind方法的返回值是基类型，并不是静态类型。
		type myint int
		var x myint = 100
		v := reflect.ValueOf(x)
		变量v的Kind依旧是reflect.Int,而不是myint这个静态类型。 Type可以表示静态类型，而Kind不可以。

*/
func main() {
	/*
		3.2 通过reflect.Value设置实际变量的值（通过反射对象修改实际变量的数值）
		reflect.Value是通过reflect.ValueOf(X)获得的，只有当X是指针的时候，才可以通过reflect.Value修改实际变量X的值。即，要修改反射类型的对象就一定要保证其值是“addressable”(地址类型)的。
		也就是说：要想修改一个变量的值，那么必须通过该变量的指针地址，取消指针的引用。
		通过refPtrVal := reflect.Valueof(&var)的方法获取指针类型，然后使用refPtrVal.elem().set(一个新的reflect.Value)来进行更改（//elem通过指针获取指针对象，然后set进行更改），传递给set()的值也必须是一个reflect.value。
			这里需要一个方法：
			func(Value)Elem
				func(v Value)Elem() Value
			解释起来就是：Elem返回接口v包含的值或指针v指向的值。如果v的类型不是interface或ptr，他会panic。如果v为零，则返回零值。
			如果你的变量是一个指针、map、slice、channel、Array。 那么你可以使用reflect.TypeOf(v).Elem()来确定包含的类型。
	*/
	var num float64 = 1.23
	fmt.Println("num的数值是：", num)

	//修改需要操作指针
	//通过reflect.ValueOf() 获取num的Value对象
	pointer := reflect.ValueOf(&num) //注意参数必须是指针才能修改其值
	newValue := pointer.Elem()       //Elem()获取pointer对应的value对象，成为newValue，然后通过对newValue的修改，达成修改num的目的。

	fmt.Println("newValue类型：", newValue.Type()) //float64
	fmt.Println("是否可以修改数据：", newValue.CanSet()) //true 可否对newValue进行修改？true 可

	//重新赋值
	newValue.SetFloat(3.14)
	fmt.Println(num)
	/*
	   说明
	   1. 需要传入的参数是*float64这个指针，然后可以通过pointer.Elem()去获取所指向的Value，注意一定要是指针。
	   2. 如果传入的参数不是指针，而是变量，那么
	   	通过Elem获取原始值对应的对象则直接panic
	   	通过CanSet方法查询是否可以设置返回false
	   3. newValue.CanSet()表示是否可以重新设置其值，如果输出的是true则可修改，否则不能修改，修改完之后再进行打印发现真的已经修改了。
	   4. reflect.Value.Elem()表示获取原始值对应的反射对象，只有原始对象才能修改，当前反射对象是不能修改的。
	   5. 也就是说如果要修改反射类型对象，其值必须是“addressable”（对应的要传入的是指针，同时要通过Elem方法获取原始值对应的反射对象）
	   6. struct或者struct的嵌套都是一样的判断处理方式

	*/
}
