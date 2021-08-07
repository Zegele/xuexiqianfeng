package main

import "fmt"

func main() {

	/*
		递归函数：一个函数自己调用自己，就叫做递归函数。
			递归函数要有一个出口，逐渐的向出口靠近。

	*/
	// 1. 求1-5的和
	sum := getSum(5)
	fmt.Println(sum)

	// 2.fibonacci数列
	n := 12
	fi1 := fibo(n)
	fmt.Printf("第%d位的fibo数列是：%d。\n", n, fi1)

	fi2 := getFibonacci(n)
	fmt.Printf("第%d位的fibo数列是：%d**\n", n, fi2)
}

func getSum(n int) int {
	fmt.Println("&&&&&&&&&")
	if n == 1 {
		return 1
	}
	return getSum(n-1) + n
}

/*
0   第一位
1	第二位
1	第三位
2	第四位
3	第五位
5	第六位
8	第七位
13	第八位
21	第九位
34	第十位
*/

func fibo(n int) int {
	if n <= 0 {
		fmt.Println("输入有误")
		return -1
	}

	if n == 1 {
		return 0
	}
	if n == 2 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)

}

//老师的做法 第一项，第二项是1  我写的第一项是0，第二项是1
/*
1	2	3	4	5	6	7	8	9	10	11	12 ...
1	1	2	3	5	8	13	21	34	55	89	144	...
*/

func getFibonacci(n int) int {
	if n == 1 || n == 2 {
		return 1
	}
	return getFibonacci(n-1) + getFibonacci(n-2)
}
