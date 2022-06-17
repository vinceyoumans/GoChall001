package main

import (
	"fmt"
	"math"
)

const FindPrimNumner = 10000

func main() {

	fmt.Println("=======  version 1 =========")
	// just checking if isPrime Works
	fmt.Println("2 - is prime", isPrime(2)) // yes
	fmt.Println("4 - is prime", isPrime(4)) // no

	fmt.Println("v1 answer = ", xPM(FindPrimNumner))

	// Another Approach
	fmt.Println("========   version two... is this better?  ===")
	fmt.Println(" 2 - is Prime", isPrime02(2))
	fmt.Println(" 4 - is Prime", isPrime02(4))

	fmt.Println("v2 answer = ", xpm2(FindPrimNumner))

}

//============================================

func xPM(numner int) int {
	pnc := 0
	pn := 0
	for x := 2; pnc <= FindPrimNumner; x++ {
		if isPrime02(x) {
			pnc++
			pn = x
		}
	}

	return pn
}

//=====

func xpm2(n int) int {
	if n == 0 {
		return 0
	}
	count := 0
	i := 2
	for ; ; i++ {
		if isPrime02(i) {
			count++
		}
		if count == n {
			break
		}
	}
	return i

}

//===========================
func isPrime(value int) bool {
	for i := 2; i <= int(math.Floor(float64(value)/2)); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}

//===========
// another Prime approach
func isPrime02(value int) bool {
	for i := 2; i <= int(math.Floor(math.Sqrt(float64(value)))); i++ {
		if value%i == 0 {
			return false
		}
	}
	return value > 1
}
