package main

import "fmt"

// 一 、方法的概念

//Go语言同时有函数和方法，方法的本质是函数，但是方法和函数又有所不同。

//1 含义不同
//
//函数（function）是一段具有独立功能的代码，
//可以被反复多次调用，从而实现代码复用。
//而方法（method）是一个类的行为功能，只有该类的对象才能调用。
//
//2 方法有接受者，而函数无接受者
//Go语言的方法（method）是一种作用于特定类型变量的函数。
//这种特定类型变量叫作接受者（receiver）

//3 函数不可以重名，而方法可以重名
//
//只要接受者不同，方法名就可以相同。

// 二、 方法的定义

//func (接受者变量 接受者类型) 方法名（参数列表）（返回值类型） {
//	// 方法体
//	接受者在func关键字和方法名之间编写，
//	接受者可以是struct类型或非struct类型，
//	可以是指针类型或非指针类型
//	。接受者中的变量在命名时，官方建议使用接受者类型的第一个小写字母
//}

// 示例 方法与函数
type Employee struct {
	name, currency string
	salary         float64
}

// printSalary() 的方法
func (e Employee) printSalary() {
	fmt.Printf("员工姓名 %s, 薪资 %s%.2f \n", e.name, e.currency, e.salary)
}

// printSalary 函数
func printSalary(e Employee) {
	fmt.Printf("员工姓名 %s, 薪资 %s%.2f \n", e.name, e.currency, e.salary)

}

func main() {
	emp1 := Employee{"cc", "$", 2000}
	emp1.printSalary() // 方法
	printSalary(emp1)  // 函数

	// 输出
	//员工姓名 cc, 薪资 $2000.00
	//员工姓名 cc, 薪资 $2000.00

}
