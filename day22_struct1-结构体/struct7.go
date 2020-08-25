package main

import "fmt"

type User struct {
	string
	byte
	int8
	float64
}

func main() {
	// 实列化 结构体
	user := User{"steven", 'm', 35, 177.5}
	fmt.Println(user)

	// 如果想依次输出 姓名、年龄、身高、性别
	fmt.Printf("姓名：%s \n", user.string)
	fmt.Printf("身高：%.2f \n", user.float64)
	fmt.Printf(" 性别 %c \n", user.byte)
	fmt.Printf(" 年龄 %d \n", user.int8)

}
