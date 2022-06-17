package main

import "fmt"

func main() {

	v1()
	fmt.Println("---------  v2")
	v2()

}

func v2() {
	max := 2520
	for i := max; ; i++ {
		can := true
		for j := 1; j <= 20; j++ {
			if i%j != 0 {
				can = false
				break
			}
		}
		if can {
			max = i
			break
		}
	}
	fmt.Println(max)
}

//===============================
func v1() {

	fmt.Println(whatIsRemainder(100, 10))
	fmt.Println(whatIsRemainder(99, 10))

	for x := 20; true; x++ {

		norem := divBy1to20(x)

		if norem {
			// at the lowest number
			fmt.Println("========== ", x)
			return
		}

	}

}

func divBy1to20(x int) bool {

	for d := 1; d <= 20; d++ {
		_, s := whatIsRemainder(x, d)
		if s == false {
			return false
		}
	}
	return true
}

func whatIsRemainder(n, denom int) (int, bool) {
	remainder := 0
	noRemainder := false
	remainder = n % denom
	if remainder == 0 {
		noRemainder = true
	} else {
		noRemainder = false
	}
	return remainder, noRemainder
}
