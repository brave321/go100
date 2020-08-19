package main

import "fmt"

// 继承、多继承

//当结构体中的成员也是结构体时，该结构体就继承了这个结构体，
// 继承了其所有的方法与属性，当然有多个结构体成员也就是多继承。

type Person struct {
	Name string
	Age  int
}

type Teacher struct {
	Salary  int
	Classes string
}

type man struct {
	sex    string
	job    Teacher // 别名 继承Teacher
	Person         // 继承 Person
}

func main() {

	var man1 = new(man)
	man1.Age = 22
	man1.Name = "cc"
	man1.job.Salary = 3000
	fmt.Println(man1, man1.job.Salary)
	// 输出 &{ {3000 } {cc 22}} 3000
}
