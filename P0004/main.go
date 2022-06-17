package main

import (
	"fmt"
	"strconv"
)

func main() {

	fmt.Println("----   9009 ", isPalendrome001(9009))
	fmt.Println("----   9008 ", isPalendrome001(9008))

	i1, i2, p := maxProd()
	fmt.Println("v1 - max pal is ", i1, " * ", i2, " = ", p)

	i1, i2, p = maxProd02()
	fmt.Println("v2 - max pal is ", i1, " * ", i2, " = ", p)

}

func maxProd02() (int, int, int) {
	max := 999
	i1 := 0
	i2 := 0
	for i := 999; i >= 100; i-- {
		for j := 999; j >= 100; j-- {
			product := i * j
			if isPalendrome001(product) && product > max {
				max = product
				i1 = j
				i2 = i
				return i1, i2, max // why is this returning wrong numbers?  I see why.
			}
		}
	}

	return i1, i2, max
}

func maxProd() (int, int, int) {
	max := 999
	i1 := 0
	i2 := 0
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			product := i * j
			if isPalendrome001(product) && product > max {
				max = product
				i1 = i
				i2 = j
			}
		}
	}

	return i1, i2, max
}

func isPalendrome001(n int) bool {
	//  convert to string
	// seems simplest, but prob not optimal

	a1 := strconv.Itoa(n)
	//fmt.Println("===", a1)
	a2 := ""
	//res := false

	for i := len(a1) - 1; i >= 0; i-- {
		//fmt.Println(string(a1[i]))
		a2 += string(a1[i])
	}

	//fmt.Println("====== +", a2)
	//if a1 == a2 {
	//	res = true
	//} else {
	//	res = false
	//}
	//return res

	return a1 == a2

}

func isPalendrome002(n int) bool {
	//  create new numeric in reverse order

	//reverse := 0
	//original := n

	for n != 0 {
		//reverse = (reverse * 10) +
	}

	return true
}
