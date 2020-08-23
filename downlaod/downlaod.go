package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func retriver(url string) []byte {

}

func main() {
	resp, err := http.Get("www.baidu.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	bytee, _ := ioutil.ReadAll(resp.Body)
	fmt.Printf("%s\n", bytes)
}
