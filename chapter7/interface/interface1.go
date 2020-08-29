package main

import (
	"fmt"
)

// 定义接口的语法格式

//type 接口名字 interface {
//	方法1 ([参数列表]) [返回值]
//	方法2 （[参数列表]）[返回值]
//	....
//	方法n ([参数列表]) [返回值]
//}

// 实现接口方法的语法格式如下
//
//func (变量名 结构体类型) 方法1 ([参数列表返回值] ) [返回值] {
//  // 方法体
//}

//示例

type Phone interface {
	call()
}

type AndriodPhone struct {
}

type IPhone struct {
}

func (a AndriodPhone) call() {
	fmt.Println("安卓手机")
}

func (i IPhone) call() {
	fmt.Println("苹果手机")
}

func main() {
	fmt.Println("接口的定义与实现")
	var phone Phone
	phone = new(AndriodPhone)
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	fmt.Println("=============")
	phone = IPhone{}
	fmt.Print("%T,%v,%p \n", phone, phone, &phone)
	phone.call()

}
