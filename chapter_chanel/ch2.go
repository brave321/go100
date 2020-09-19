package main

import (
	"fmt"
)

func change_test() {
	var ch1 chan int
	fmt.Println(ch1)
	if ch1 == nil {
		fmt.Println("nil ....nil")
	}
	// chanel 是引用类型
	ch2 := make(chan int)
	fmt.Println(ch2)
	//
	ch3 := make(chan interface{})
	fmt.Println(ch3)
	//
	type Eurip struct {
		name string
	}

	ch4 := make(chan *Eurip)

	fmt.Println("ch4 is", ch4)

	// 使用chanel 发送数据
	// 通过channel 发送数据需要使用特殊的操作符号
	// <-
	// channel发送的值的类型必须与channel的元素类型一致。
	//如果接收方一直没有接收，那么发送操作将持续阻塞
	//使用channel时要考虑发生死锁（deadlock）的可能。
	//如果Goroutine在一个channel上发送数据，
	//其他的Goroutine应该接收得到数据；如果没有接收，
	//那么程序将在运行时出现死锁。
	//如果Goroutine正在等待从channel接收数据
	//，其他一些Goroutine将会在该channel上写入数据；
	//如果没有写入，程序将会死锁。

}

func change_receive1() {
	ch1 := make(chan string)
	go sendData(ch1)

	for {
		data := <-ch1
		// 如果通道关闭，通道中传输的数据则为各数据类型的默认值
		// chan int 的默认值为0
		if data == "" {
			break
		}
		fmt.Println("从通道中读取的数据方式为1", data)
	}
}

// 循环接收数据方式2
func change_receive2() {
	ch1 := make(chan string)
	for {
		data, ok := <-ch1
		fmt.Println(ok)
		if !ok {
			break
		}
		fmt.Println("从通道中读取的数据方式2:", data)
	}
}

func sendData(ch1 chan string) {
	defer close(ch1)
	for i := 0; i < 4; i++ {
		ch1 <- fmt.Sprintf("发送数据%d\n", i)
	}
	fmt.Println("发送数据完毕。。。。。。")

}

func change_receive3() {
	ch1 := make(chan string)
	for value := range ch1 {
		fmt.Println("从通道中读取的数据方式3", value)
	}

}

func main() {
	//fmt.Println("hello world")
	//change_test()
	//change_receive1()
	//change_receive2()
	change_receive3()
}
