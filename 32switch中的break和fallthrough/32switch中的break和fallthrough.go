package main

import "fmt"

func main() {
	/*
		switch中的break和fallthrough语句
		break：可以使用再switch中，也可以使用在for循环中。
			强制结束case语句，从而结束switch分支。

		fallthrough：用于穿透switch
			当switch中某个case匹配成功之后，就执行该case
			如果遇到fallthrough，那么后面紧邻的case，无需匹配，执行穿透窒执行。

			fallthrough应该位于某个case的最后一行。
	*/
	n := 2
	switch n {
	case 1:
		fmt.Println("你是熊大")
		fmt.Println("你是熊大")
		fmt.Println("你是熊大")
	case 2:
		fmt.Println("你是熊二")
		fmt.Println("你是熊二")
		break // 用于强制结束case，意味着switch也被强制结束。
		fmt.Println("你是熊二")
	case 3:
		fmt.Println("你是光头强")
	}

	fmt.Println("..............")
	m := 2
	switch m {
	case 1:
		fmt.Println("第一季度")
	case 2:
		fmt.Println("第二季度")
		fallthrough
	case 3:
		fmt.Println("第三季度")
		fallthrough //必须放在该case的最后一行
	case 4:
		fmt.Println("第四季度")
	}
	fmt.Println("main over.")
}
