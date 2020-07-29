package main

import (
	"fmt"
	"strings"
	"time"
)

// 匿名函数

// go 支持匿名函数，如果某个函数只是希望使用一次，可以考虑使用匿名函数，匿名函数可以实现多次调用
// 定义匿名函数的同时 调用

// 必包
// 关于闭包的理解，
// 必包把返回的函数赋给了一个变量，虽然函数在执行完一瞬间会销毁其执行环境
// 如果有闭包的话，必包会保存外部函数的活动对象
// 闭包会一直存在内存中，垃圾收集器不会销毁闭包占用的内存

// 闭包1

func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

//闭包2
func miu(num int) func(int) int {
	return func(i int) int {
		num += i
		return num
	}
}

// 闭包 3 ，入参是一个字符串 ，返回值是一个函数
// return 是一个函数
func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}

}

// 闭包4 返回值是一个函数

// 入参是一个int 类型的数值，返回值是2个函数,返回值是int 类型

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i // base=base +i
		return base
	}
	sub := func(i int) int {
		base -= i // base = base -i
		return base
	}
	return add, sub
}

func main() {
	fmt.Println("hello world")
	// a := func() {
	// 	fmt.Println(1)
	// }
	// a()

	// 匿名函数1
	// res1 := func(n1 int, n2 int) int {
	// 	return (n1 + n2)
	// }(10, 20)
	// fmt.Println("res1", res1)

	// 输出
	//res1 30

	// 匿名函数2 将匿名函数赋值给一个变量

	// a := func(n1 int, n2 int) int {
	// 	return n1 - n2
	// }
	// res2 := a(20, 19)
	// fmt.Println("res2", res2)
	// 输出
	//res2 1
	// go 闭包的引用
	// func A() func(aa int) int {
	// 	sum := 0
	// 	return func(cc int) int {
	// 		sum += cc
	// 		fmt.Println("aa=", aa, "bb=", bb, "	sum=", sum)
	// 		return sum
	// 	}
	// }
	// 闭包1
	// tmp1 := add(10)
	// fmt.Println(tmp1(1), tmp1(10))
	// 输出
	// 11 21

	// num1 := miu(10)
	// fmt.Println(num1(10))
	// 输出
	//20

	// 闭包的例子3
	// fun1 := makeSuffixFunc(".jpg")
	// fmt.Println(fun1)

	// fun2 := makeSuffixFunc(".pdf")
	// fmt.Println(fun1("tes1t"))
	// fmt.Println(fun2("test2"))

	// 输出
	// tes1t.jpg
	// test2.pdf

	//闭包4
	f1, f2 := calc(13)
	fmt.Println(f1, f2)
	fmt.Println(f1(10), f2(12)) // 13 +10 , 13-1

	// 闭包不5
	for i := 0; i < 5; i++ {
		go func() {
			fmt.Println(i)
		}()
	}
	time.Sleep(time.Second)
}
