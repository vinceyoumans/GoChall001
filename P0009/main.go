package main

import "fmt"

func main() {

	var output = 0

out:
	for a := 1; a < 1000; a++ {

		for b := 1; b < 1000; b++ {

			for c := 1; c < 1000; c++ {

				f1 := (a*a)+(b*b) == (c * c)
				f2 := (a + b + c) == 1000
				f3 := (a < b)
				f4 := (b < c)

				if f1 && f2 && f3 && f4 {
					fmt.Println("===")
					output = a + b + c
					fmt.Printf("%d + %d +%d = 1000, product = %d\n", a, b, c, output)
					break out
				}
			}
		}
	}
}
