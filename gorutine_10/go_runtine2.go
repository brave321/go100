package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var a [100]int
	fmt.Println("running", runtime.Version())
	for i := 0; i < 100; i++ {
		go func(i int) {
			for {
				a[i]++
				//runtime.Gosched()
			}
		}(i)
	}
	time.Sleep(time.Millisecond)
	fmt.Println(a)
}
