package pk1

import (
	"fmt"
	"xuexiqianfeng/进阶/2init函数/utils"
)

func init() {
	fmt.Println("pk1包下的init()函数，我是怎么排序的？")
}

/*
多个包中如果都有init()函数，如果他们不是相互依赖的，则是按照main函数的调用顺序初始化的。先import的先初始化。
*/
func MyTest1() {
	fmt.Println("pk1包下的，MyTest1()函数。。。")
	utils.Count()
}
