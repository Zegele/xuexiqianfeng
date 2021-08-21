package main

import (
	"fmt"
	"net"
)

func main() {
	addr, err := net.LookupHost("www.baidu.com")
	fmt.Println(err)  //nil
	fmt.Println(addr) //[14.215.177.38 14.215.177.39]

	addr1, err1 := net.LookupHost("www.baidu1246498167.com")
	fmt.Println(err1) //lookup www.baidu1246498167.com: no such host

	if ins, ok := err1.(*net.DNSError); ok { //断言这个错误是不是net.DNSError中的错误
		if ins.Timeout() {
			fmt.Println("操作超时。。。")
		} else if ins.Temporary() {
			fmt.Println("临时性错误")
		} else {
			fmt.Println("通常错误")
		}

	}
	fmt.Println(addr1) //[]
}
