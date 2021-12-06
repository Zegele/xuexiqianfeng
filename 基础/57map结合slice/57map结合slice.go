package main

import "fmt"

func main() {
	/*
		map 和 slice的结合使用：
			1. 创建map用于存储人的信息
				name, age, sex, address

			2. 每个map存储一个人的信息

			3. 将这些map存入到slice中

			4. 打印遍历输出

	*/
	//slice0 := make([]string, 4)//用于每人的age，sex等
	//fmt.Printf("%T", slice0)

	map1 := make(map[string]string) //用于1个人
	//fmt.Printf(":%T:", map1)

	map1["name"] = "小明"
	map1["age"] = "18"
	map1["sex"] = "男"
	map1["address"] = "高新区"
	fmt.Println(map1)

	map2 := map[string]string{"name": "小李", "age": "19", "sex": "男", "address": "繁荣区"} //用于1个人

	slice1 := make([]map[string]string, 2)

	slice1[0] = map1
	slice1[1] = map2
	fmt.Println(slice1)

	for _, val := range slice1 {
		fmt.Println(val["name"], val["age"], val["sex"], val["address"])
	}

	fmt.Println("............................老师的做法")
	for i, v := range slice1 {
		fmt.Printf("第%d个人的信息是：\n", i+1)
		fmt.Printf("\t姓名： %s\n", v["name"])
		fmt.Printf("\t年龄： %s\n", v["age"])
		fmt.Printf("\t性别： %s\n", v["sex"])
		fmt.Printf("\t地址： %s\n", v["address"])
	}

}
