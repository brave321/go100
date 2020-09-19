package main

import "fmt"

// 鸭子类型

//duck typing

type ISayHello interface {
	SayHello() string
}

type Duck struct {
	name string
}

type Person struct {
	name string
}

func (d Duck) SayHello() string {
	return d.name + "叫：ga ga !!"
}

func (p Person) SayHello() string {
	return p.name + "say: hello go "
}

func main() {
	// 定义接口的对象 不同的类型打印不同的
	duck := Duck{"鸭子"}
	person := Person{"Steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	// 输出
	//鸭子叫：ga ga !!
	//Stevensay: hello go
	// 定义接口类型的变量
	var i ISayHello
	i = duck
	fmt.Printf("%T,%v,%p \n", i, i, &i)
	fmt.Println(i.SayHello())
	// 输出
	//
	//main.Duck,{鸭子},0xc000010220
	//鸭子叫：ga ga !!
	i = person
	fmt.Printf("%T,%v,%p \n", i, i, &i)
	fmt.Println(i.SayHello())
	// 输出
	// main.Person,{Steven},0xc000010220
	// Stevensay: hello go

}
