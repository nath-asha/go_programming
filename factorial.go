package main

import (
	"fmt"
)

func Factorial(n int) int {
	fact := 1
	i := 1
	for i = 1; i <= n; i++ {
		fact *= i
	}
	return fact
}
func main() {
	number := 5
	factorial := Factorial(number)
	fmt.Printf("The factorial of %d is %d\n", number, factorial)
}
