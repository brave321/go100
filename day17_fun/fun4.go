package main

import "fmt"

// 若增加切片的容量 ，必须创建一个更大的切片

func s1test() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to, n, "===")

	fmt.Printf("Copied %d elements\n", n)
	fmt.Println(sl_from, len(sl_from), sl_to)

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)
}

// append 追加 元素

func AppendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

// 从字符串生成自切片

func str1() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c\n", i, c)
	}
	// 输出值为  0:ÿ2:界
	// 修改字符串中的某个字符
	s2 := "hello"
	c1 := []byte(s2)
	c1[0] = 'c'
	s3 := string(c1)
	fmt.Println("s2", s2, "c1==", c1, "s3===", s3)
}

// 字节数组对比函数

// func Compare(a, b []byte) int {
// 	for i:=0; i < len(a) && i < len(b); i++ {
//         switch {
//         case a[i] > b[i]:
//             return 1
//         case a[i] < b[i]:
//             return -1
// 		}
// 	// 判读数组长度
//      switch {
// 	 case len(a)<len(b):
// 		return - 1
// 	 }
// 	case len(a) > len(b):
// 		return 1
// }
// return 0
// }

// 练习
// 编写一个函数，要求其接受两个参数，原始字符串 str 和分割索引 i，然后返回两个分割后的字符串。

// 练习 7.13

// 假设有字符串 str，那么 str[len(str)/2:] + str[:len(str)/2] 的结果是什么？

func str2test() {
	s1 := "hello go"
	fmt.Println(len(s1))

	for i, k := range s1 {
		fmt.Println(i, s1[i], k)
	}
}

// go 字符串反转

func reverse(str string) string {
	var result string
	strLen := len(str)
	for i := 0; i < strLen; i++ {
		result = result + fmt.Sprintf("%c", str[strLen-1])
	}
	return result
}

func main() {
	// fmt.Println("hello world")
	// s1test()
	// str1()
	// str2test()
	str3 := "Goog"
	result := reverse(str3)
	fmt.Println(result)
}
