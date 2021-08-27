package utils1 // 该utils1叫包名，也就是该包叫 utils1包。 目录（文件夹）名也叫utils1，一般报名和文件夹名相同。
import "fmt"

// 估计是方便查找之类的
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

func Count() {
	fmt.Println("随便一个函数！")
}
