package main

import (
	"fmt"
)

func main() {
	fmt.Print("Enter two numbers: ")
	var a, b int
	fmt.Scan(&a, &b)
	if b == 0 {
		fmt.Println("ERROR: unable to divide by 0")
		return
	}
	fmt.Println("a / b =", a/b)
}
