package main

import "fmt"

func main() {
	/*
		Go语言的数据类型： 1.基本数据类型； 2. 复合数据类型（衍生数据类型）
		1. 基本数据类型：
			1.1布尔类型: bool
				取值: true, false
			1.2数值类型:
				整数：
				浮点数：小数
				复数： complex
			1.3字符串： string

		2. 符合数据类型
		array(数组)， slice（切片）， map（地图）， function（函数）， pointer（指针）， struct（结构体）， interface（接口）， channel（通道）...


	*/

	// 1.1 布尔类型
	var b1 bool
	b1 = true
	fmt.Printf("%T, %t\n", b1, b1)
	b2 := false //简短定义
	fmt.Printf("%T, %t\n", b2, b2)

	// 1.2 数值型
	/* 1.2.1整数型
	整数分：	有符号：最高位表示符号位，0正数，1负数，其余位表示数值
			无符号：所有的位都表示数值
	int8
	有符号8位整型（-128到127）
	长度：8bit

	int16
	有符号16位整型（-32768到32767）

	int32
	有符号32位整型（-2147483648到2147483647）

	int64
	有符号64位整型（-9223372036854775808到9223372036854775807）

	uint8
	无符号8位整型（0到255）

	Unit16
	无符号16位整型（0到65535）

	uint32
	无符号32位整型（0到4294967295）

	uint64
	无符号64位整型（0到18446744073709551615）

	int和uint: 根据底层平台，表示32或64位整数。
	除非需要使用特定大小的整数，否则通常应该使用int来表示整数。
	大小：32位系统32位，64位系统64位。

	注意：
	byte:uint8
	rune:int32
	*/
	//1.2.1 整数
	var i1 int8 //-128~127
	i1 = 100
	fmt.Println(i1)
	//i1 = 200 //溢出

	var i2 uint8
	i2 = 200
	fmt.Println(i2)

	var i3 int //我的系统是64位的，所以这个int，系统会默认为int64
	i3 = 1000
	fmt.Println(i3)
	// 语法角度： 系统默认int位int64，但与直接设定的int64，不是同一种类型。

	//var i4 int64
	//i4 = i3 //cannot use i3 (type int) as type int64 in assignment

	var i5 uint8
	i5 = 100
	var i6 byte
	i6 = i5 //uint8类与byte类一致，可通用
	fmt.Println(i5, i6)

	var i7 = 888 //隐形定义，也就是类型推断的时候是int型
	fmt.Printf("%T,%d\n", i7, i7)

	/* 1.2.2 浮点型
	float32
	IEEE-754 32位浮点型数（单精度）

	float64
	IEEE-754 64位浮点型数（双精度）

	complex64
	32位实数和虚数

	complex128
	64位实数和虚数
	*/

	//浮点
	var f1 float32
	f1 = 3.14
	var f2 float64
	f2 = 4.67
	fmt.Printf("%T,%f\n", f1, f1)   //默认6位
	fmt.Printf("%T,%.2f\n", f2, f2) //%.2f 就是2位
	fmt.Printf("%T,%.3f\n", f1, f1) //%.3f 就是3位
	fmt.Println(f1, f2)             //本身是几位小数就是几位

	f3 := 2.55                  //类型推断是float64类型的。
	fmt.Printf("%T,%f", f3, f3) // float64,2.550000

}
