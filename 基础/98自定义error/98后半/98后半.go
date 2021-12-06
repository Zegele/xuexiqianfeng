package main

import "fmt"

func main() {
	/*

	 */
	length, width := -5.5, -4.4
	area, err := rectArea(length, width)
	if err != nil {
		fmt.Println(err)
		if err, ok := err.(*areaError); ok { //断言 此error是不是*areaError类型，并且可以具体看err中的数据
			if err.lengthNegative() { //调用lengthNegative，返回bool，如果true，继续执行。
				fmt.Printf("error:长度：%.2f，小于零\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error:宽度：%.2f，小于零\n", err.width)
			}

		}
		return
	}
	fmt.Println("矩形的面积是：", area)

}

type areaError struct {
	msg           string  //必须有个string，而且必须是第一个。 //错误的描述
	length, width float64 //发生错误的时候，长、宽的长度
}

func (e *areaError) Error() string {
	return e.msg
}

func (e *areaError) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError) widthNegative() bool {
	return e.width < 0
}

func rectArea(length, width float64) (float64, error) {
	msg := ""
	if length < 0 {
		msg = "长度为负了"
	}
	if width < 0 {
		if msg == "" {
			msg = "宽度为负了"
		} else {
			msg += "，宽度也为负了"
		}
	}

	if msg != "" {
		return -1, &areaError{msg, length, width}
	}
	return length * width, nil
}
