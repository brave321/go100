package main

import "fmt"

// 必包

// 闭包是由函数体及其引用环境组合而成的实体,

func clourse() func(int) int {
	var x int
	return func(a int) int {
		x++
		return a + x
	}
}

func main() {
	fmt.Println("go 必包")
	A := clourse()
	fmt.Println("x:", A(1))
	fmt.Println("x:", A(2))
}
