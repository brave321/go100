package main

import "fmt"

// 什么是指针
// 指针是存储变量内存地址的变量

// 指针的声明

// 指针变量的类型为 *T ，该指针指向一个T 类型的变量, 比如 int、string
//& 用于获取操作符号的变量地址 .

func test1() {
	b := 255
	s1 := "hello"
	var a *int = &b
	var s2 *string = &s1
	fmt.Printf("type of a is %T\n", a)
	fmt.Println("address of b is ", a)
	fmt.Printf("type of s2 is %T\n", s2)
	fmt.Println("address of s2 is ", s2)

	// 使用指针来修改b 的值
	*a++
	fmt.Println("new value of b is", b)

}

func main() {
	fmt.Println("指针")
	test1()
}
