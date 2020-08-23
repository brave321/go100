package main

import (
	"fmt"
	"log"
	"os"
)

// LookupEnv 检查环境变量
func main() {
	fmt.Println(os.Getwd())
	key := "DB_CONN"
	os.Setenv(key, "DB_CONN")
	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("环境变量 %s 没有设置\n", key)
	}
	fmt.Println(connStr)
}
x