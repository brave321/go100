package main

//在Go语言中处理错误的方式通常是将返回的错误与nil进行比较。
//nil值表示没有发生错误，而非nil值表示出现错误。如果不是nil，需打印输出错误。
import (
	"fmt"
	"math"
)

func main() {
	// 异常情况1
	res := math.Sqrt(-100)
	fmt.Println(res)
	res, err := Sqrt(-100)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}
}
