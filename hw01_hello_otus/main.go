package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
)

const (
	inputText = "Hello, OTUS!"
)

func main() {
	// Place your code here.

	fmt.Println(reverse.String(inputText))
}
