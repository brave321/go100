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
	fmt.Printf("name is %s age is %d\n",p1.Username,p1.Age)
}

//struct1 初始化方法2
func  structtest2()  {
	p2 :=People{
		Username: "山海经",
		Age: 500,
	}
	fmt.Println(p2.Username,p2.Age) // 输出 山海经 500
}

//struct 初始化的默认值

func structtest3 ()  {
	var p3 People
	fmt.Printf("%#v\n",p3) // 输出值 main.People{Username:"", Age:0} // int 初始化的值是0, string 的初始化默认值是“”
}

// 结构体类型的指针

// 什么是结构体 , go 语言中数组可以存储同一类型的数据，但在结构体中我们可以为不同项定义不同的数据类型

// 结构体表示一项记录

// 结构体类型的指针

func structtest4()  {
	var p4 *People = &People{}
	// %v 默认格式 
	// %#p 16 进制格式
	// %#v 相应值的go 语法表示
	fmt.Printf("%p  %#v\n",p4,p4)  // 输出 0xc000092000  &main.People{Username:"", Age:0}
}

func structtest5()  {
	var p5 *People = &People{
		Username : "lunyu",
		Age: 2000,
	}
	fmt.Println(p5.Username,p5.Age) // 输出 lunyu 2000
	fmt.Printf("%p %#v\n",p5,p5) //输出 0xc000074020 &main.People{Username:"lunyu", Age:2000}

}

func structtest6()  {
	var p6 People=new(People)
	p6.Username = "mengzi"
	p6.Age = 1900
	fmt.Println(p6.Username,p6.Age) 
}




func main()  {
	// structtest1()
	// structtest2()
	// structtest3() 
	// structtest4()
	// structtest5()
	structtest6()
}