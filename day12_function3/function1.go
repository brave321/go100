package main

import "fmt"

// 递归函数
// 斐波那契数列 ,从第三个数开始，每个数均为前2个数之和
// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, 89, 144, 233, 377, 610, 987, 1597, 2584, 4181, 6765, 10946, …

// 描述3个数 n-1,n-2 ,n
//2  1   3
// 第n 个数 n-2,n-1 ,n

// n=2
// res=f(2)+(f0)
// 	1  + 0

func fibonacci(n int) (res int) {
	if n <= 1 {
		res = 1
	} else {
		res = fibonacci(n-1) + fibonacci(n-2)
	}
	return
}

// 层层嵌套的递归函数

func even(nr int) bool {
	if nr == 0 {
		return true
	}
	return odd(RevSign(nr) - 1)
}

func odd(nr int) bool {
	if nr == 0 {
		return false
	}
	return even(RevSign(nr) - 1)
}

func RevSign(nr int) int {
	if nr < 0 {
		return -nr
	}
	return nr
}

// 使用递归函数从10打印到1
// 递归自己调自己 ，有明确的退出条件

// 从10 打印到1 普通的函数

func print_test() {
	for i := 10; i > 0; i-- {
		fmt.Println(i)
	}
}

// 从10 打印到1的递归函数 ???

// 将函数作为参数

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func callback(y int, f func(int, int)) {
	f(y, 2) // this becomes Add(1, 2)
}

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}

}

func main() {
	// fmt.Println("递归函数")

	// fmt.Printf("%d is even: is %t\n", 16, even(16)) // 16 is even: is true
	// fmt.Printf("%d is odd: is %t\n", 17, odd(17))

	// fmt.Printf("%d is odd: is %t\n", 18, odd(18))

	// 输出

	// 	递归函数
	// 16 is even: is true
	// 17 is odd: is true
	// 18 is odd: is false

	// result := 0
	// for i := 0; i <= 10; i++ {
	// 	result = fibonacci(i)
	// 	fmt.Printf("fibonacci(%d) is: %d\n", i, result)
	// }

	// print_test()
	// callback(1, Add)
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Printf("0+1+...%d=%d\n", i, a(i))
	}

}
