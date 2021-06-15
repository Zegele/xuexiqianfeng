package main

import "fmt"

func main() {
	/*
		一、 数据类型：
			基本数据类型： int, float, string, bool
			复合数据类型： array, slice, map, function, pointer, struct..

				array: [size]数据类型
				slice: []数据类型
				map: map[key的类型]value的类型  //value的类型也可以是一个map

		二、 存储特点：
			值类型： int, float, string, bool, array, struct
			引用类型：slice, map
	*/

	map1 := make(map[int]string)
	map2 := make(map[string]float64)
	fmt.Printf("%T\n", map1)
	fmt.Printf("%T\n", map2)

	map3 := make(map[string]map[string]string) //map中又有一个map
	m1 := make(map[string]string)
	m1["name"] = "小李"
	m1["age"] = "31"
	m1["salary"] = "3000"
	map3["hr"] = m1

	m2 := make(map[string]string)
	m2["name"] = "小刘"
	m2["age"] = "32"
	m2["salary"] = "8000"
	map3["总经理"] = m2

	fmt.Println(map3)

	fmt.Println("..................")
	map4 := make(map[string]string)
	map4["小王"] = "高富帅"
	map4["翠花"] = "白富美"

	fmt.Println(map4)

	map5 := map4
	fmt.Println(map5)

	map5["go"] = "工具人"
	fmt.Println(map5)
	fmt.Println(map4)
}
