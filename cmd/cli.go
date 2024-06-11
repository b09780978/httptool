package main

import (
	"fmt"

	"github.com/b09780978/httptool/cmd/httptool"
)

func main() {
	client := httptool.HttpClient{}
	fmt.Printf("client: %v\n", client)
	fmt.Printf("client: %v\n", httptool.DefaultClient)
}
