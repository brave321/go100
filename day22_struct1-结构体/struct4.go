package main

import "fmt"

// 结构体深浅拷贝

type Dog struct {
	name  string
	color string
	age   int8
	kind  string
}

func main() {
	// 实现结构体的深拷贝
	// struct 是值类型，默认的复制就是深拷贝
	d1 := Dog{"tom", "black", 2, "中华田园犬"}
	fmt.Printf("d1: %T ,%v,%p \n", d1, d1, d1)
	// 输出
	//d1: main.Dog ,{tom black 2 中华田园犬},%!p(main.Dog={tom black 2 中华田园犬})

	fmt.Println("/////")
	d2 := d1 // 深拷贝
	fmt.Printf("d2: %T ,%v,%p \n", d2, d2, d2)
	d2.name = "毛毛"
	fmt.Println("d2 修改后d2:", d2)
	fmt.Println("d1:", d1)
	// 输出
	//	d2 修改后d2: {毛毛 black 2 中华田园犬}
	//d1: {tom black 2 中华田园犬}

	// 实现结构体浅拷贝，直接赋值指针地址 &
	d3 := &d1
	fmt.Printf("d3: %T ,%v,%p \n", d3, d3, d3)
	d3.name = "求求"
	d3.color = "蓝色"
	d3.kind = "萨摩耶"
	fmt.Println("d3 修改后d3:", d3)
	fmt.Println("d1", d1)
	// 输出
	//d3 修改后d3: &{求求 蓝色 2 萨摩耶}
	//d1 {求求 蓝色 2 萨摩耶}

	// d1因为d3的改变而改变，所以是浅拷贝

	// 3 实现结构体浅拷贝，通过new() 函数来实现实例化 对象
	d4 := new(Dog)
	fmt.Println(d4)
	d4.name = "多多"
	d4.color = "棕色"
	d4.age = 1
	d4.kind = "巴哥犬"
	d5 := d4
	fmt.Printf("d4: %T ,%v,%p \n", d4, d4, d4)
	fmt.Printf("d5: %T ,%v,%p \n", d5, d5, d5)
	d5.color = "金色"
	d5.kind = "金毛"
	fmt.Println("d5 修改后d5:", d5)
	fmt.Println("d4", d4)
	// 输出
	//
	//d5 修改后d5: &{多多 金色 1 金毛}
	//d4 &{多多 金色 1 金毛}

}
