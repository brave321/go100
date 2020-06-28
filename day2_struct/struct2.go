package main
import "fmt"

// 匿名字段和内嵌结构体

// 结构体可以包含一个或多个匿名（或内嵌）字段
// 匿名字段本身可以是一个结构体类型 结构体可以包含内嵌结构体

type inner_val struct{
	value1 int
	value2 int
}

type outer_val struct{
	b int
	c float32
	int // 匿名字段
	inner_val //匿名字段
}

func structtest1()  {
	outer :=new(outer_val)
	outer.b = 1
	outer.c = 12.3
	outer.int = 20
	outer.value1 = 5
	outer.value2=10 
   // 输出
	fmt.Printf("outer.b is %d\n",outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
    fmt.Printf("outer.int is: %d\n", outer.int)
    fmt.Printf("outer.value1 is: %d\n", outer.value1)
	fmt.Printf("outer.value2 is: %d\n", outer.value2)
	
	// 输出
// 	outer.b is 1
// outer.c is: 12.300000
// outer.int is: 20
// outer.value1 is: 5
// outer.value2 is: 10

//  使用结构体字面量
outer2 := outer_val{6, 7.5, 60, inner_val{5, 10}}
fmt.Println("outer2 is:", outer2)

// 输出 
// outer2 is: {6 7.5 60 {5 10}}

}



func main()  {
	fmt.Println("hello go")
	structtest1()
}