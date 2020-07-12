package main

import (
	"fmt"
)

// for 结构

// 如果想要执行 某些语句，go 语言只有for 结构可以使用

// for 1.go

func fortest1() {
	for i := 0; i < 5; i++ {
		fmt.Printf("this is the %d\n", i)
	}
}

// 两个for 循环嵌套
func fortest2() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			fmt.Println("---", i, "====", j)
		}
	}
}

// 使用for 无限循环

func fortest3() {
	for {
		fmt.Println("hello go")
	}
	// 死循环  hello go
}

// 使用 for-range 结构

func fortest4() {
	var nums = [...]int{1, 3, 11, 23, 33, 3}
	for _, v := range nums {
		fmt.Println(v)
	}

	// 输出
	// 	1
	// 3
	// 11
	// 23
	// 33
	// 3

}

//

func main() {
	// fmt.Println("hello world")
	// fortest1()
	// fortest2()
	// fortest3()
	// fortest4()
}
