package main

import "fmt"

// struct 匿名成员(字段 属性)

// 结构体中，每个成员不一定都有名称，也允许字段没有名字，即匿名成员。

// 匿名成员的一个重要作用，可以用来实现oop中的继承。

// 同一种类型匿名成员只允许最多存在一个。

// 当匿名成员是结构体时，且两个结构体中都存在相同字段时，优先选择最近的字段。

type Person struct {
	Name string
	Age  int
	// Person
}

type Student struct {
	score string
	Age   int
	Person
}

type innerS struct {
	in1 int
	in2 int
}

type outerS struct {
	b int
	c float32
	int
	innerS
}

func main() {
	// var stu = new(Student)
	// stu.Age = 22
	// fmt.Println(stu.Person, stu.Age)
	// 输出

	// { 0} 22

	outer := new(outerS)
	// fmt.Println(outer)
	outer.b = 6
	outer.c = 8.4
	outer.int = 11
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)
	// 输出
	// 	outer.b is: 6
	// outer.c is: 8.400000
	// outer.int is: 11
	// outer.in1 is: 5
	// outer.in2 is: 10

	// 使用结构体字面量

	outer2 := outerS{5, 5.4, 3, innerS{9, 3}}
	fmt.Println("outer2 is", outer2)

	// 输出

}
