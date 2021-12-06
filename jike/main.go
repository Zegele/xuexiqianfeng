package main

import (
	"fmt"
)

type Person struct {
	num    int
	name   string
	sex    string
	waist  float64
	weight float64
	bfr    float64
}

func main() {

	var list []Person

	for i := 1; i <= 2; i++ {
		var person Person

		person.num = i

		var name_in string
		fmt.Println("请输入您的姓名：")
		fmt.Scanln(&name_in)
		person.name = name_in

		var sex_in string
		fmt.Println("请输入您的性别：")
		fmt.Scanln(&sex_in)
		person.sex = sex_in

		var waist_in float64
		var weight_in float64

		if sex_in == "男" || sex_in == "man" {

			fmt.Println("请依次输入您腰围、体重：")
			fmt.Scanf("%f、%f\n", &waist_in, &weight_in)
			person.waist = waist_in
			person.weight = weight_in
			person.bfr = person.bfrman()

			list = append(list, person)

		} else if sex_in == "女" || sex_in == "woman" {
			fmt.Println("请依次输入您腰围、体重：")
			fmt.Scanf("%f、%f\n", &waist_in, &weight_in)
			person.waist = waist_in
			person.weight = weight_in
			person.bfr = person.bfrwoman()

			list = append(list, person)

		} else {
			fmt.Println("输入错误，请重新输入")
			return
		}
	}

	aBFR(list) //计算平均体脂率
}

func (p Person) bfrman() (bfr float64) {
	a := p.waist * 0.74         //79*0.74=58.46
	b := p.weight*0.082 + 44.74 //65*0.082+44.74=50
	fat_weight := a - b
	bfr = fat_weight / p.weight
	fmt.Printf("您的体脂率为%.2f%\n", bfr*100)
	return bfr //12.91%!
}

func (p Person) bfrwoman() (bfr float64) {
	a := p.waist * 0.74
	b := p.weight*0.082 + 34.89
	fat_weight := a - b
	bfr = fat_weight / p.weight
	fmt.Printf("您的体脂率为%.2f%\n", bfr*100)
	return bfr //28.06%!
}

func aBFR(list []Person) {
	var tbfr float64
	for i := 0; i < len(list); i++ {
		tbfr += list[i].bfr
	}

	abfr := int(tbfr*100) / len(list)
	fmt.Printf("平均体脂率为：%d%", abfr)
}
