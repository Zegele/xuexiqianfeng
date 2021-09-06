package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		time包：
			1年 = 365天， day
			1天 = 24小时， hour
			1小时 = 60分钟， minute
			1分钟 = 60秒， second
			1秒钟 = 1000毫秒， millisecend
			1毫秒 = 1000微秒， microsecond -> μs
			1微秒 = 1000纳秒， nanosecond -> na
			1纳秒 = 1000皮秒， picosecond -> ps

	*/
	// 1.获取当前的时间
	t1 := time.Now()
	fmt.Printf("%T\n", t1) //time.Time 结构体类型
	fmt.Println(t1)        //2021-09-06 21:37:04.9135578 +0800 CST m=+0.003103401  年 月 日 时 分 秒 时区

	// 2.获取指定的时间
	t2 := time.Date(2008, 6, 7, 8, 9, 10, 11, time.Local) //2008-06-07 08:09:10.000000011 +0800 CST
	// 纳秒后跟着时区， time.Local 意思是当地的时区
	fmt.Println(t2)

	// 3.格式化时间样子 （其实就是time -> string之间的转换）

	/*
		time -> string
		t1.Format("2006-1-2 15:04:05")
		母板的日期必须是固定的：06-1-2-3-4-5
	*/
	s1 := t1.Format("2006年1月2日 15：04：05") //把t1的时间，按照字符串的格式，改造完成后，赋值给s1
	//时间的母板是不能变的，1月，2日，下午3点，4分，5秒，06年
	fmt.Println(s1)
	s2 := t1.Format("2006-1-2")
	fmt.Println(s2)
	/*
		   string -> time
			time.Parse("模板"，str) --> time, err 返回一个time类型，一个err
						模板	  要转换的对象
	*/
	s3 := "2022年1月1日"                      //string
	t3, err := time.Parse("2006年1月2日", s3) //1月与01月不匹配
	if err != nil {
		fmt.Println("err:", err)
	}
	fmt.Println(t3)

	fmt.Printf("%T\n %v\n", t1.String(), t1.String()) //t1.String()把time类型直接变为string类型，输出。

	// 4.根据当前时间，获取指定的内容
	year, month, day := t1.Date() //年，月，日
	fmt.Println(year, month, day)
	hour, min, sec := t1.Clock() //时，分，秒
	fmt.Println(hour, min, sec)

	year2 := t1.Year() //年
	fmt.Println("年：", year2)
	fmt.Println(t1.YearDay()) //t1.YearDay() 今年过去了多少天

	month2 := t1.Month()
	fmt.Println("月：", month2)
	fmt.Println("日：", t1.Day())
	fmt.Println("时：", t1.Hour())
	fmt.Println("分：", t1.Minute())
	fmt.Println("秒：", t1.Second())
	fmt.Println("纳秒：", t1.Nanosecond())

	fmt.Println(t1.Weekday()) //t1.Weekday() 星期几

	//5. 时间戳 指定的日期，距离1970年1月1日0点0分0秒 UTC时区 的时间差值：秒，纳秒
	t4 := time.Date(1970, 1, 1, 9, 0, 0, 0, time.UTC) //time.UTC 标准时区 如果是time.Local当地时区，那得加8小时 因为北京时区是东8区
	fmt.Println(t4)
	timeStamp1 := t4.Unix() //秒的差值
	fmt.Println(timeStamp1) //3600 秒
	timeStamp2 := t1.Unix() //当前时间与1970的时间插值，单位秒
	fmt.Println(timeStamp2) //1630938078

	timeStamp3 := t4.UnixNano() //单位为纳秒的差值
	fmt.Println(timeStamp3)     //3600 000 000 000 纳秒
	timeStamp4 := t1.UnixNano()
	fmt.Println(timeStamp4) //1630938078026636900

	//6.时间间隔
	t5 := t1.Add(1000) //Add 里的时间是纳秒为单位的，也直接写 time.Minute 之类的，就是一分钟
	// Add() 给时间加 小时，分，秒，等
	fmt.Println(t1)
	fmt.Println(t5)
	fmt.Println(t1.Add(24 * time.Hour)) //如果是负数，就是之前的时间。

	t6 := t1.AddDate(1, 0, 1)
	// AddDate(year, months, days) 给时间加年，月，日
	fmt.Println(t6)

	//两个时间的间隔 单位：会显示最简单位
	d1 := t5.Sub(t1)
	fmt.Println(d1)

	d2 := t1.Sub(t5)
	fmt.Println(d2)

	// 7. 睡眠
	time.Sleep(4 * time.Second) //time.Sleep()让当前程序进入睡眠状态。
	fmt.Println("main..over")

	//睡眠（1-10）的随机数
	rand.Seed(time.Now().UnixNano())
	randNum := rand.Intn(10) + 1 //rand.Intn(10) 这个范围是0-9
	fmt.Printf("%T\n", randNum)
	fmt.Println(randNum)
	time.Sleep(time.Duration(randNum) * time.Second) //time.Duration(randNum)把randNum int类型转化为 Duration int64类型
	fmt.Println("睡好没？")
}
