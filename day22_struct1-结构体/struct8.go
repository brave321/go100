package main

import (
	"fmt"
)

// 结构体嵌套模拟继承关系

type Person struct {
	name string
	age  int
	sex  string
}

type Student struct {
	Person
	schoolName string
}

func printInfo(s1 Student) {
	fmt.Println(s1)
	fmt.Printf("%+v \n", s1)
	fmt.Printf(" 姓名：%s, 年龄: %d, 性别：%s, 学校：%s\n", s1.name, s1.age, s1.sex, s1.schoolName)
	fmt.Println("============")
}

func main() {
	// 1 、实例化 Persoon
	p1 := Person{"steven", 35, "男"}
	fmt.Println(p1)
	fmt.Println("=========")

	// 2 实例化 并初始化struct
	// 写法1
	s1 := Student{p1, "北航软件学校"}
	fmt.Println(s1)
	// 输出
	// {{steven 35 男} 北航软件学校}

	// 写法2

	s2 := Student{Person{"josh", 30, "男"}, "北航软件学校"}
	fmt.Println(s2)
	// 输出 {{josh 30 男} 北航软件学校}

	// 写法3

	s3 := Student{Person{"joh2", 19, "女"}, "北大元培学院"}
	fmt.Println(s3)
	// 输出 {{joh2 19 女} 北大元培学院}

	// 写法4
	s4 := Student{}
	s4.name = "Danile"
	s4.sex = "男"
	s4.age = 24
	s4.schoolName = "北京师范大学"
	fmt.Println(s4)
	//{{Danile 12 男} 北京师范大学}
	printInfo(s4)
	// 输出
	// 姓名：Danile, 年龄: 24, 性别：男, 学校：北京师范大学

}
