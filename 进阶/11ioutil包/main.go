package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	/*
		ioutil 包：
			ReadFile()
			WriteFile()
			ReadDir()
			...
	*/

	//1. 读取文件中的所有的数据
	fileName := "进阶/11ioutil包/aa.txt"
	data, err := ioutil.ReadFile(fileName) //data是byte（字节）类型的切片
	fmt.Println(err)
	fmt.Printf("%T, %d,%q\n", data, data, data)
	fmt.Println(string(data)) //字节类型的切片可以直接转为string类型？？

	//2.写出数据（写之前会先清空）
	fileName1 := "进阶/11ioutil包/bbb.txt"
	s1 := "helloword"                                            //要写入的内容
	err1 := ioutil.WriteFile(fileName1, []byte(s1), os.ModePerm) //第一个参数是要写入的文件名，第二个参数是，写入内容，必须是字节类型的切片，第三个参数是，如果该文件不存在，则创建该文件的权限，os.ModePerm即0777。
	//返回值是个错误
	fmt.Println(err1)

	//3. ReadAll()  该方法，可读取的类型比较多，只要是io.Reader的实现类都可以使用该方法读取。
	s2 := "好好学习，天天向上！"
	r1 := strings.NewReader(s2)
	data2, err2 := ioutil.ReadAll(r1) //参数需要io.Reader的实现类对象
	fmt.Println(err2)
	fmt.Println(string(data2))
	//表示读取r1这个对象，所指向的数据。在这距离r1指向的就是s2，所以读取的就是s2中的数据。

	data3, err3 := ioutil.ReadAll(r1) //尝试再读取一次，但没有数据了。也就是说strings.NewReader(s2)使用一次后，就会垃圾回收了。
	fmt.Println(err3)
	fmt.Println(data3)

	//4. ReadDir() 读取文件夹 （只读取一层文件夹中的内容）
	dirName := "进阶/11ioutil包"
	fileInfos, err4 := ioutil.ReadDir(dirName) //返回值是一个切片(包含所有文件的信息，如文件名，文件类型，生成时间等)和err
	if err4 != nil {
		fmt.Println(err4)
		return
	}
	fmt.Println(fileInfos)
	fmt.Println(len(fileInfos))
	for i := 0; i < len(fileInfos); i++ {
		//fmt.Printf("%T\n", fileInfos[i])
		fmt.Printf("第 %d 个：名称：%s, 是否是目录：%t\n", i, fileInfos[i].Name(), fileInfos[i].IsDir())
	}

	//5. 创建临时文件夹和临时文件
	//临时文件夹
	dir5, err5 := ioutil.TempDir("进阶/11ioutil包", "Test") //第一个参数：在该目录下创建这个临时文件夹，第二个参数：该临时文件的名字开头前缀，返回值是该临时文件夹的路径，和err
	if err5 != nil {
		fmt.Println(err5)
		return
	}

	defer os.Remove(dir5) //删除文件夹，就是删除路径
	fmt.Println(dir5)     //进阶\11ioutil包\Test167978347

	//临时文件
	file, err6 := ioutil.TempFile(dir5, "Test.txt") //第一个参数在该文件夹下，创建一个文件，第二个参数：该文件名的前缀，返回值是该文件的路径和err
	if err6 != nil {
		fmt.Println(err6)
		return
	}
	file.Close()                 //关闭文件后，才能执行删除文件。不明白为什么教程中没有这一步就执行删除了，而我这里得先关闭文件。先跟着脚印走，然后自己细细看，仔细琢磨。
	defer os.Remove(file.Name()) //file是路径，file.Name()才是文件名

	fmt.Println(file)
}
