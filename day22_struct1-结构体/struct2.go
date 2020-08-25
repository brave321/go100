package main

import (
	"fmt"
)

// 语法糖作用提升程序的可读性，从而减少程序代码出错的机会
// 结构体和数组中语法糖
// 使用内置函数 new()对结构体进行 实列化，从而减少程序代码出错的机会

// 定义结构体
type Emp struct {
	name string
	age  int8
	sex  byte
}

// 其他数据结构的语法糖
func Syntacticsugar() {
	// 数组中的语法糖
	arr := [4]int{10, 20, 4, 92}
	fmt.Println(arr)
	// [10 20 4 92]

	//
	arr2 := &arr
	// 获取aarr2 中的某个元素
	fmt.Println((*arr2)[len(arr)-2])
	// 输出 4

	// 切片中的愈发糖
	arr3 := []int{100, 34, 323, 234}
	arr4 := &arr3
	fmt.Println((*arr4)[len(arr3)-2])
	// 输出 323

}

func main() {
	fmt.Println("结构体 语法糖")
	// 使用new() 内置函数实例 化struct
	emp1 := new(Emp)
	fmt.Println(emp1)
	// 输出
	//&{ 0 0}

	fmt.Printf("emp1 %T ,%v,%p \n", emp1, emp1, emp1)
	//
	(*emp1).name = "cc"
	(*emp1).age = 30
	(*emp1).sex = 1
	fmt.Printf("emp1 %T ,%v,%p \n", emp1, emp1, emp1)
	//emp1 *main.Emp ,&{cc 30 1},0xc000110000
	// 语法糖写法
	emp1.name = "make"
	emp1.age = 33
	emp1.sex = 1
	fmt.Println(emp1)
	// 输出
	//&{make 33 1}
	fmt.Println("=============")
	Syntacticsugar()

}
