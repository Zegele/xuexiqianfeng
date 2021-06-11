package main

import "fmt"

func main() {
	/*
		map: 映射，是一种专门用于存储键值对的集合。 属于引用类型。

		存储特点：
			A: 存储的是无序的键值对
			B： 键不能重复，并且和value值一一对应。
				map中的key不能重复，如果重复，那么新的value会覆盖原来的，程序不会报错。

		语法结构：
			1. 创建map
				var map1 map[key类型]value类型
					nil map  无法直接使用

				var map2 = make(map[key类型]value类型)

				var map3 = map[key类型]value类型{key:value, key2:value2,...}

		每种数据类型：
			int: 0
			float: 0.0 -->0
			string: ""
			array:[00000]

			//引用类型创建时没有初始化，就是nil
			slice:nil 在后面赋值时，可以使用
			map:nil 创建好后，不可以赋值，不可使用。只是声明了
	*/

	//创建map

	var map1 map[int]string         //没有初始化，nil var 本来就是变量声明。所以这里只是声明了，并没有初始化
	var map2 = make(map[int]string) // make创建，并且已经初始化了的，所以不是nil。 map2 := make(map[int]string)    // slice1 := make([]int, 4, 5)
	map3 := map[string]int{"go": 98, "ptyhon": 87, "java": 79, "html": 93}
	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	//map1[1] = "hello" //panic: assignment to entry in nil map 没有办法向nil map中分配数据
	if map1 == nil {
		map1 = make(map[int]string) //如果map是 nil ，就先make创建下。这样就不是nil了。
		fmt.Println(map1 == nil)
	}
}
