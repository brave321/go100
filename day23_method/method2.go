package main

import (
	"fmt"
	"math"
)

// 三、方法和函数

//一段程序可以用函数来写，却还要使用方法，主要有以下两个原因。
//• Go不是一种纯粹面向对象的编程语言，它不支持类。
//因此其方法旨在实现类似于类的行为。
//• 相同名称的方法可以在不同的类型上定义，
//而具有相同名称的函数是不允许的。
//假设有一个正方形和一个圆形，可以分别在正方形和圆形上定义一个名为Area的求取面积的方法。

type Rectangle struct {
	width, height float64
}

type Circle struct {
	radius float64
}

// 定义Rectangle 的方法
// 方法 （接受者变量 接受者类型) 方法名（参数列表）（返回值类型）

func (r Rectangle) Area() float64 {
	return r.width * r.height
}

// 定义Circle 的方法
func (c Circle) Area() float64 {
	return c.radius * c.radius * math.Pi
}

func main() {
	r1 := Rectangle{10, 4}
	r2 := Rectangle{12, 5}
	c1 := Circle{1}
	c2 := Circle{10}
	fmt.Println("r1 的面积", r1.Area())
	fmt.Println("r2 的面积", r2.Area())
	fmt.Println("c1 的面积", c1.Area())
	fmt.Println("c2 的面积", c2.Area())
	// 输出
	//r1 的面积 40
	//r2 的面积 60
	//c1 的面积 3.141592653589793
	//c2 的面积 314.1592653589793
	fmt.Println("================")

}
