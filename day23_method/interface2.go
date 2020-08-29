package main

import "fmt"

// 接口 duck typing

// go 没有 Go没有implements或extends关键字，这类编程语言叫作duck typing编程语言。

//Go类型系统采取了折中的办法，其做法如下。
//第一，结构体类型T不需要显式地声明它实现了接口I。只要类型T实现了接口I规定的所有方法，它就自动地实现了接口I
//。这样就像动态类型语言一样省了很多代码，少了许多限制。
//第二，将结构体类型的变量显式或者隐式地转换为接口I类型的变量i。
//这样就可以和其他静态类型语言一样，在编译时检查参数的合法性。

// 定义  duck typing

type ISsayHello interface {
	SayHello() string
}

type Duck struct {
	name string
}

type Person struct {
	name string
}

func (d Duck) SayHello() string {
	return d.name + "叫：ga ga !! ga!"
}

func (p Person) SayHello() string {
	return p.name + "说：helllo hello "

}

func main() {
	// 定义实现接口的对象
	duck := Duck{"Yaya"}
	person := Person{"steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println("============")
	// 输出
	//Yaya叫：ga ga !! ga!
	//	steven说：helllo hello
	//============

	// 定义接口类型的变量
	var i ISsayHello
	i = duck
	fmt.Println("%T,%v,%p\n", i, i, &i)
	fmt.Println(i.SayHello())
	// 输出
	//{Yaya} {Yaya} 0xc000010220
	//Yaya叫：ga ga !! ga!
	i = person
	fmt.Println("%T,%v,%p\n", i, i, &i)
	fmt.Println(i.SayHello())
	// 输出
	//{steven} {steven} 0xc000110200
	//steven说：helllo hello

	//一个函数如果接受接口类型作为参数，
	//那么实际上它可以传入该接口的任意一个实现类的对象作为参数。
	//定义一个接口变量，实际上可以赋值给任意一个实现了该接口的对象。
	//如果定义了一个接口类型的容器（数组或切片），实际上该容器可以存储任意一个实现类对象

}
