package main

import "fmt"

const maxNumbers = 100

func main() {
	fmt.Println(SoS())

	res := SquareOfSum() - SoS()

	fmt.Println("===   res ", res)

	fmt.Println("========  v2")
	// better solution
	sumSquares := 0
	squaresOfTheSum := 0
	for i := 1; i <= 100; i++ {
		sumSquares += i * i
		squaresOfTheSum += i
	}
	diffSums := (squaresOfTheSum * squaresOfTheSum) - sumSquares
	fmt.Println(diffSums)

}

func SquareOfSum() int {
	r := 0
	for x := 1; x <= maxNumbers; x++ {
		r += x
	}
	return r * r
}

func SoS() int {
	ans := 0
	for x := 1; x <= maxNumbers; x++ {
		ans += x * x
	}
	return ans
}
