package main

import "fmt"

// defer

// defer 允许我们推迟到函数返回之前（或任意位置执行return 语句之后）一刻才执行某个语句或者函数

// 示例

func function1() {
	fmt.Printf("In function1 at the top\n")

	defer function2()

	fmt.Printf("In function1 at the bottom!\n")

}

func function2() {
	fmt.Printf("function2: Deferred until the end of the calling function!\n")
}

// 示例2
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// 当有多个 defer 行为被注册时，它们会以逆序执行（类似栈，即后进先出）：

func f() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d ", i)

	}
}

// 模拟go 连接数据库情况
func connectToDB() {
	fmt.Println("ok, connected to db")
}

func disconnectFromDB() {
	fmt.Println("ok, disconnected from db")
}

func doDBOperations() {
	connectToDB()
	fmt.Println("Defering the database disconnect.")
	defer disconnectFromDB()
	fmt.Println("Doing some DB operations ...")
	fmt.Println("Oops! some crash or network error ...")
	fmt.Println("Returning from function here!")
	return

}

func main() {
	// fmt.Println("hello ")
	// function1()

	// 输出，最后输出defer 的值
	// 	In function1 at the top
	// In function1 at the bottom!
	// function2: Deferred until the end of the calling function!
	// a()
	// f()
	// 输出
	//4 3 2 1 0
	doDBOperations()
	// 输出，最后执行 defer disconnectFromDB() // 有多个defer 的状态时候，执行后进先出
	// 执行

	// 	ok, connected to db
	// Defering the database disconnect.
	// Doing some DB operations ...
	// Oops! some crash or network error ...
	// Returning from function here!
	// ok, disconnected from db

}
