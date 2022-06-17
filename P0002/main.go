package main

import "fmt"

func main() {
	maxFib := 4000000
	n1 := 0
	n2 := 1
	n3 := 0
	sumEven := 0
	sumAll := 0

	for n3 < maxFib {
		n3 = n1 + n2
		if n3%2 == 0 {
			sumEven += n3
		}
		sumAll += n3
		n1 = n2
		n2 = n3
	}

	fmt.Println("===  the even sum result is", sumEven)
	fmt.Println("===   the sum All is ", sumAll)

}
