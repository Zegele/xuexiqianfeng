package main //主函数所在的包 程序入口

import (
	"fmt"

	"xuexiqianfeng/进阶/1包的使用/pk1"
	"xuexiqianfeng/进阶/1包的使用/utils1" //绝对路径

	"xuexiqianfeng/进阶/1包的使用/utils1/timeutils"

	p "xuexiqianfeng/进阶/1包的使用/person" //取别名
)

func main() {
	/*
		关于包的使用：
		1. 一个目录下的统计文件归属一个包，package的声明要一致。
		2. package声明的包和对应的目录名可以不一致，但习惯上还是写成一致的。
		3. 包可以嵌套
		4. 同包下的函数不需要导入包，可以直接使用
		5. main包，main()函数所在的包，其他的包不能使用。
		6. 导入包的时候，路径要从src下开始写。
		7. 起别名，如下；
			import(
				p1 "package1"
				p1 "package1"
			)
			p1.Method()
		8. _操作
			涉及到init函数。
	*/
	utils1.Count() // 包名.函数名 //utils1包的Count()函数
	//注意这里时包名，不是文件名，也不是文件夹名（目录名）
	//包名是可以与文件名，目录名不同的。但一般设置是一致的，方便查找。比如不用打开文件专门查看包名，通过看目录名就知道是什么包。提高使用效率。
	timeutils.PrintTime()

	pk1.Mytest2()

	utils1.Mytest1() //相同的包下，不需要导包。如，Mytest1() 本身就在utils1包下，且该包下有Count()函数，则可以直接使用该函数，不需要导入包。
	//详情查看utils1包下的test1.go

	fmt.Println("?????????????")
	p1 := p.Person{Name: "二狗", Age: 10, Sex: "男"}
	fmt.Println(p1)
}

/*
包需要满足：
1. 一个目录下的同级文件归属一个包，也就是说，在同一个包下面的所有文件的package名（包名），都是一样的。
2. 在同一个包下面的文件名，package名（包名）都建议设为该目录名，但也可以不是。也就是说，包名可以与其目录不同名。
3. 包名为main的包为应用程序的入口包，其他包不能使用。

导入包：
导入一个包：
	import "package1"

导入多个包：
	import(
		"package1"
		"package2"
		)

*/
