package main

import (
	"flag"
	"fmt"
	"net/url"
	"strings"

	"github.com/b09780978/httptool/cmd/httptool"
)

func validMethod(method string) bool {
	switch method {
	case httptool.GET:
		return true
	case httptool.POST:
		return true
	case httptool.PUT:
		return true
	case httptool.DELETE:
		return true
	case httptool.HEAD:
		return true
	case httptool.OPTIONS:
		return true
	}
	return false
}

func main() {
	var method string
	var urlStr string
	var Url *url.URL
	flag.StringVar(&method, "method", httptool.GET, "HTTP method, default is GET")
	flag.StringVar(&urlStr, "url", "", "Url")

	flag.Parse()

	method = strings.ToUpper(method)
	if !validMethod(method) {
		fmt.Printf("Invalid method: %s\n", method)
		return
	}

	Url, err := url.Parse(urlStr)

	if err != nil {
		fmt.Printf("Invalid url: %s\n", urlStr)
		return
	}

	fmt.Printf("method: %s\n", method)
	fmt.Printf("url: %v\n", Url)

	//client := httptool.New()

}
