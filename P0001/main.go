package main

import "fmt"

func main() {
	sum := 0
	tn := 1000
	fmt.Println(" hi there")

	for i := 0; i < tn; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}

	fmt.Println("the sum of these numbers is", sum)

}
