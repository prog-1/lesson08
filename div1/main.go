package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter two numbers: ")
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println("a / b =", a/b)
}
