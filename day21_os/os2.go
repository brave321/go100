package main

import (
	"fmt"
	"log"
	"os"
)

func main() {

	key := "DB_CONN"
	fmt.Println(key)

	connStr, ex := os.LookupEnv(key)
	if !ex {
		log.Printf("环境变量 %s 没有设置\n", key)
	}
	fmt.Println(connStr)
}
