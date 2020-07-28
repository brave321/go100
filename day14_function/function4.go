package main

import "fmt"

// go 将函数作为参数

func callback(y int, f func(int, int)) {
	f(y, 2)
}

func Add(a, b int) {
	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func main() {
	fmt.Println("将函数作为参数")
	callback(1, Add)

}
