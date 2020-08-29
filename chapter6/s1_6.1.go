package main

import (
	"fmt"
	"unicode/utf8"
)

// 字符串处理包介绍

// 因此熟悉Go的strings包后基本就能借此封装符合自己需求的、应用于特定场景的字符串处理函数了。
// 而strconv包实现了字符串与其他基本数据类型之间的类型转换

func main() {
	fmt.Println("字符串处理包介绍")
	s := "我爱go语言"
	fmt.Println("字符串长度", len(s))
	fmt.Println("==============")
	// for range 遍历循环字符串
	len := 0
	for i, ch := range s {
		fmt.Printf("%d:%c \ns", i, ch)
		len++
		//fmt.Println(len)
	}

	fmt.Println("\n 字符串长度", len)
	fmt.Println("==============")

	// 遍历所有字节
	for i, ch := range []byte(s) {
		fmt.Printf("%d:%x\n", i, ch)
	}
	fmt.Println()
	fmt.Println("==============")
	// 遍历所有字符
	count := 0
	for i, ch := range []rune(s) { //rune
		fmt.Printf("%d:%c", i, ch)
		count++
	}
	fmt.Println()
	fmt.Println("字符串长度", count)
	fmt.Println("字符串长度", utf8.RuneCountInString(s))

}

// 如果字符串涉及中文，遍历字符推荐使用rune。
//因为一个byte存不下一个汉语文字的unicode值
