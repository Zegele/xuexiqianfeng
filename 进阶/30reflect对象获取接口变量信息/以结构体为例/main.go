package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
	Age  int
	Sex  string
}

func (p Person) Say(msg string) { //Person的Say方法
	fmt.Println("hello", msg)
}
func (p Person) PrintInfo() { //Person的PrintInfo方法
	fmt.Printf("姓名： %s, 年龄： %d， 性别： %s\n", p.Name, p.Age, p.Sex)
}
func main() {
	p1 := Person{"小李", 20, "女"}
	GetMessage(p1)
}

//获取input的信息
func GetMessage(input interface{}) {
	getType := reflect.TypeOf(input) //先获取input的类型
	fmt.Println("get Type is :", getType.Name())
	fmt.Println("get Kind is :", getType.Kind())

	getValue := reflect.ValueOf(input)
	fmt.Println("get all Fields is :", getValue) //{小李 20 女}

	//获取字段
	/*
		step1: 先获取Type对象：reflect.Type，
			NumField()//可以知道有几个字段
			Field(index)//通过下标可获得字段
		step2: 通过Filed()获取每一个Filed字段
		step3： Interface()，得到对应的Value
	*/

	for i := 0; i < getType.NumField(); i++ {
		filed := getType.Field(i)
		value := getValue.Field(i).Interface() //获取一个数值
		fmt.Printf("字段名称： %s，字段类型；%s, 字段数值： %v\n", filed.Name, filed.Type, value)
	}

	//i := getType.NumField()
	d := getType.Field(1)
	fmt.Println(d)
	fmt.Println(d.Name)

	//获取方法

	for i := 0; i < getType.NumMethod(); i++ {
		method := getType.Method(i)
		fmt.Printf("方法名称：%s, 方法类型：%v\n", method.Name, method.Type)
	}

}

/*
通过运行结果可以得知获取位置类型的interface的具体变量及其类型的步骤为：
	1. 先获取interface的reflect.Type,然后通过NumFiled进行遍历
	2. 再通过reflect.Type的Field(index)获取其Field的类型全部内容， 继续 Field.Name可获得类型名，Field.Type可获得类的类型等
	3. 最后通过reflect.Value.Field(index)的Interface() 得到对应的value

通过运行结果可以得知获取未知类型的interface的所属方法（函数）的步骤为：
	1. 先获取interface的reflect.Type,然后通过NumMethod进行遍历
	2. 再分别通过reflect.Type的Method获取对应的真实的方法（函数）
	3. 最后对结果取其Name，和Type得知具体的方法名。
	4. 也就是说反射可以将“反射类型对象”再重新转换为“接口类型变量”
	5. struct或者struct的嵌套都是一样的判断处理方法。
*/
