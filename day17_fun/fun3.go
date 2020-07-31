package main

import "fmt"

// 切片的循环遍历

func s1test() {
	var slice1 []int = make([]int, 4)
	slice1[0] = 1
	slice1[1] = 2
	slice1[2] = 3
	slice1[3] = 4

	for ix, value := range slice1 {
		fmt.Printf("slice at %d is:%d\n", ix, value)
	}

	//输出
	// 	slice at 0 is:1
	// slice at 1 is:2
	// slice at 2 is:3
	// slice at 3 is:4

}

//

func s2test() {
	seasons := []string{"Spring", "Summer", "Autumn", "Winter"}
	for ix, value := range seasons {
		fmt.Printf("slice at %d is:%s\n", ix, value)
		// 输出
		// 		slice at 0 is:Spring
		// slice at 1 is:Summer
		// slice at 2 is:Autumn
		// slice at 3 is:Winter

	}
}

//写一个 Sum 函数，传入参数为一个 32 位 float 数组成的数组 arrF，返回该数组的所有数字和。
func s3test(arr [3]float32) {
	var sum float32
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
	}
	fmt.Println("数组的和为", sum)

}

//

func s4test(slice1 []int) {
	var sum int
	for i := 0; i < len(slice1); i++ {
		sum += slice1[i]
	}
	fmt.Println("切片的和为", sum)

}

// 切片与重组
// slice1 := make([]type, start_length, capacity)
// start_length 作为切片初始长度而 capacity 作为相关数组的长度。

//
func s5test() {
	slice1 := make([]int, 0, 10)
	for i := 0; i < cap(slice1); i++ {
		slice1 = slice1[0 : i+1]
		slice1[i] = i
		fmt.Printf("The length of slice is %d\n", len(slice1))
	}

	// 输出
	// The length of slice is 1
	// The length of slice is 2
	// The length of slice is 3
	// The length of slice is 4
	// The length of slice is 5
	// The length of slice is 6
	// The length of slice is 7
	// The length of slice is 8
	// The length of slice is 9
	// The length of slice is 10

}

func main() {
	fmt.Println("hello world")
	// s1test()
	// s2test()
	// arr1 := [3]float32{11.3, 23.4, 5.8}
	// s3test(arr1)
	// 输出 数组的和为 40.5
	// var s []int = []int{1, 2, 3}
	// s4test(s)
	s5test()
}
