package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(math.Sqrt(4))
	n := 600851475143
	fmt.Println("=====", int(math.Sqrt(float64(n))))
	printprimes(n)
}

func printprimes(n int) {
	for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			fmt.Println(i)
			n /= i
			i--
		}
	}
	if n > 0 {
		fmt.Println(n)
	}
}
