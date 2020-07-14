package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bytes, _ := ioutil.ReadFile("test.txt")

	fmt.Println(string(bytes))
}
