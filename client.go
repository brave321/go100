package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func main() {
	fmt.Print("hello http response")
	resp, err := http.Get("http://www.k8sre.com")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", s)
}
