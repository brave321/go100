package main

import (
	"fmt"
	"log"
	"os"
)

//Getenv  检查环境变量

func main() {
	conStr := os.Getenv("DB_CONN")
	fmt.Println(conStr, "====")
	log.Println("连接:", conStr, "====")
}
