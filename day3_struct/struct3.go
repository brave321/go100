package main
import "fmt"

// 内嵌结构体可以来自其他包

type A struct{
	ax,ay int
}

type B struct{
	A 
	bx,by float32 
}


type C struct{
   value int
	cx,cy string
}


func structtest1()  {
	b := B{A{1,11},3.4,5.0}
	fmt.Println(b,b.ax,b.ay) 
	//输出 {{1 11} 3.4 5} 1 11
	fmt.Println(b.A)
	// 输出 {1 11}
}

func structtest2()  {
	c1 :=C{12,"hello","go"}
	fmt.Println(c1)
	// 输出 {12 hello go}
}

// 命名冲突
// 1、外层名字会覆盖内层名字（但是两者的内层空间保留），这提供来一种重载字段或方法的方式
//如果相同的名字在同一级别出现了两次，如果这个名字被程序使用了，将会引发错误




func main()  {
	fmt.Println("hello world")
	// structtest1()
	// structtest2()
}

