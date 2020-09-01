package main

import "fmt"

// fmt 打印

//打印格式     打印内容
//%v          值的默认表示
//%+v         类似%v,但输出结构体时会添加字段名
//%#v         值的go 语法表示
//%T          值的类型的go 语法表示


// 示例
func main()  {
	str :="steven"
	fmt.Printf("%T,%v,\n",str,str)
}
\


