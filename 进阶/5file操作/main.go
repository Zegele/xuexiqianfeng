package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

func main() {
	/*
		首先，file类是在os包中的，封装了底层的文件描述符和相关信息，同时封装了Read和Write的实现。
		一、FileInfo接口
		FileInfo接口中定义了File信息相关的方法。
		type FileInfo interface{
			name() string // base name of the file 文件名.扩展名 aa.txt  name函数，返回值是字符串类型
			Size() int64 // 文件大小，单位：字节 abcd 4个字节
			Mode() FileMode // 文件权限 -rw-rw-rw-
			ModTime() time.Time // 修改时间 2020-08-10 16：12：30 +8000 CST
			IsDir() bool // 是否文件夹
			Sys() interface{} // 基础数据接口（can return nil）

		通过 func Stat(name string)(FileInfo, error)函数，可以获取FileInfo返回值。
	*/

	// 1. 获取FileInfo信息 （FileInfo是一个接口）
	// 通过 os.Stat 这个Stat函数就可以获得FileInfo，通过FileInfo就可以获得文件信息。
	fileInfo, err := os.Stat("D:/gogo/src/xuexiqianfeng/进阶/5file操作/a.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Printf("%T\n", fileInfo)

	// 获取文件名（查看文件名）
	fmt.Println(fileInfo.Name())
	// 获取（查看）文件大小
	fmt.Println(fileInfo.Size())
	//获取（查看）文件权限
	fmt.Println(fileInfo.Mode()) //-rw-rw-rw-
	//获取（查看）文件修改时间
	fmt.Println(fileInfo.ModTime())
	//产看是否文件夹
	fmt.Println(fileInfo.IsDir())

	/*
			文件操作：
			1. 路径：
				相对路径： relative
					ab.txt
					相对于当前工程
				绝对路径：absolute
					D:/gogo/src/xuexiqianfeng/进阶/5file操作/a.txt

			. 表示当前目录
			.. 表示上一层目录

			2. 创建文件夹，如果问价夹存在，创建失败
		 		os.MkDir(), 创建一层
		 		os.MkDirAll(), 创建多层

			3. 创建文件，Create采用模式0666 （任何人都可读写，不可执行）创建一个名为name的文件，如果文件已存在会截断它（成为空文件）
				os.Create() ，创建文件

			4. 打开文件，让当前的程序，和执行的文件之间建立一个连接。
				os.Open(filename)
				os.OpenFile(filename, mode, perm)

			5. 关闭文件：程序和文件之间的连接断开。
				file.Close()

			6. 删除文件：慎用，慎用，再慎用！(代码删除是不会有回收站的，删除就没有了，所以要慎重！)
				os.Remove() //删除空文件和目录
				os.RemoveAll() //删除所有

	*/
	// 1.路径
	fileName1 := "D:/gogo/src/xuexiqianfeng/进阶/5file操作/a.txt"
	fileName2 := "ab.txt"
	fmt.Println(filepath.IsAbs(fileName1)) // filepath文件路径的结构体 IsAbs(路径) 函数，表示该路径是否为绝对路径
	fmt.Println(filepath.IsAbs(fileName2))
	fmt.Println(filepath.Abs(fileName1)) //Abs(路径)，把路径变为绝对路径
	fmt.Println(filepath.Abs(fileName2)) //相对该工程，本软件相对该工程就是xuexiqianfeng ，Sublime软件相对该工程是 5file操作
	fmt.Println(filepath.Abs("123"))     //不需要这个路径真实有个文件啥的。

	fmt.Println("获取父目录", path.Join(fileName1, "..")) //Join(路径, "..") 获取该路径的上一层路径  .. 上一层路径， .本层路径

	//2. 创建目录（文件夹）
	//创建单层目录
	dir := os.Mkdir("D:/gogo/src/xuexiqianfeng/进阶/5file操作/abc", os.ModePerm) //os.Mkdir(路径, 权限) // ModePerm 是系统设置的常数 就表示0777
	//Mkdir() 创建一层文件夹

	if dir != nil {
		fmt.Println("err:", dir)
		return
	}
	fmt.Println("文件夹创建成功")

	// 创建多层目录
	err1 := os.MkdirAll("D:/gogo/src/xuexiqianfeng/进阶/5file操作/cd/ef/gh", os.ModePerm)
	if err != nil {
		fmt.Println("err:", err1)
		return
	}
	fmt.Println("多层文件创建成功！")

	//3. 创建文件

	//3.1 创建绝对路径的文件
	file1, err := os.Create("D:/gogo/src/xuexiqianfeng/进阶/5file操作/1.txt") // os.Create 函数创建文件
	//注意：使用os.Create函数创建文件，如果该文件存在，再创建，就会变为空文件。
	//通过返回值file1，就可以对文件进行读，写等操作。
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file1) //&{0xc0000a8000} 是个指针类型

	//3.2 创建相对路径的文件
	file2, err := os.Create(fileName2) //创建相对路径的文件，是以当前工程为参照的 golang软件以 xuexiqianfeng为工程的路径为参照，所以ab.txt在xuexiqianfeng下。
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file2)

	//4. 打开文件：打开文件，就是让当前该程序与文件产生一个连接。然后就可以通过程序操作文件内从数据了。
	//4.1 只读模式打开文件。
	file3, err := os.Open(fileName1) //os.Open() 函数是打开文件，然后 只读 权限的。

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file3)

	//4.2 指定模式打开文件。
	file4, err := os.OpenFile(fileName1, os.O_RDONLY|os.O_WRONLY, os.ModePerm) // os.O_RDONLY(只读)|os.O_WRONLY（只写） == os.O_RDWR（可读可写）
	//os.OpenFile(文件名, 指定权限, 如果文件不存在新建文件的权限) //
	//第二参数：文件的打开方式 （意思是，是以什么权限打开）
	//const (
	//	// Exactly one of O_RDONLY, O_WRONLY, or O_RDWR must be specified.
	//	O_RDONLY int = syscall.O_RDONLY // open the file read-only.
	//	O_WRONLY int = syscall.O_WRONLY // open the file write-only.
	//	O_RDWR   int = syscall.O_RDWR   // open the file read-write.
	//	// The remaining values may be or'ed in to control behavior.
	//	O_APPEND int = syscall.O_APPEND // append data to the file when writing.
	//	O_CREATE int = syscall.O_CREAT  // create a new file if none exists.
	//	O_EXCL   int = syscall.O_EXCL   // used with O_CREATE, file must not exist.
	//	O_SYNC   int = syscall.O_SYNC   // open for synchronous I/O.
	//	O_TRUNC  int = syscall.O_TRUNC  // truncate regular writable file when opened.
	//)
	//第三参数是，当该文件名的文件不存在时，新建一个文件，该新建文件的权限。

	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println(file4)

	//5. 关闭文件
	file4.Close() // Close() 关闭文件

	//6. 删除文件或文件夹
	// 删除文件
	err = os.Remove("D:/gogo/src/xuexiqianfeng/进阶/5file操作/a.txt")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	fmt.Println("删除文件成功！")

	// 删除目录
	err = os.Remove("D:/gogo/src/xuexiqianfeng/进阶/5file操作/cd/ef/gh") //os.Remove() 使用该函数删除文件夹，该文件夹必须是空文件夹。
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("删除文件夹成功！")

	err = os.RemoveAll("D:/gogo/src/xuexiqianfeng/进阶/5file操作/cd") //os.RemoveAll() 使用该函数完全删除整个文件夹及内容。
	//慎重使用os.RemoveAll() 函数
	if err != nil {
		fmt.Println("err: ", err)
		return
	}
	fmt.Println("删除文件夹成功！")
}

/*
扩展
权限：Linux中，有2中文件权限表示方式，即“符号表示”和“八进制表示”。
1. 符号表示方式：
	- 	--- 	--- 	---
  type	owner	group	others
type中，-表示文件，d表示目录（文件夹） |连续符号
r表示可读，w表示可写，x表示可执行，如果没有该权限用-代替
例如： -rwxr-xr-x 表示：这是一个文件，对于拥有者可读可写可执行，对于组（group）可读可执行不可写，对于其他人(other)可读可执行不可写

2. 八进制表示方式：
r -> 004
w -> 002
x -> 001
- -> 000

例如：
0755 对拥有者可读可写可执行，对组可读可执行，对其他人可读可执行。
0777
0555
0444
0666 对拥有者，组，其他人，都是可读可写可执行了。
*/

/*
在课程6：30 有打开模式，File操作的文字
*/
