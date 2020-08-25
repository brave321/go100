package main

import "fmt"

// 结构体做为函数参数，若复制一份到函数中，对函数参数进行修改
// 不会影响到实际参数

type Human struct {
	name string
	age  int8
	sex  byte
}

func changename(h Human) {
	h.name = "Daniel"
	h.age = 23
	fmt.Printf(" 函数体修改后为： %T,%v, %p \n", h, h, &h)

}

func main() {
	//1、 初始化human
	h1 := Human{"cc", 35, 1}
	fmt.Printf("h1 %T,%v, %p \n", h1, h1, &h1)
	fmt.Println("=========")
	// 复制结构体对象
	h2 := h1
	h2.name = "David"
	h2.age = 30
	fmt.Printf("h2 修改后：%T,%v %p \n", h2, h2, &h2)
	fmt.Println("=========")
	// 3 将结构体对象作为参数传递
	changename(h1)
	fmt.Printf("h1 %T,%v, %p \n", h1, h1, &h1)
	// 输出
	// 函数体修改后为： main.Human,{Daniel 23 1}, 0xc00000c120
	// h1 main.Human,{cc 35 1}, 0xc00000c060
	// h1 传入函数进行修改，依然没有对h1 造成影响，结构体是值类型

}
