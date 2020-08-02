package main

import "fmt"

// 快速回顾

// 数组练习

// 数组的定义

func arr1test() {
	var arr1 [5]int
	fmt.Println(arr1)

	for i := 0; i < len(arr1); i++ {
		// fmt.Println(i)
		arr1[i] = i * 2 // 数组赋值
		fmt.Println("==", arr1[i], "==")
	}

	// 获取 数组的索引
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("Array at index is %d value is %d\n", i, arr1[i])
	}

}

// 切片

// 切片是对数组一个连续片段的引用 ，切片是可以索引的

func slicetest1() {
	var arr1 [5]int              // 数组
	var slice1 []int = arr1[1:4] //从数组的第二个取值
	fmt.Println(arr1, slice1)    // 打印数组，与切片
	// 输出
	// [0 0 0 0 0] [0 0 0]

	// 数组赋值
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
		// fmt.Println(arr1)
		//输出 [0 1 2 3 4] // 从0-5
	}

	// 打印切片
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
		// 输出
		// Slice at 0 is 1
		// Slice at 1 is 2
		// Slice at 2 is 3
	}

}

//// 将切片传递给函数

func sum(a []int) int {
	s := 0
	for i := 0; i < len(a); i++ {
		s += a[i]
	}
	return s
}

// 使用make 创建切片
// 使用make 函数创建一个切片

func main() {
	fmt.Println("切片与数组的练习")
	// arr1test()
	// slicetest1()
	var a1 = [5]int{1, 4, 5, 6, 7}
	// fmt.Println(a1)
	fmt.Println(sum(a1[:]))
	// 输出 23
}
