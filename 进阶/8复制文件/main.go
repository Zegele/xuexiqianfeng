package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {

	/*
			拷贝文件：
		方法1：自己写拷贝代码
		方法2：io.Copy
		方法3：ioutil包
			使用ioutil包中的 ioutil.WriteFile() 和 ioutil.ReadFile()，
			但由于使用一次性读取文件，再一次性写入文件的方式，所以该方法不适用于大文件。容易内存溢出。

	*/

	srcFile := "D:/gogo/src/xuexiqianfeng/进阶/8复制文件/test.jpeg"   //源文件地址
	destFile := "D:/gogo/src/xuexiqianfeng/进阶/8复制文件/test1.jpeg" //目标文件地址

	//total, err := CopyFile1(srcFile, destFile) //调用方法一
	//total, err := CopyFile2(srcFile, destFile) //调用方法二
	total, err := CopyFile3(srcFile, destFile) //调用方法三
	fmt.Println(total, err)

}

// 方法一： 该函数用于通过IO操作实现文件的拷贝，参数是两个文件的路径，返回值是拷贝的总数量（字节），error
func CopyFile1(srcFile, destFile string) (int, error) {
	file1, err := os.Open(srcFile) // 这里对源文件，只需要只读
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()

	//读写操作
	bs := make([]byte, 1024, 1024)
	n := -1 // 读取的数量
	total := 0
	for {
		n, err = file1.Read(bs)
		if err == io.EOF || n == 0 {
			fmt.Println("拷贝完毕！")
			break
		} else if err != nil {
			fmt.Println("报错了！！！")
			return total, err
		}
		total += n
		file2.Write(bs[:n]) //n 读到n个就写出n个，如果n是100个，但容量是1024个，也不用担心，把之前的数据重复写出。
	}
	return total, nil
}

//方法二：io.Copy 函数 (推荐)
func CopyFile2(srcFile, destFile string) (int64, error) {
	file1, err := os.Open(srcFile)
	if err != nil {
		return 0, err
	}
	file2, err := os.OpenFile(destFile, os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer file1.Close()
	defer file2.Close()
	return io.Copy(file2, file1) //io.Copy(目标文件，源文件)（int64, err） 把源文件拷贝到目标文件
}

//方法三：ioutil包下的ioutil.WriteFile() 和 ioutil.ReadFile()
func CopyFile3(srcFile, destFile string) (int, error) {
	bs, err := ioutil.ReadFile(srcFile) //ioutil.ReadFile(源文件)（字节切片，error）  ioutil.ReadFile()函数读取源文件，并返回字节切片和err。
	if err != nil {
		return 0, err
	}
	err = ioutil.WriteFile(destFile, bs, 0777)
	if err != nil {
		return 0, err
	}
	return len(bs), nil
}

/*
扩展：
CopyN(dst, src, n) 为复制src中n个字节到dst
CopyBuffer(dst, src, buf) 为指定一个buf缓存区，以这个大小完全复制。
*/
