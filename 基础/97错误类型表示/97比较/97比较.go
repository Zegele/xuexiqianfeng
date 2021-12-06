package main

import (
	"fmt"
	"path/filepath"
)

func main() {
	/*
		如果返回的错误是我们比较关心的错误，我们就可以直接比较，该错误是不是和我们关心的错误是同一个类型的。就用这种方法。
	*/
	files, err := filepath.Glob("test.txt")
	fmt.Println(err)
	if err != nil && err == filepath.ErrBadPattern { // filepath.ErrBadPattern 就是将filepath中的错误和想要对比的错误进行比较
		//如果类型相同，就没啥。如果不同就会返回错误：ErrBadPattern
		fmt.Println(err)
		return
	}
	fmt.Println("files:", files)
}
