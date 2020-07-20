package main

import "fmt"

// 函数参数与返回值

//在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的引用类型都是默认使用引用传递（即使没有显式的指出指针）

// 函数的返回值

func functiontest1(a, b, c int) int {
	return a * b * c
}

// 非命名返回值

func getX2AndX3(input int) (int, int) {
	return 2 * input, 3 * input
}

// 命名返回值
func getX2AndX3_2(input int) (x2 int, x3 int) {
	x2 = 2 * input
	x3 = 3 * input
	return
}

// 变长参数

func min(a ...int) int {
	if len(a) == 0 {
		return 0
	}
	min := a[0]
	for _, v := range a {
		if v < min {
			min = v
		}
	}
	return min
}

// 打印变成参数里的每一个参数
func funargs(p ...int) int {
	for k, v := range p {

		fmt.Println(k, v)

	}
	return 10
}

func main() {
	// fmt.Printf(" 函数的返回值为%d\n", functiontest1(12, 4, 6))
	// fmt.Println(getX2AndX3_2(12))
	// 输出 24 36
	// fmt.Println(getX2AndX3(11))
	// 输出 22 33

	// x := min(1, 33, 445, 23)
	// fmt.Printf("The minimum is: %d\n", x)
	// arr := []int{11, 34, 33, 212}
	// x = min(arr...)
	// fmt.Printf("The minimum in the array arr is: %d", x)
	// 输出
	//The minimum is: 1
	//The minimum in the array arr is: 11

	p := funargs(1, 33, 89)
	fmt.Println(p)

}
