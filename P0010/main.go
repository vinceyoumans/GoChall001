package main

import (
	"fmt"
	"math"
)

const (
	ICount = 2000000
)

func main() {

	var res = 0
	for i := 2; i <= ICount; i++ {

		if isPrime(i) {
			fmt.Println(i)
			res += i
		}
	}
	fmt.Println("---  res = ", res)
}

func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
