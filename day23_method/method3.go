package main

import "fmt"

// 定义指针，指针作为接受者

type Rectangle struct {
	width, height float64
}

func (r Rectangle) setValue() {
	fmt.Printf("setValue 方法中r 的地址：%p \n", &r)
	r.height = 10
}

func (r *Rectangle) setValue2() {
	fmt.Printf("setValue2 方法中r 的地址：%p \n", &r)
	r.height = 20
}

func main() {
	r1 := Rectangle{5, 8}
	r2 := r1
	r1.setValue()
	// 打印对象中的内存地址
	fmt.Printf("r1 的内存地址%p \n", &r1)
	fmt.Printf("r2 的内存地址%p \n", &r2)

	fmt.Println("r1.height=", r1.height)
	fmt.Println("r2.height=", r2.height)
	fmt.Println("===================")
	r1.setValue2()

	fmt.Println("r1.height=", r1.height)
	fmt.Println("r2.height=", r2.height)
	fmt.Println("===================")

	// 输出

	//setValue 方法中r 的地址：0xc000016080
	//r1 的内存地址0xc000016060
	//r2 的内存地址0xc000016070
	//r1.height= 8
	//r2.height= 8
	//===================
	//setValue2 方法中r 的地址：0xc00000e030
	//r1.height= 20
	//r2.height= 8
	//===================

}
