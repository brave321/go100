package main

import (
	"fmt"
)

// 接口的定义与实现

// 1、定义接口的语法格式

//type 接口名称 interface {
//	方法1([参数列表]) [返回值]
//	方法2([参数列表]) [返回值]
//
//}

// 2、定义接口方法的语法格式

//func (变量名 结构体类型) 方法1([参数列表])  [返回值]  {
//    // 方法体
//}

// 3 示例 通过传入不同的方法，打印不同的函数名称

type Phone interface {
	call()
}

type AndroidPhone struct {
}

type Iphone struct {
}

func (a AndroidPhone) call() {
	fmt.Println("安卓手机")
}

func (i Iphone) call() {
	fmt.Println("苹果手机")
}

func main() {
	// 定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	// 输出
	//*main.AndroidPhone,&{},0xc00008e1e0
	//安卓手机
	phone = new(Iphone)
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()
	// 输出
	//*main.Iphone,&{},0xc000010200
	//苹果手机
	phone = Iphone{}
	fmt.Printf("%T,%v,%p \n", phone, phone, &phone)
	phone.call()

}
