package main

import (
	"encoding/json"
	"fmt"
)

// struct 场景示例,json 序列化操作

type Student struct {
	Name string `"json:name"`
	Age  int    `"json:age"`
}

func main() {
	var stu = Student{Name: "wd", Age: 22}
	data, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("json encode failed err:", err)
		return
	}
	fmt.Println(string(data))
}
