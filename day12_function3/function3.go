package main

import "fmt"

// 函数参数与返回值

func test1(a int, b int, c int) int {
	return a * b * c
}

// 命名返回值 x2 x3
func test2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

//非命名返回值

func test3(input int) (int, int) {
	return 2 * input, 4 + input
}

//传递变长参数
// 如果函数的最后一个参数是采用， .... type 的形式，那么这个函数可以处理一个变长的参数，长度可以为0 ，
// 这样的函数称为变参函数

//

func main() {
	fmt.Println("hello go")
	// fmt.Println("计算函数test1 三个返回值相乘", test1(3, 4, 5))
	// 输出
	// 计算函数test1 三个返回值相乘 60
	// 命名返回值
	// fmt.Println(test2(11))
	// 非命名返回值
	// fmt.Println(test3(11))
	// 输出 22 15

}
