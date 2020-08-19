package main

// 结构体定义和方法

import "fmt"

// 结构体的一般定义方法

// 使用 new
// 使用 new 函数给一个新的结构体变量分配内存，
// 它返回指向已分配内存的指针：var t *T = new(T)，
// 如果需要可以把这条语句放在不同的行（比如定义是包范围的，但是分配却没有必要在开始就做）。

type struct1 struct {
	i1  int
	f1  float32
	str string
}



type myStruct struct {
	i1 int 
	var v myStruct 
	var p *myStruct
}

func main() {
	// fmt.Println("hello world")
	// ms := new(struct1)
	// fmt.Println(ms)
	// ms.i1 = 10
	// ms.f1 = 15.4
	// ms.str = "Chris"
	// fmt.Printf("The int is:%d\n", ms.i1)
	// fmt.Printf("The float is:%f\n", ms.f1)
	// fmt.Printf("The string is:%s\n", ms.str)
	// fmt.Println(ms)

	// 输出
	// &{0 0 }
	// The int is:10
	// The float is:15.400000
	// The string is:Chris
	// &{10 15.4 Chris}

 // struct 例子

 








	

}
