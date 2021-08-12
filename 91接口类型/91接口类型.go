package main

import "fmt"

func main() {
	/*
			接口：interface
				在Go中，接口是一组方法签名。

				当某个类型为这个接口中的所有方法提供了方法的实现，他被称为实现接口。

				Go语言中，接口和类型的实现关系，是非侵入式的。

		1. 当需要接口类型的对象时，可以使用任意实现类对象代替。
			//实现类，谁实现了接口中的方法，谁就是实现类。
		2. 接口对象不能访问实现类中的属性

		多态：一个十五的多种形态
			go语言通过接口模拟多态

		理解多态：
		就是一个接口的实现
			1. 看成实现本身的类型，能够访问实现类中的属性和方法。
				如本身就是U盘，就能访问U盘本身的属性和方法。

			2. 看成是对应的接口类型，那就只能够访问接口中的方法。

		接口的用法：
			1. 一个函数如果接受接口类型作为参数，那么实际上可以传入该接口的任意实现类型对象作为参数。 返回值同理
				//如41行：testInterface(m1) 参数是USB接口类型，传入的参数是接口实现类的对象，m1

			2. 定义一个类型为接口类型，实际上可以赋值为任意实现类的对象。
				//如47：var usb USB //定义usb是USB接口类型
				//	usb = m1 //将m1这个实现类赋值给该接口类型

		鸭子类型：


	*/
	//1.创建Mouse类型
	m1 := Mouse{name: "小黑"}
	fmt.Println(m1.name)
	//2.创建FlashDisk
	f1 := FlashDisk{name: "金士"}
	fmt.Println(f1.name)

	testInterface(m1) //该函数 让m1实现testInterface里的两个方法 （全部方法）
	testInterface(f1)

	var usb USB //定义usb是USB接口类型
	usb = m1    //m1实现了USB接口的方法。 所以m1可以代替usb
	usb.start() //让m1实现接口中的某一个方法（某一个方法）
	usb.end()

	f1.deleteData()
	//usb.deleteData()//不可实现U盘自己的方法

	var arr [3]USB
	arr[0] = m1
	arr[1] = f1
	fmt.Println(arr)

}

//1.定义接口
type USB interface {
	start() //USB设备开始工作 方法的声明 方法的名字(方法的参数)(方法的返回值)
	end()   //USB设备结束工作
}

//2.实现类
type Mouse struct {
	name string
}

type FlashDisk struct {
	name string
}

func (m Mouse) start() {
	fmt.Println(m.name, "鼠标准备就绪，可以开始工作，点点点。。。")
}

func (m Mouse) end() {
	fmt.Println(m.name, "结束工作，可以安全退出。。")
}

func (f FlashDisk) start() {
	fmt.Println(f.name, "闪盘准谱开始工作，可以进行数据存储。。")
}
func (f FlashDisk) end() {
	fmt.Println(f.name, "结束，可以弹出。。")
}

//3.测试方法
func testInterface(usb USB) { //参数是接口
	usb.start()
	usb.end()
}

func (f FlashDisk) deleteData() {
	fmt.Println(f.name, "删除数据。。。")
}
