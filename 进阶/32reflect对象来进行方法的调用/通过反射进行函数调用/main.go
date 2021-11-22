package main

import (
	"fmt"
	"reflect"
	"strconv"
)

/*
说明：
	1. 要通过反射来调用起对应的方法，必须要先通过reflect.ValueOf(interface)来获取到reflect.Value，得到“反射类型对象”后才能做下一步处理。
	2. reflect.Value.MethodByName这个MethoByName，需要指定准确真实的方法名字，如果错误将直接panic，MehtodByName返回一个函数值对应的refect.Value方法名字
	3. []reflect.Value，这个是最终需要调用的方法的参数，可以没有或者一个或者多个，根据实际参数来定。
	4. reflect.Value的Call这个方法，这个方法将最终调用真实的方法，参数务必保持一致，如果reflect.Value.Kind不是一个方法，那么将直接panic
	5. 本来可以用对象访问方法直接调用的，但是如果要通过反射，那么首先要将方法注册，也就是MethodByName,然后通过反射调用methodValue.Call
*/
func main() {
	//函数的反射
	/*
		思路：函数也是看作接口变量类型
		step1：函数-->反射对象，Value
		step2：kind --> func //kind后，如果是func类型，就Call()调用
		step3：Call()
	*/
	f1 := fun1 //注意fun1表示函数1，fun1()表示调用函数1
	value := reflect.ValueOf(f1)
	fmt.Printf("kind:%s, type:%s\n", value.Kind(), value.Type()) //kind:func, type:func()

	value2 := reflect.ValueOf(fun2)
	value3 := reflect.ValueOf(fun3)
	fmt.Printf("kind:%s, type:%s\n", value2.Kind(), value2.Type()) //kind:func, type:func(int, string)
	fmt.Printf("kind:%s, type:%s\n", value3.Kind(), value3.Type()) //kind:func, type:func(int, string) string

	//通过反射调用函数
	value.Call(nil)
	value2.Call([]reflect.Value{reflect.ValueOf(1), reflect.ValueOf("hello")})
	resultValue := value3.Call([]reflect.Value{reflect.ValueOf(12), reflect.ValueOf("helloWorld!")})
	fmt.Printf("%T\n", resultValue) //[]reflect.Value //Call()的返回值是个切片类型 是reflect.Value类型的切片
	fmt.Println(resultValue)
	fmt.Println(len(resultValue))                                                  //1
	fmt.Printf("kind:%s, type:%s\n", resultValue[0].Kind(), resultValue[0].Type()) //kind:string, type:string

	//用interface()把反射的对象转换为实际的数据类型
	s := resultValue[0].Interface().(string)
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

func fun1() {
	fmt.Println("我是函数fun1(),无参的。。。")
}

func fun2(i int, s string) {
	fmt.Println("我是函数fun2(),有参的。。。", i, s)
}
func fun3(i int, s string) string {
	fmt.Println("我是函数fun3(),有参的，有返回值。。。", i, s)
	return s + strconv.Itoa(i)
}
