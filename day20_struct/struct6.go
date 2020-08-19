package main

import "fmt"

// 结构体中的方法

// go 语言中的方法是作用在特定类型的变量上，因此自定义的类型都可以有方法，不仅仅是结构体中
// go 中的方法和传统的类的方法不太一样，方法和类并非组织在一起，传统的oop 方法和类放在一个文件里面，而go 语言只要在同一个包里就可以

type Person struct {
	Name string
	Age  int
}

func (p Person) Getname() string {
	fmt.Println(p.Name)
	return p.Name
}

func main() {
	fmt.Println("结构体中的方法")
	var person1 = new(Person)
	person1.Age = 22
	person1.Name = "wd"
	person1.Getname()
}
