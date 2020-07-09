package main

// go 常量，术语常量用于表示固定的值
// 关键字 const 用于表示常量，常量不允许重新赋值

import (
	"fmt"
	"math"
)

func constest1() {
	// int 类型的常量
	const a = 55
	// a = 89 // 不允许重新赋值
	fmt.Println("the value of a is ", a)
	// 输出
	// the value of a is  55
	fmt.Println("Hello go")
	var a1 = math.Sqrt(4)
	fmt.Println(a1)
	// const b1 = math.Sqrt(4) // 报错  const initializer math.Sqrt(4) is not a constant
	// fmt.Println(b1)
	var name = "kk"
	fmt.Printf("type %T value %v\n", name, name)

	//字符串类型常量
	var defaultName = "make"
	type myString string
	var customName myString = "make"
	// customName = defaultName // 不允许 报错 cannot use defaultName (type string) as type myString in assignment
	fmt.Println(defaultName, customName)

	// 创建变量 defaultName  字符串类型为string
	// 创建一个myString 的新类型 它是字符串的别名 ，然后创建myString 的变量customName

	// 布尔类型常量，规则类似常量字符串

	// const trueConst = true
	// type myBool bool
	// var defaultBool = trueConst
	// var customBool myBool = trueConst
	// defaultBool=customBool

}

func main() {
	fmt.Println("hello world")
	constest1()
}
