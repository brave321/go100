package main

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
	// fmt.Println(p.firstName,p.lastName)
}

type Student struct {
	name  string
	age   int
	Class string
}

// 自定义构造函数

func Newstu(name1 string, age1 int, class1 string) *Student {
	return &Student{name: name1, age: age1, Class: class1}
}

func main() {
	// var pers1 Person
	// pers1.firstName = "chris"
	// pers1.lastName = "Woodward"
	// upPerson(&pers1)
	// fmt.Printf("the name of person is %s %s\n", pers1.firstName, pers1.lastName)
	// fmt.Println(pers1.firstName, pers1.lastName)

	// 输出

	// the name of person is CHRIS WOODWARD
	// CHRIS WOODWARD

	// var stu1 Student
	// var stu2 *Student = &Student{}
	// var stu3 *Student = new(Student)
	// fmt.Println(stu1, stu2, stu3)
	// 输出
	// { 0 } &{ 0 } &{ 0 }

	// 在struct 中无论使用的是指针的方式声明还是普通方式声明 访问成员都使用.
	// struct 分配内存使用new,返回的是指针
	// struct 没有构造函数，但是我们可以自己定义构造函数
	// struct 是我们自己定义的类型，不能和其他类型进行强制转换

	// var stu1 Student
	// fmt.Println(stu1)
	// // 输出 { 0 }

	// stu1.age = 2
	// stu1.name = "wd"
	// stu1.Class = "class1"
	// fmt.Println(stu1)
	// //{wd 2 class1}

	// var stu2 *Student = new(Student)
	// stu2.name = "jack1"
	// stu2.age = 33
	// fmt.Println(stu2.name, (*stu2).name)
	// // jack1 jack1

	// var stu3 *Student = &Student{name: "rose", age: 18, Class: "class3"}
	// fmt.Println(stu3.name, (*stu3).name)

	// rose rose

	stu2 := Newstu("cc", 22, "math")
	fmt.Println(stu2)
	// 输出
	// &{cc 22 math}

}
