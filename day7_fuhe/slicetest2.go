package main

import "fmt"

// 切片的长度,使用make 创建切片
//func make([]T,len,cap) []T 通过传递类型，长度和容量来创建切片
// 容量是可选参数，默认值为切片长度。 make 函数创建一个数组，并返回引用该数组的切片

//

func slicetest2() {
	i := make([]int, 5, 5)
	fmt.Println(i)
	// 输出 [0 0 0 0 0]
	i[0] = 11
	fmt.Println(i)
	// 输出 [11 0 0 0 0]

}

// 切片追加元素
func slicetest3() {
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "old length", len(cars), "old capacity", cap(cars))
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "new1 length", len(cars), "new1 capacity", cap(cars))

	// cars = append(cars, "BYD")
	// fmt.Println("cars:", cars, "new2 length", len(cars), "new2 capacity", cap(cars))
	// cars = append(cars, "QQ")
	// fmt.Println("cars:", cars, "new3 length", len(cars), "new3 capacity", cap(cars))
	// cars = append(cars, "benz")
	// fmt.Println("cars:", cars, "new4 length", len(cars), "new4 capacity", cap(cars))

	// 输出

	// cars: [Ferrari Honda Ford] old length 3 old capacity 3
	// cars: [Ferrari Honda Ford Toyota] new1 length 4 new1 capacity 6
	// cars: [Ferrari Honda Ford Toyota BYD] new2 length 5 new2 capacity 6
	// cars: [Ferrari Honda Ford Toyota BYD QQ] new3 length 6 new3 capacity 6
	// cars: [Ferrari Honda Ford Toyota BYD QQ benz] new4 length 7 new4 capacity 12
	// 初始的时候 长度为3 ，容量为3
	// 增加了一个元素，长度变为了3 ，容量变为了6 ，容量了翻了一倍
	// 当长度和容量都为6的时候，在添加了一个元素，容量变为了12，容量翻了一倍

	// 切片类型的零值为nil ，一个nil 切片的长度和容量为0，可以使用 append 函数将值追加到 nil 切片
	var names []string
	// fmt.Println(names)
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}

	// 输出
	// 	slice is nil going to append
	// names contents: [John Sebastian Vinay]

	// 将两个切片合并到一起
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)

	// 输出
	//food: [potatoes tomatoes brinjal oranges apples]

}

// 切片的函数传递

func slicetest4(numbers []int) {
	for i := range numbers {
		numbers[i] -= 2
	}
}

// 切片的内存优化
// 切片持有对底层数组的引用，只要切片在内存中，数组不能被垃圾回收。
// 在内存管理方面，这是需要注意的。
//让我们假设我们有一个非常大的数组，我们只想处理它的一小部分，然后，我们由这个数组创建一个切片，
//并开始处理切片。这里需要重点注意的是，在切片引用时数组仍然存在内存中
//一种解决方法是使用 copy 函数 func copy(dst，src[]T)int 来生成一个切片的副本。这样我们可以使用新的切片，原始数组可以被垃圾回收。

func slicetest5() []string {
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
	// fmt.Println(countries, "new countries", neededCountries)
	countriesCpy := make([]string, len(neededCountries))
	fmt.Println("countriesCpy", countriesCpy, "neededCountries", neededCountries)
	copy(countriesCpy, neededCountries) // copy 生成一个新的副本，现在 countries 数组可以被垃圾回收, 因为 neededCountries 不再被引用
	fmt.Println("new countriesCpy is", countriesCpy)
	return countriesCpy
	// 输出
	// 	countriesCpy [  ] neededCountries [USA Singapore Germany]
	// new countriesCpy is [USA Singapore Germany]
	// [USA Singapore Germany]

}

func main() {
	// slicetest2()
	// slicetest3()
	// nos := []int{11, 45, 12}
	// fmt.Println("slice before function call", nos)
	// slicetest4(nos)
	// fmt.Println("slice after function call", nos)
	//
	// 	slice before function call [11 45 12]
	// slice after function call [9 43 10]
	// slicetest5()
	countriesNeeded := slicetest5()
	fmt.Println(countriesNeeded)

}
