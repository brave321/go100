package main

import "fmt"

// slice make 内建函数
// slice 删除元素 中间的值 首尾的值

func slictest2(arr []int) {
	fmt.Println(len(arr))
}

func slictest() {
	arr := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("arr[2:6]", arr[2:6])
}

func main() {
	slictest()
}
