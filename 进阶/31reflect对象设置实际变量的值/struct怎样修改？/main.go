package main

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name   string
	Age    int
	School string
}

func main() {
	s1 := Student{"小明", 18, "自学"}
	//通过反射，更改对象的数值：前提是数据可以被更改
	fmt.Println("%T,%v\n", s1, s1)
	p1 := &s1
	fmt.Printf("%T,%v\n", p1, p1)
	fmt.Println(s1.Name)
	fmt.Println((*p1).Name, p1.Name) //(*p1).Name,p1.Name这两种写法都可以

	//通过反射，改变struct数值
	value := reflect.ValueOf(&s1) //或reflect.ValueOf(p1)
	fmt.Println(value)
	if value.Kind() == reflect.Ptr { //如果value.Kind()（种类） 等于 指针类型
		newValue := value.Elem()
		fmt.Println(newValue)
		fmt.Println(newValue.CanSet()) //true

		f1 := newValue.FieldByName("Name")
		f1.SetString("小红")
		f3 := newValue.FieldByName("School")
		f3.SetString("家里蹲")
		fmt.Println(s1)
	}
}
