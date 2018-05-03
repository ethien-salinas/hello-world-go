package main

import (
	"fmt"

	"github.com/ethien-salinas/hello-world-go/stringutil"
)

var reverseString = "!oG gninrael ma I 😎"

func main() {
	fmt.Println("Hello Ethien! ☕️")
	fmt.Printf(stringutil.Reverse(reverseString))
}
