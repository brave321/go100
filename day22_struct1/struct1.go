package main

// 结构体的使用类型

// 单一的结构体，已经满足不了现实开发需求，于是需要定义复杂的数据类型
// 结构体的定义

// 代码示例

//type 类型名 struct {
//	成员属性1 类型1
//	成员属性2 类型2
//	成员属性3, 成员属性4
//}
import "fmt"

// 定义Teacher 结构体
type Teacher struct {
	name string
	age  int8
	sex  byte
}

func main() {
	// var 声明方式实例化结构体，初始化方式为：对象.属性=值
	var t1 Teacher
	fmt.Println(t1)
	fmt.Printf("t1:%T,%v,%q \n", t1, t1, t1)
	t1.name = "cc"
	t1.age = 30
	t1.sex = 1
	fmt.Println(t1)
	fmt.Println("========")
	// 2 短变量声明实列化结构体
	t2 := Teacher{}
	t2.name = "kk"
	t2.age = 19
	t2.sex = 1
	fmt.Println(t2)
	fmt.Println("==========")

	// 3 ，变量简短声明氏 实例化结构体
	t3 := Teacher{
		name: "john",
		age:  28,
		sex:  1,
	}
	fmt.Println(t3, "t3 before")
	fmt.Println("==========")
	t3 = Teacher{name: "john2", age: 25, sex: 1}
	fmt.Println("==========")
	fmt.Println(t3, "t3 after")

	// 4 变量简短式声明实力化结构体

}
