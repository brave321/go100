package main

import (
	"fmt"
	"time"
)

// 并发：同一时间内执行多个操作

// 并行： 同一时刻执行多个操作

// 多线程

// 线程是由操作系统进行管理，也就是处于内核态

// 现成之间进行切换，需要发生用户态到内核态的切换

// 当系统中✅大量线程，系统会变得非常慢

// 用户态的线程，支持大量线程创建，也就协程或goroutine

// 创建goroutine

func hello() {
	fmt.Println("hello wolrd goroutine")
}

// 启动多个goroutine

func numbers() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Millisecond)
		fmt.Printf("%d\n", i)
	}
}

func loop() {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d \n", i)
	}
}

func main() {
	go loop()
	//go numbers()
	//go hello()
	//time.Sleep(time.Millisecond)
	fmt.Println("main function")
	//fmt.Println("hello")
}
