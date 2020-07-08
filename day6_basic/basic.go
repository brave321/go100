package main

import "fmt"

// go 语言数据类型

// 基本数据类型

// 数据类型 int float32 float64

// 布尔值 bool true flase

// 字符串类型 string

func basic_test1() {
	var i int
	var b1 bool
	var s1 string
	fmt.Println(i, b1, s1)
	fmt.Printf("the default value is %d %s %q", i, b1, s1)
	// 输出
	// 	0 false
	// the default value is 0 %!s(bool=false) ""
}

// bool 类型

func booltest() {
	a := true
	b := false
	fmt.Println("a is", a, "b is", b)
	c := a && b //逻辑 AND 运算符。 如果两边的操作数都是 True，则条件 True，否则为 False。
	fmt.Println("c is", c)

	d := a || b // 逻辑 OR 运算符。 如果两边的操作数有一个 True，则条件 True，否则为 False。
	fmt.Println("d is", d)
	// 输出
	// 	a is true b is false
	// c is false
	// d is true

}

// 有符号整形
// int：根据不同的底层平台（Underlying Platform），表示 32 或 64 位整型。除非对整型的大小有特定的需求，否则你通常应该使用 int 表示整型。
// 大小：在 32 位系统下是 32 位，而在 64 位系统下是 64 位。
// 范围：在 32 位系统下是 -2147483648～2147483647，而在 64 位系统是 -9223372036854775808～9223372036854775807

func inttest() {
	var a int = 100
	b := 87
	fmt.Printf("the value of a is %d b is %d", a, b)
	// 输出
	// the value of a is 100 b is 87
}

// 浮点类型
func floattest() {
	a, b := 3.33, 8.95
	fmt.Printf("type of a is %T b is %T", a, b)
	// output
	// type of a is float64 b is float64
}

// 字符串类型

func stringtest1() {
	first := "java"
	second := "c++"
	third := "go"
	language := first + " " + second + " " + third
	fmt.Println(language)
	// output
	//java c++ go
}

// 类型转换

func changevalue() {
	i := 45    // int
	j := 78.89 //float
	// sum := i + j // 直接加报错 invalid operation: i + j (mismatched types int and float64)
	sum := i + int(j)
	fmt.Println(sum)
	// output 123

}

// 变量的声明

func basictest2() {
	var age int
	fmt.Println("the age is", age)
	age = 18 // 赋值
	fmt.Println("the age is", age)
	// 声明变量并初始化
	var age2 int = 19
	fmt.Println("the age2 is", age2)
	// 类型推断赋值
	var age3 = 27
	fmt.Println("the age3 is", age3)

	// 短变量声明 ，操作符的左边至少有一个变量是尚未声明的
	age4 := 28
	fmt.Println("the age4 is", age4)
	a, b := 10, 32
	fmt.Printf("a is %d b is %d\n", a, b)
	b, c := 11, 23
	fmt.Printf("b is %d c is %d\n", b, c)

	//声明多个变量
	var (
		name1 = "kk"
		name2 = "jenkins"
	)
	fmt.Println(name1, name2)

	// 输出
	// 	the age is 0
	// the age is 18
	// the age2 is 19
	// the age3 is 27
	// the age4 is 28
	// a is 10 b is 32
	// b is 11 c is 23
	// kk jenkins

}

// 常量

func main() {
	// basic_test1()
	// basictest2()
	// booltest()
	// inttest()
	// floattest()
	// stringtest1()
	// changevalue()
}
