package main

import (
	"log"
	"os"
)

// 原理 
// 环境变量的获取和设置分别可以通过 os 包中 Getenv 和 Setenv 方法实现。方法名称已经很明确说明了自身的功能。

// 但 Getenv 方法有个缺点，即使在未设置环境变量的情况下，它也返回一个空的字符串。

// os 包中还有一个很有用的方法， LookupEnv ，该方法返回两个值，一个是变量的值，另一个则是变量在环境中是否设置的布尔值。而 LookupEnv 方法则可以克服上面 Getenv 的缺点。

// 当我们要判断是否设置了环境变量以及实现一个默认环境变量的方法，都应该使用 LookupEnv 。因为如果未设置环境变量，则第二个值会返回 false


func GetEnvDefault(key, defVal string) string {
	val, ex := os.LookupEnv(key)
	if !ex {
		return defVal
	}
	return val
}

func main() {
	key := "DB_CONN"
	// 设置环境变量
	os.Setenv(key, "postgres://as:as@example.com/pg?sslmode=verify-full")
	val := GetEnvDefault(key, "postgres://as:as@localhost/pg?sslmode=verify-full")
	log.Println("值是 :" + val)

	os.Unsetenv(key)
	val = GetEnvDefault(key, "postgres://as:as@127.0.0.1/pg?sslmode=verify-full")
	log.Println("默认值是 :" + val)

}
