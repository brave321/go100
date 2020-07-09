package main

import "fmt"

// 符合数据类型
// 数组是同一类型元素的集合

// 数组的声明 [n]T n 表示数组中元素的数量，T 代表每个元素的类型
//[n]T 个数+类型

func arraytest1() {
	// 数组的声明
	var a [3]int // 表示声明了一个长度为3的类型
	fmt.Println(a)
	a[0] = 11 // 赋值，从0开始
	a[1] = 22
	a[2] = 78
	fmt.Println(a) //[11 22 78]

	//短变量声明
	b := [3]int{12, 37, 89}
	fmt.Println(b) //[12 37 89]

	// 短变量声明，赋部分数值
	c := [3]int{11}
	fmt.Println(c)

	// 短变量声明 忽略声明数组的长度 使用 ... 代替
	d := [...]int{11, 2121, 3}
	fmt.Println(d)

	// 数组的大小是类型的一部分，数组不能比较大小
	// e1 := [3]int{11, 12, 3}
	// var f1 [5]int
	// e1 = f1 // 编译报错 cannot use f1 (type [5]int) as type [3]int in assignment
	// fmt.Println(e1)
}

//数组是值类型
// go 中的数组是值类型而不是引用类型，这意味着当数赋值给一个新的变量的时候，
// 该变量 会得到一个原始数组的一个副本，若对新变量进行更改，则不对影响原始数组

func arraytest2() {
	a := [...]string{"USA", "China", "India", "Germany", "France"}
	b := a
	fmt.Println("before a", a)
	// fmt.Println(b)
	// 修改b 值
	b[0] = "Singapore"
	fmt.Println("after a", a)
	fmt.Println(b) //修改b 不影响a

}

//数组作为参数传递给函数时，他们是按值传递，而原始数组保持不变

func arraytest3(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)

}

func main() {
	// arraytest1()
	// arraytest2()
	num := [...]int{1, 4, 5, 567, 9}
	fmt.Println("before passing to function ", num)
	arraytest3(num)
	fmt.Println("after passing to function ", num) // num 数组不会因为调用而改变

}
