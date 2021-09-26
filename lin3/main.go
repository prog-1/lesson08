package main

import "fmt"

func main() {
	fmt.Println("This program solves ax*b=0 equation")
	fmt.Print("Enter a and b: ")
	var a, b float64
	fmt.Scan(&a, &b)
	if a == 0 {
		fmt.Println("ERROR: a must not be zero")
	} else {
		x := -b / a
		fmt.Println("x =", x)
	}
}
