package main

import (
	"errors"
	"fmt"
)

func main() {
	/*
		error: 内置的数据类型，内置的接口
			定义方法：Error() string
			会打印出“标准”错误信息 如open test.txt: The system cannot find the file specified.

		使用go语言提供好的包：
			errors包下的函数：New()，创建一个error对象 就是对错误信息的描述，如：该数据为负，非法。
			fmt包下的Error()函数：
				func Errorf(format string, a ...interface{}) error
	*/
	//1. 创建一个error返回数据
	err1 := errors.New("创建一个错误反馈信息！！！")
	fmt.Println(err1)
	fmt.Printf("%T\n", err1) //*errors.errorString

	//2. 另一个error返回数据
	err2 := fmt.Errorf("错误的信息码：%d", 100)
	fmt.Println(err2)
	fmt.Println(err2.Error()) //调用Error()方法，上面是简写。
	fmt.Printf("%T\n", err2)  //*errors.errorString

	fmt.Println("________________")
	err3 := checkAge(-30)
	if err3 != nil {
		fmt.Println(err3)
		return
	}
	fmt.Println("程序正常")

}

//设计一个函数：验证年龄是否合法，如果为负数，就返回一个error
func checkAge(age int) error {
	if age < 0 {
		//返回error对象
		//return errors.New("年龄不合法") //第一种方法
		err := fmt.Errorf("您给定的年龄是：%d,不合法", age) //Errorf()该函数就是针对error的
		return err                               //第二种方法
	}
	fmt.Println("年龄是：", age)
	return nil
}
