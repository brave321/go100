package main

import "fmt"

// 什么是channel

channel 用于线程间的通信

// channel 如何定义

// 示例

//  使用场景


func main() {

	ch1 :=make(chan int)
	ch2 :=make(chan interface{})
	type Euip struct {

	}

	ch3 :=make(chan *Euip)
	fmt.Println(ch1,ch2,ch3)
}
