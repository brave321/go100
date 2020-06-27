package main
import "fmt"

// struct 声明和定义
// go 中面向对象是通过struct 来实现的，struct 是用户自定义的类型

type People struct {
	Username string
	Age int

}

// struct1 初始化方法1 
func structtest1()  {
	var p1 People
	p1.Username="zhangsan"
	p1.Age = 18
	fmt.Printf("name is %s age is %d",p1.Username,p1.Age)
}

func main()  {
	structtest1()
}