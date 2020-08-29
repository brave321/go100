package main

import (
	"fmt"
	"strings"
	//"unicode"
)

// 6.2 strings包的字符串处理函数

// 判断字符串的常用方法

func TestContainers() {
	fmt.Println(strings.Contains("seafood", "foo"))
	fmt.Println(strings.Contains("seafood", "bar"))
	fmt.Println(strings.Contains("seafood", ""))
	fmt.Println(strings.Contains("", ""))
	fmt.Println(strings.Contains("steven土2020", "土"))
	// 输出
	//true
	//false
	//true
	//true
	//true

}

// 判读字符串中是否包含另一个字符串中的任一字符

func TestContainersAny() {
	fmt.Println(strings.ContainsAny("team", "i"))
	fmt.Println(strings.ContainsAny("failure", "u & i"))
	fmt.Println(strings.ContainsAny("failure", ""))
	fmt.Println(strings.ContainsAny("", ""))
	// 输出
	//false
	//true
	//false
	//false
}

// 判断字符串是否包含unicode 码值
func TestContainsRune() {
	fmt.Println(strings.ContainsRune("一丁", '丁'))
	fmt.Println(strings.ContainsRune("一丁", 19969))
	// 输出
	//true
	//true
}

// 返回字符串包含另一个字符串的个数

func main() {
	//TestContainers()
	//TestContainersAny()
	TestContainsRune()
	//
	//false
	//true
	//false
	//false

}
