package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	/*
		遍历文件夹：

	*/
	dirname := "进阶"
	listFiles(dirname, 0)
}

func listFiles(dirname string, level int) {
	//level用来记录当前递归的层次，生成带有层次感的空格
	s := "|--"
	for i := 0; i < level; i++ {
		s = "|  " + s
	}
	fileInfos, err := ioutil.ReadDir(dirname) //第一个返回值是一个切片，包含所有文件信息
	if err != nil {
		log.Fatal(err)
	}
	for _, fi := range fileInfos {
		filename := dirname + "/" + fi.Name()
		fmt.Printf("%s%s\n", s, filename)
		if fi.IsDir() { //如果是文件夹
			//递归调用方法
			listFiles(filename, level+1)
		}
	}

}
