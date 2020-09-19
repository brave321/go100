package main

import (
	"fmt"
	"time"
)

func main() {
	//fmt.Println("hello")
	for i := 0; i < 10000; i++ {
		go func(i int) {
			for {
				fmt.Println("hello from"+"goruntine %d", i)
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
}
