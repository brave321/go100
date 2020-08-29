package main

import (
	"fmt"
	"time"
)

func chanDemo() {
	c := make(chan int)
	go func() {
		for {
			n := <-c
			fmt.Println(n)
		}
	}()
	c <- 1
	c <- 2
	time.Sleep(time.Millisecond)
	//n :=<-c
	fmt.Println()

}

func main() {
	chanDemo()
}
