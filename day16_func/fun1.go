package main

import "fmt"

// 数组与切片

// 数组是具有相同 唯一类型的一组已编号且长度固定的数据序列

// 数组的定义

func arrtest1() {
	var arr1 [5]int
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i * 2
	}

	// for i := 0; i < len(arr1); i++ {
	// 	fmt.Printf("Array at index %d is %d\n", i, arr1[i])
	// }
	// 也可以使用for range

	for k, _ := range arr1 {
		fmt.Printf("Array at index %d is %d\n", k, arr1[k])
	}

}

// 数组是一种值类型，可以通过new 来创建

func arrtest2() {
	a := [...]string{"a", "b", "c", "d"}
	for i := range a {
		fmt.Println("Array item", i, "is", a[i])
	}

	var arr2 = new([4]int)
	fmt.Println(arr2)
	// fmt.Println(*arr2)
	// 输出 &[0 0 0 0]

	var arr3 [4]int
	fmt.Println(arr3)

	// 输出 [0 0 0 0]
	// var arr1 = new([5]int)
	// arr2 := *arr1
	// arr2[2] = 100
	// fmt.Println(arr2)
}

func f(a [3]int) { fmt.Println(a) }

func fp(a *[3]int) {
	fmt.Println(a)
}

func arrtest4() {
	var arr1 = new([5]int) // 通过new 创建的数组
	// var arr2 [5]int
	// fmt.Println("arr1", arr1, "arr2", arr2)

	//这样的结果就是当把一个数组赋值给另一个时，需要在做一次数组内存的拷贝操作。例如：

	arr2 := *arr1
	arr2[2] = 100
	fmt.Println(arr1, arr2)
}

//练习7.2：for_array.go: 写一个循环并用下标给数组赋值（从 0 到 15）并且将数组打印在屏幕上。

func arrtest5() {
	var a [16]int
	for i := 0; i < len(a); i++ {
		a[i] = i
	}
	for k, v := range a {
		fmt.Println(k, v)
	}

}

// 数组常量

func arraytest6() {
	var arrKeyValue = [5]string{3: "Chris", 4: "Ron"}
	for i := 0; i < len(arrKeyValue); i++ {
		fmt.Printf("Person at %d is %s\n", i, arrKeyValue[i])
	}
}

func arraytest7(a *[3]int) {
	fmt.Println(a)
}

func main() {
	fmt.Println("数组与切片")
	// arrtest1()
	// 输出
	// 	Array at index 0 is 0
	// Array at index 1 is 2
	// Array at index 2 is 4
	// Array at index 3 is 6
	// Array at index 4 is 8
	// arrtest2()
	// var ar [3]int
	// f(ar)
	// fp(&ar)
	// arrtest4()
	// 输出
	//arr1 &[0 0 0 0 0] arr2 [0 0 0 0 0]
	// arrtest5()
	// arraytest6()
	// 你可以取任意数组常量的地址来作为指向新实例的指针

	// for i := 0; i < 3; i++ {
	// 	arraytest7(&[3]int{i, i * i, i * i * i})
	// }

	// 输出
	// &[0 0 0]
	// &[1 1 1]
	// &[2 4 8]

}
