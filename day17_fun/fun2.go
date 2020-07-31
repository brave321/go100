package main

import (
	"fmt"
)

// 切片 是对数组一个连续片段的引用
// 切片是可以索引的
// 切片是一个长度可变的数组
// 切片提供了计算容量的函数cap()
// 优点 切片是引用类型
// 声明切片的格式

// var identifier []type

// 一个切片在未初始化之前默认为 nil ,长度为0

// 切片的初始化格式是：var slice1 []type=arr1[start:end]

// 示例

func slicetest1() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5]

	// 数组赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	//

	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Println(arr1, "=====")
	fmt.Println(slice1)
	// 输出
	//     Slice at 0 is 2
	// Slice at 1 is 3
	// Slice at 2 is 4
	// [0 1 2 3 4 5] =====
	// [2 3 4]

}

// 切片传给函数
// identify []int
func slicetest2(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// 使用make 创建一个切片

// 格式  var slice1 [] type = make([] type,len) //len 代表长度
// var slicence2 []type = make([]type ,len,cap)

func slicetest3() {
	var s1 []int = make([]int, 3)
	for i := 0; i < len(s1); i++ {
		s1[i] = 5 * i
	}

	for i := 0; i < len(s1); i++ {
		fmt.Printf("slice at %d is %d\n", i, s1[i])
	}

	fmt.Println(s1, len(s1), cap(s1))
}

// new() 和make() 的区别

// new(T) 为每个新的类型T 分配一片内存，初始化为0 并且返回类型为 *T 的内存地址 ，返回一个指向类型为T

//make(T) 返回一个类型为T 的初始值

// new 函数分配内存  make 函数初始化

func slicetest4() {
	var p1 *[]int = new([]int)
	p2 := new([]int)

	fmt.Println(p1) // 输出 &[]
	fmt.Println(p2) // 输出 &[]

	// fmt.

	v := make([]int, 10, 50)
	fmt.Println(v)

}

// bytes 包

func main() {
	fmt.Println("hello world")
	// slicetest1()
	// var arr = [5]int{1, 3, 4, 6, 5}
	// fmt.Println(slicetest2(arr[:]))
	// 输出 19
	// slicetest3()
	// slicetest4()

}
