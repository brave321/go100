package main

import (
	"fmt"
	"regexp"
	"strconv"
)

// strconv 包函数转换

func strconvtest() {
	fmt.Println(strconv.ParseBool("1"))
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	// 	true <nil>
	// true <nil>
	// true <nil>
}

func strconvtest2() {
	// 目标字符串
	searchIn := "John: 2578.34 William: 4567.23 Steve: 5632.18"
	pat := "[0-9]+.[0-9]+"

	f := func(s string) string {
		v, _ := strconv.ParseFloat(s, 32)
		return strconv.FormatFloat(v*2, 'f', 2, 32)
	}
	if ok, _ := regexp.Match(pat, []byte(searchIn)); ok {
		fmt.Println("Match Found!")
	}

	re, _ := regexp.Compile(pat)
	//将匹配到的部分替换为"##.#"
	str := re.ReplaceAllString(searchIn, "##.#")
	fmt.Println(str)
	//参数为函数时
	str2 := re.ReplaceAllStringFunc(searchIn, f)
	fmt.Println(str2)
	// 输出
	//    Match Found!
	//    John: ##.# William: ##.# Steve: ##.#
	//    John: 5156.68 William: 9134.46 Steve: 11264.36

}

// 锁和 sync 包

import "sync"




func main() {
	fmt.Println("hello world")
	// strconvtest()
	// strconvtest2()

}
