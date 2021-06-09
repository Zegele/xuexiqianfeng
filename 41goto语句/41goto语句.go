package main

import "fmt"

func main() {
	/*
		goto：可以无条件地转移到过程中指定的行。一般配合条件语句使用。
		goto label;
		..
		..
		label: statement;

	*/

	var a = 10
LOOP:
	for a < 20 {
		if a == 15 {
			a += 1
			goto LOOP
		}
		fmt.Printf("a的值为：%d\n", a)
		a++
	}

	fmt.Println("...............")
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println(i, j)
			if j == 2 {
				goto breakHere
			}
		}
	}
	return //手动返回，避免执行进入标签

breakHere:
	fmt.Println("done....")

}
