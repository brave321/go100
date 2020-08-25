package main

import (
	"fmt"
	"math"
)

//匿名 结构体

// 1、匿名结构体 就是没有名字的结构体,
// 无须通过type关键字定义就可以直接使用
// 创建匿名 结构体时，同时要创建对象。
// 匿名结构体由结构体定义和键值对初始化两部分组成

func main() {
	fmt.Println("hello")
	// 匿名函数
	res := func(a, b float64) float64 {
		return math.Pow(a, b)
	}(2, 3)
	fmt.Println(res)
	// 匿名 结构体
	add := struct {
		province, city string
	}{"陕西声", "西安市"}
	fmt.Println(add)

	cat := struct {
		name, color string
		age         int8
	}{
		name:  "绒毛",
		color: "黑白",
		age:   1,
	}
	fmt.Println(cat)

	// 结构体的匿名字段

	// 实列化结构体

}
