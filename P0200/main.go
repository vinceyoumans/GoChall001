package main

import "fmt"

const (
	x = 7
	y
	z
)

func main() {
	//fmt.Println(x, y, z)
	//
	//primeNums := [6]int{2, 3, 5, 7, 11, 13}
	//slice1 := primeNums[1:4]

	//fmt.Println(slice1)

	//var x int
	//arr := [3]int{3, 5, 2}
	//x, []int = arr[0], arr[1:]
	//fmt.Println(x, arr)  // 3.  5, 2

	//arr1 := [2]int{2, 3}
	//arr2 := [...]int{2, 3}
	//fmt.Println(arr1 == arr2)
	//
	a := 1
	b := 1
	res := 2

	for x := 0; x < 50; x++ {
		res = a + b
		fmt.Println(res)
		a = b
		b = res
	}

}
