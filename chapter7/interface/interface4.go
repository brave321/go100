package main

import "fmt"

// 多态

//多态就是事物的多种形态，
//Go语言中的多态性是在接口的帮助下实现的——定义接口类型，创建实现该接口的结构体对象。

type Income interface {
}

func main() {
	fmt.Println("hello world")
}
