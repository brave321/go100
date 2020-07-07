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
	// 输出
	// 	type of a is *int
	// address of b is  0xc000016070
	// type of s2 is *string
	// address of s2 is  0xc000010200

	// 使用指针来修改b 的值
	*a++
	fmt.Println("new value of b is", b)
	// 输出
	// new value of b is 256

}

// 指针的零值

func test2() {
	a := 25
	var b *int
	if b == nil {
		fmt.Println("b is", b)
		b = &a // & 用于获取操作符号的变量地址  ,将a 的地址赋指给b
		fmt.Println("b after initialization is", b)
		// 输出
		// 		b is <nil>  b 初始化为nil
		// b after initialization is 0xc000016070

	}

}

// 指针的解引用
// 指针的解引用可以获取指针指向的变量的值  ,将a 解引用的语法是 *a

func test3() {
	b := 255
	a := &b
	fmt.Println("address of b is ", a)
	fmt.Println("value of b is", *a)
	// 输出
	// 	address of b is  0xc0000b4010
	// value of b is 255

}

// 向参数传递指针参数 ，如下向函数 val 传递指针参数 *int

func change(val *int) {
	*val = 55

}

// 不要向函数传递数组的指针，而应该使用切片
// 假如我们想要在函数内修改一个数组，并希望调用函数的地方能得到修改后的数组
// 一种解决方案是把一个指向数组的指针传递给这个函数

func modify(arr *[3]int) {
	(*arr)[0] = 90

}

// 使用切片获取  上述代码

func modify2(sls []int) {
	sls[0] = 900
}

// 指针运算

func calc() {
	b := [...]int{11, 45, 121}
	p := &b
	fmt.Println(b)
	fmt.Println(p, *p)
	p++

}
func main() {
	fmt.Println("指针")
	// test1()
	// test2()
	// test3()
	// a := 58
	// fmt.Println("value before is", a)
	// b := &a // 将a 的变量地址赋值给b
	// change(b)
	// fmt.Println("value after is", a)
	// 输出
	// 	value before is 58
	// value after is 55

	// a := [3]int{11, 989, 32}
	// modify(&a)
	// fmt.Println(a)
	// 输出

	// [90 989 32]

	// a := [3]int{11, 45, 98}
	// modify2(a[:])
	// fmt.Println(a)
	// 输出 [900 45 98]
	calc()
	// 输出 p++ (non-numeric type *[3]int) go 不支持指针运算

}
