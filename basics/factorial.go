package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)

	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}

	fmt.Println("Factorial:", result)
}