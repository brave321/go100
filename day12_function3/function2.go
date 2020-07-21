package main

import "fmt"

// 递归函数

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

//将函数作为参数

//函数可以作为其它函数的参数进行传递，然后在其它函数内调用执行，一般称之为回调

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

// 必包
// 什么是必包

func main() {
	// result := 0
	// for i := 0; i <= 10; i++ {
	// 	result = fibonacci(i)
	// 	fmt.Printf("fibonacci(%d) is:%d\n", i, result)
	// }

	// 输出

	// 	fibonacci(0) is:1
	// fibonacci(1) is:1
	// fibonacci(2) is:2
	// fibonacci(3) is:3
	// fibonacci(4) is:5
	// fibonacci(5) is:8
	// fibonacci(6) is:13
	// fibonacci(7) is:21
	// fibonacci(8) is:34
	// fibonacci(9) is:55
	// fibonacci(10) is:89

	//将函数当作参数传入

	// callback(1, Add)
	// 输出 The sum of 1 and 2 is: 3

}
