package main

import "fmt"

type Human struct {
	name, phone string
	age         int
}

type Student struct {
	Human
	school string
}

type Employee struct {
	Human
	company string
}

func (h *Human) SayHi() {
	fmt.Printf("大家好！我是%s ,我 %d,我的联系方式是: %s\n", h.name, h.age, h.phone)
}

func main() {
	fmt.Println("hello world")
	s1 := Student{Human{"Danile", "15032...", 13}, "十一中学"}
	e1 := Employee{Human{"kk", "178213...", 32}, "1000phone"}
	s1.SayHi()
	e1.SayHi()
	// 输出
	//大家好！我是Danile ,我 13,我的联系方式是: 15032...
	//大家好！我是kk ,我 32,我的联系方式是: 178213...

}
