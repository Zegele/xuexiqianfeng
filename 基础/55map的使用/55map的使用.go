package main

import "fmt"

func main() {
	/*
		map语法结构：
		1. 创建map 54课

		2. 添加\修改
			map[key] = value
			如果key不存在，就是添加数据
			如果key存在，就是修改数据

		3. 获取
			map[key] ---> value //一般不这样获取，因为无法判断key是否真的存在map中。
			value, ok := map[key]
				如果ok为true，说明key真实存在，value就是对应的值。
				如果ok为false，说明key不存在，value就是值类型的默认值。

		4. 删除数据：
			delete(map, key)
				如果key存在，就可以直接删除
				如果key不存在，删除失败。但也不影响map。

		5. 长度：
			len()

		注意map是无序的，每次打印map，结果的顺序可能不同。 map必须通过key值对应相应的value。
	*/

	var map1 map[int]string //只是声明了map1，还不能直接用。
	if map1 == nil {
		map1 = make(map[int]string) //创建了，也就是map1已经初始化了，可以用了。
	}

	// 2. 存储键值对到map中
	map1[1] = "hello"
	map1[2] = "world"
	map1[3] = "haha"
	map1[4] = "ooo"
	map1[5] = "" //怎么判断key 5 是否真的存在，而不是输出的默认值。

	// 3. 获取数据，根据key获取对应的value值
	// 根据key获取对应的value，如果key存在，获取数值，如果key不存在，获取的是value值类型的零值。
	fmt.Println(map1)
	fmt.Println(map1[4])  // 根据key为4，获取对应的value值。
	fmt.Println(map1[40]) // 没有key为40，得到的是value值类型的默认值。

	v1, ok := map1[40]
	if ok {
		fmt.Println("对应的数值是：", v1)
	} else {
		fmt.Println("操作的key不存在，获取到的是value类型的零值：", v1)
	}

	// 4. 修改数值
	fmt.Println(map1)
	map1[3] = "你好"
	fmt.Println(map1)

	// 5. 删除数据
	delete(map1, 3)
	fmt.Println(map1)
	delete(map1, 30)
	fmt.Println(map1)

	// 7. 长度 //有几个键值对
	fmt.Println(len(map1))

}
