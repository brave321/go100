package main

import "fmt"

// 结构体作为函数的参数及返回值有两种形式：
// 值传递和引用传递

type Flower struct {
	name, color string
}

// 传结构体对象

func changeInfo1(f Flower) {
	f.name = "月季花"
	f.color = "粉色"
	fmt.Printf("函数 changeInfo1 内f1: %T,%v,%p \n", f, f, &f)
}

func changeInfo2(f *Flower) {
	f.name = "蔷薇"
	f.color = "紫色"
	fmt.Printf("函数 changeInfo2 内f1: %T,%v,%p %p \n", f, f, f, &f)
}

// 返回结构体对象 结构体前面不加*
func getFlower1() (f Flower) {
	f = Flower{"牡丹", "白色"}
	fmt.Printf("函数 getFlower1 内f1: %T,%v,%p \n", f, f, &f)
	return
}

// 返回结构体指针

func getFlower2() (f *Flower) {
	temp := Flower{"芙蓉", "红"}
	fmt.Printf("函数 getFlower2 内f: %T,%v,%p \n", f, f, &f)
	f = &temp
	fmt.Printf("函数 getFlower2 内 f: %T,%v,%p,%p\n", f, f, f, &f)
	return
}

func main() {
	fmt.Println("hello")
	// 1 结构体 作为参数的用法
	f1 := Flower{"玫瑰", "红色"}
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1)
	// 输出
	//f1:main.Flower,{玫瑰 红色},0xc00000c060
	// 2 结构体   对象作为参数
	changeInfo1(f1)
	fmt.Printf("f1:%T,%v,%p\n", f1, f1, &f1)
	fmt.Println("=======")
	// 输出
	//	函数 changeInfo1 内f1: main.Flower,{月季花 粉色},0xc000118060
	//f1:main.Flower,{玫瑰 红色},0xc000118000
	//	=======

	// 3 、将结构体指针作为参数
	changeInfo2(&f1)
	fmt.Printf("f1:%T,%v,%p \n", f1, f1, &f1)
	fmt.Println("=======")

	// 输出
	//	函数 changeInfo2 内f1: *main.Flower,&{蔷薇 紫色},0xc0000a6020 0xc0000ae020
	//f1:main.Flower,{蔷薇 紫色},0xc0000a6020
	//	=======

	// 4 结构体作为返回值的用法
	// 结构体对象作为返回值
	f2 := getFlower1()
	f3 := getFlower1()
	fmt.Println("更改前", f2, f3)
	fmt.Printf("f2 更改前地址为 %p\n ,f3 地址为%p\n", &f2, &f3)
	f2.name = "杏花"
	fmt.Println("更改后", f2, f3)
	fmt.Printf("f2 更改后地址为 %p\n ,f3 地址为%p\n", &f2, &f3)
	// 输出
	//f2 更改前地址为 0xc00000c1c0
	//,f3 地址为0xc00000c240
	//更改后 {杏花 白色} {牡丹 白色}
	//f2 更改后地址为 0xc00000c1c0
	//,f3 地址为0xc00000c240

	// 返回结构体指针
	f4 := getFlower2()
	f5 := getFlower2()
	fmt.Println("更改前", f4, f5)
	fmt.Printf("f4 更改前址为 %p\n ,f5 更改前地址为%p\n", &f4, &f5)
	f4.name = "桃花"
	fmt.Println("更改后", f4, f5)
	fmt.Printf("f4 更改后的址为 %p\n ,f5 更改后的地址为%p\n", &f4, &f5)

}
