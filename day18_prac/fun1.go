package main

import "fmt"

// map 的声明和初始化

// map 是引用类型

// var map1 map[keytype][valuetype]
// var map2 map[string]int

// 示例

func map1() {
	var mapLit map[string]int

	var mapAssigned map[string]int

	fmt.Println(mapLit, "====", mapAssigned)

	mapLit = map[string]int{"one": 1, "two": 2}

	mapCreated := make(map[string]float32)

	mapAssigned = mapLit

	fmt.Println("===", mapLit, "==", mapCreated, "==", mapAssigned)

	// map 元素添加 ,key 里面的是双引号

	mapCreated["key1"] = 43.6
	mapCreated["key2"] = 24.6
	mapCreated["key3"] = 14.6
	mapCreated["key4"] = 44.63

	//根据key 值获取元素值

	// fmt.Println(mapLit["one"])
	fmt.Printf("Map liter one is %d\n", mapLit["one"])
	fmt.Printf("Map mapCreated one is %f\n", mapCreated["key1"])
	// 输出 Map mapCreated one is 43.599998
	// 获取不存在的元素 为
	fmt.Printf("Map mapCreated one is %f\n", mapCreated["key10"])
	// 输出Map mapCreated one is 0.000000

	// 循环遍历 map 中的元素

	for k, v := range mapCreated {
		// fmt.Println(k, v)
		fmt.Printf("the key is: %s value is %f\n", k, v)
		// 输出
		// 		the key is: key1 value is 43.599998
		// the key is: key2 value is 24.600000
		// the key is: key3 value is 14.600000
		// the key is: key4 value is 44.630001

	}

}

// map 的数值可以是任意类型的 // map 的值是函数

func map2() {
	mf := map[int]func() int{
		1: func() int { return 10 },
		2: func() int { return 20 },
		5: func() int { return 50 },
	}
	fmt.Println(mf)
}

// map 容量

// // map 的值可以根据key，value 动态伸缩，也可以选择标明 初始容量
// 如 map2 := make(map[string]float32, 100)

// 用切片作为 map 的值

//mp1 := make(map[int][]int)

// val1 = map1[key1] 的方法获取 key1 对应的值 val1。如果 map 中不存在 key1，val1 就是一个值类型的空值。
// 为了解决这个问题，我们可以这么用：val1, isPresent = map1[key1]
//  如果你只是想判断某个 key 是否存在而不关心它对应的值到底是多少，你可以这么做：

// if _, ok := map1[key1]; ok {
//     // ...
// }

// map 删除元素  delete(map1, key1) 就可以s

func map3() {
	var mapLit map[string]int
	mapLit = map[string]int{"one": 1, "two": 2}

	for k, v := range mapLit {
		if _, ok := mapLit[k]; ok { // 判断元素如果有的话就输出

			fmt.Printf("the key is %s ,the value is %d\n", k, v)
		}

	}

	// 增加元素

	mapLit["three"] = 3
	mapLit["four"] = 4

	fmt.Println(mapLit)

	// 删除元素
	delete(mapLit, "four")

	fmt.Println(mapLit)

	// 输出结果
	// the key is one ,the value is 1
	// the key is two ,the value is 2
	// map[four:4 one:1 three:3 two:2]
	// map[one:1 three:3 two:2]

}

func map4() {
	var value int
	var isPresent bool
	map1 := make(map[string]int, 0)
	map1["New Delhi"] = 55
	map1["Beijing"] = 20
	map1["Washington"] = 25
	value, isPresent = map1["Beijing"]
	if isPresent {
		fmt.Printf("the value of \"Beijing\" in map1 is:%d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}
	value, isPresent = map1["Paris"]
	fmt.Printf("Is \"Paris\" in map1 ?:%t\n", isPresent)
	fmt.Printf("Value is: %d\n", value)

	delete(map1, "Washington")
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}

	// 输出

	// 	the value of "Beijing" in map1 is:20
	// Is "Paris" in map1 ?:false
	// Value is: 0
	// map1 does not contain Washington

}

//打印

func map5() {

	// var mapLit map[string]int
	// mapLit = map[string]int{"one": 1, "two": 2}
	// var value string
	var isPresent bool

	var map1 map[string]string

	map1 = map[string]string{"1": "Monday", "2": "Tuesday", "3": "Wednesday", "4": "Thursday", "5": "Friday", "6": "Saturday", "7": "Sunday"}
	fmt.Println(map1, isPresent)

	// value, isPresent = map1["Tuesday"]
	// if isPresent {
	// 	fmt.Printf("the value of \"Tuesday\" in map1 is: %d\n", value)
	// } else {
	// 	fmt.Println("map1 dose not container Tuesday ")
	// }

	//
	// fmt.Println("==========================")

	// value, isPresent = map1["Thursday"]
	// if isPresent {
	// 	fmt.Printf("the value of \"Thursday\" in map1 is: %d\n", value)
	// } else {
	// 	fmt.Println("map1 dose not container Thursday ")
	// }

	// value, isPresent = map1["Tuesday"]

	// value, isPresent = map1["Tuesday"]
	// if isPresent {
	// 	fmt.Printf("the value of \"Tuesday\" in map1 is:%d\n", value)
	// } else {
	// 	fmt.Printf("map1 does not contain Tuesday")
	// }
}

func main() {
	fmt.Println("hello world")
	// map1()
	// mapCreated()
	// map2()
	// map3()
	// map4()
	map5()
}
