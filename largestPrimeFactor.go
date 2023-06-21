package main

import (
	"fmt"
	"math"
)

func LargestPrimeFactor(n int) int {
	largest := 1

	for n%2 == 0 {
		largest = 2
		n /= 2
	}
	//Check for divisibility by odd numbers starting from 3 upto the //square root of n
	for i := 3; i <= int(math.Sqrt(float64(n))); i += 2 {
		for n%i == 0 {
			largest = i
			n /= i
		}
	}
	//if n is prime number greater than 2, it is the largest prime factor
	if n > 2 {
		largest = n
	}
	return largest
}
func main() {
	number := 14 //Change this to desired number
	largest := LargestPrimeFactor(number)
	fmt.Printf("The largest prime factor of %d is %d\n", number, largest)
}
