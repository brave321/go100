package main

import "fmt"

// 切片
// 切片是由数据建立的一种方便、灵活且功能强大的包装，切片本身不拥有任何数据
//只是对数组的 引用
// 一个数组的表示形式为 [n]T ，切片的表示形式为 []T 表示

// 切片的创建
func slicetest1() {
	a := [5]int{11, 121, 34, 78, 36}
	var b []int = a[1:4]
	fmt.Println(b)
}

//切片的创建方法2

func slicetest2() {
	c := []int{1, 2, 3, 4, 5}
	fmt.Println(c, len(c))
	for i := 0; i < len(c); i++ {
		fmt.Println(i)
	}

	fmt.Println("============")
	for i, v := range c {
		fmt.Printf("the key is %d value is %d\n", i, v)
	}
}

// 切片的修改 修改了切片的长度，原始的数组也变了

func slicetest3() {
	darr := [...]int{1, 3, 4, 9, 811, 98}
	dslice := darr[2:5]
	fmt.Println("array bdfore", darr, dslice)
	for i := range dslice {
		// fmt.Println(dslice[i])
		dslice[i]++
	}
	fmt.Println(dslice)
	fmt.Println(darr)
}

// 切片的长度和容量
func slicetest4() {
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))
	// 输出 length of slice 2 capacity 6

}

func main() {
	fmt.Println("hello")
	// slicetest1()
	// slicetest2()
	// slicetest3()
	slicetest4()
}
