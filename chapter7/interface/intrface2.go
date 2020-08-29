package main

import "fmt"

// 鸭子类型

type ISayHello interface {
	SayHello() string
}

func (d Duck) SayHello() string {
	return d.name + "叫: ga ga ga ！！！"
}

func (p Person) SayHello() string {
	return p.name + "说: hello go"
}

type Duck struct {
	name string
}

type Person struct {
	name string
}

func main() {
	fmt.Println("hello world")
	// 定义实现接口的对象
	duck := Duck{"Yaya"}
	person := Person{"Steven"}
	fmt.Println(duck.SayHello())
	fmt.Println(person.SayHello())
	fmt.Println("================")

	// 定义接口类型的变量
	var i ISayHello
	var j ISayHello
	i = duck
	fmt.Printf("%T,%v,%p \n", i, i, &i)
	fmt.Println(i.SayHello())
	fmt.Println("=======")
	j = person
	fmt.Printf("%T,%v,%p \n", j, j, &j)
	fmt.Println(j.SayHello())

}
