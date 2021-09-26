package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("This program solves the square equation a*x^2+b*x+c=0")
	fmt.Print("Enter a, b and c: ")
	var a, b, c float64
	fmt.Scan(&a, &b, &c)
	if a == 0 {
		// Solving as a linear equation.
		// ...
		return // <- early return
	}
	D := b*b - 4*a*c
	if D < 0 {
		fmt.Println("No real roots")
	} else if D == 0 {
		x := (-b + math.Sqrt(D)) / (2 * a)
		fmt.Println("x1, x2 =", x)
	} else {
		x1 := (-b + math.Sqrt(D)) / (2 * a)
		x2 := (-b - math.Sqrt(D)) / (2 * a)
		fmt.Println("x1 =", x1)
		fmt.Println("x2 =", x2)
	}
}
