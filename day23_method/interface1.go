package main

import "fmt"

// 接口的定义与方法实现如下

//type 接口名字 interface {
//	方法1([参数列表]) [返回值]
//	方法2([参数列表]) [返回值]
//   .....
//   方法n ([参数列表]) [返回值]
// }

//实现接口方法的语法格式如下
//func (变量名 结构体类型) 方法1([参数列表]) [返回值]  {
//	// 方法体
//}

// 具体使用方式如下

// 定义接口

type Phone interface {
	call()
}

type AndroidPhone struct {
}

type Iphone struct {
}

// 实现接口的方法

func (a AndroidPhone) call() {
	fmt.Println("我是安卓手机，可以打电话了")
}

func (i Iphone) call() {
	fmt.Println("我是苹果手机，可以打电话了")
}

func main() {
	fmt.Println(" 接口的定义与实现")
	// 定义接口类型的变量
	var phone Phone
	phone = new(AndroidPhone)
	fmt.Printf("%T ,%v,%p \n", phone, phone, &phone)
	// 输出
	// *main.AndroidPhone ,&{},0xc000010200
	phone.call()
	// 我是安卓手机，可以打电话了
	phone = new(Iphone)
	fmt.Printf("%T ,%v,%p \n", phone, phone, &phone)
	phone.call()
	// 输出
	//*main.Iphone ,&{},0xc00008e1e0
	//我是苹果手机，可以打电话了
	//new(AndroidPhone)以及new(IPhone)可以通过隐式实现接口直接赋值给接口变量phone。

}
