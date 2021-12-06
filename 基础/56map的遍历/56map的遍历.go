package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		map 的遍历：
			1. 使用：for range
				每次打印的数据顺序不一定，因为map是无序的。

			2. 使用，数组或切片对key进行存储，并排序
				然后按照顺序打印value。
	*/

	map1 := make(map[int]string)
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "!!!"
	map1[4] = "hahaha"
	map1[5] = "6666"

	//1. 遍历map
	//1.1for range
	for k, v := range map1 {
		fmt.Println(k, v) //每次打印的顺序不一定
	}

	// 1.2
	for i := 0; i < len(map1); i++ { //如果key不是连续的数，会参杂空值
		fmt.Println(i, "对应的值是：", map1[i])
	}

	// 1.3利用切片当中介
	/*
		1. 获取所有的key，利用slice或array
		2. 对slice或array进行排序
		3. 遍历key，对应拿到value
	*/
	slice1 := make([]int, 0, len(map1))
	for k, _ := range map1 {
		slice1 = append(slice1, k)
	}
	fmt.Println(slice1)

	//给slice中介排序
	//冒泡排序，或者使用sort包下的排序方法
	sort.Ints(slice1) //sort.Ints是升序排列
	fmt.Println(slice1)

	for _, key := range slice1 {
		fmt.Printf("%d ---> %s\n", key, map1[key])
	}

	//延伸字母排序
	s1 := []string{"Apple", "Windows", "abc", "acc", "什么", "?"}
	fmt.Println(s1)
	sort.Strings(s1) //对字符串从小到大排序。
	fmt.Println(s1)
}
