package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	grid = `08 02 22 97 38 15 00 40 00 75 04 05 07 78 52 12 50 77 91 08
49 49 99 40 17 81 18 57 60 87 17 40 98 43 69 48 04 56 62 00
81 49 31 73 55 79 14 29 93 71 40 67 53 88 30 03 49 13 36 65
52 70 95 23 04 60 11 42 69 24 68 56 01 32 56 71 37 02 36 91
22 31 16 71 51 67 63 89 41 92 36 54 22 40 40 28 66 33 13 80
24 47 32 60 99 03 45 02 44 75 33 53 78 36 84 20 35 17 12 50
32 98 81 28 64 23 67 10 26 38 40 67 59 54 70 66 18 38 64 70
67 26 20 68 02 62 12 20 95 63 94 39 63 08 40 91 66 49 94 21
24 55 58 05 66 73 99 26 97 17 78 78 96 83 14 88 34 89 63 72
21 36 23 09 75 00 76 44 20 45 35 14 00 61 33 97 34 31 33 95
78 17 53 28 22 75 31 67 15 94 03 80 04 62 16 14 09 53 56 92
16 39 05 42 96 35 31 47 55 58 88 24 00 17 54 24 36 29 85 57
86 56 00 48 35 71 89 07 05 44 44 37 44 60 21 58 51 54 17 58
19 80 81 68 05 94 47 69 28 73 92 13 86 52 17 77 04 89 55 40
04 52 08 83 97 35 99 16 07 97 57 32 16 26 26 79 33 27 98 66
88 36 68 87 57 62 20 72 03 46 33 67 46 55 12 32 63 93 53 69
04 42 16 73 38 25 39 11 24 94 72 18 08 46 29 32 40 62 76 36
20 69 36 41 72 30 23 88 34 62 99 69 82 67 59 85 74 04 36 16
20 73 35 29 78 31 90 01 74 31 49 71 48 86 81 16 23 57 05 54
01 70 54 71 83 51 54 69 16 92 33 48 61 43 52 01 89 19 67 48`
	//gridRowLen = 20 // columns
	//adjacent   = 4
)

type wheel struct {
	Spoke01 Spoke
	Spoke02 Spoke
	Spoke03 Spoke
	Spoke04 Spoke
	Spoke05 Spoke
	Spoke06 Spoke
	Spoke07 Spoke
	Spoke08 Spoke
}

type Spoke struct {
	//prevTested bool // has this combination been tested in past, if true, can be skipped
	inBounds bool //  There are sufficient nodes to sum.
	dir      int  // Which spoke  1 - 8  ( there are 8 total spokes. 1 is up.  3 is right, 5 is down.  going clockwise
	valSum   int
	a00      arr // The x,y of the line to be counted
	a01      arr
	a02      arr
	a03      arr
}
type arr struct {
	row, col int
}

var MTX [][]int              // may as well make global
var MaxSpoke Spoke           // the top spoke in the wheels
var MPrevTest map[string]int // keeps track of prev tested combinations so they can be skipped.

func (s *Spoke) SIndex() string {
	// returns the unique index for the Spoke

	return "this is the unique index"
}

func main() {
	MPrevTest = make(map[string]int) // Declaired Globally, but still must be initialized

	MaxSpoke := Stage000Begin(grid)

	fmt.Println("==========================  FINISH ===================")
	fmt.Println(" MaxSpoke is:", MaxSpoke)
}

func Stage000Begin(gr string) Spoke {

	// Step 1 - Create MTX[][]int grid
	MTX = Step001_MakeMTXGrid(gr) //works
	var resultSpokes []Spoke

	for r := 0; r < len(MTX); r++ {
		for c := 0; c < len(MTX[r]); c++ {
			resultSpokes = Step002_GetMaxSpoke(r, c)
			for _, aaa := range resultSpokes {

				//fmt.Println("+++++++++++++", aaa.SIndex())
				
				prevIndex := getPrevIndex2(aaa.a03, aaa.dir) // will return row,col of mirror val

				if isAlreadyTested2(prevIndex) {
					//  this combination has already been discovered in previous test.  This is just a reverse sum.
					continue
				} else {
					//add prevIndex to the map of prev tested MTX
					addPrevTested(prevIndex)
				}

				if aaa.inBounds {
					vv := MTX[aaa.a00.row][aaa.a00.col]
					vv *= MTX[aaa.a01.row][aaa.a01.col]
					vv *= MTX[aaa.a02.row][aaa.a02.col]
					vv *= MTX[aaa.a03.row][aaa.a03.col]
					if MaxSpoke.valSum < vv {
						MaxSpoke = aaa
						MaxSpoke.valSum = vv
					}
				}
			}
		}
	}

	return MaxSpoke
}

//===============================================================================
//===============================================================================
func Step001_MakeMTXGrid(gr string) [][]int {
	//m1 := strings.Split(grid, "\n")
	m1 := strings.Split(gr, "\n")
	MTX := make([][]int, len(m1))
	for r := 0; r < len(m1); r++ {
		m2 := strings.Split(m1[r], " ")
		MTX[r] = make([]int, len(m2))
		for c := 0; c < len(m2)-1; c++ {
			MTX[r][c], _ = strconv.Atoi(m2[c])
		}
	}
	return MTX
}

//===============================================================================
//===============================================================================

func addPrevTested(index string) {
	// just adding randomly "2"...  It means nothing.
	//fmt.Println("====   MPrevTested adding - ", index)
	MPrevTest[index] = 2
}

func isAlreadyTested2(pt string) bool {
	_, ok := MPrevTest[pt]
	return ok
}

//=====================
func getPrevIndex2(aaa arr, dir int) string {
	z0 := strconv.Itoa(aaa.row)
	z1 := strconv.Itoa(aaa.col)
	z2 := "0"

	switch dir {
	case 1:
		z2 = strconv.Itoa(5)
	case 2:
		z2 = strconv.Itoa(6)
	case 3:
		z2 = strconv.Itoa(7)
	case 4:
		z2 = strconv.Itoa(8)
	case 5:
		z2 = strconv.Itoa(1)
	case 6:
		z2 = strconv.Itoa(2)
	case 7:
		z2 = strconv.Itoa(3)
	case 8:
		z2 = strconv.Itoa(4)
	}

	zz := fmt.Sprintf("%s-%s-%s", z0, z1, z2)

	return zz
}

//===============================================================================

func Step002_GetMaxSpoke(r, c int) []Spoke {

	//fmt.Println("----   Step002  ---", r, "- ", c)
	var a0 arr
	var a1 arr
	var a2 arr
	var a3 arr

	var sp Spoke
	var sps []Spoke

	for d := 1; d <= 8; d++ {

		sp.dir = d
		switch d {
		case 1:
			a0 = arr{0, 0}
			a1 = arr{-1, 0}
			a2 = arr{-2, 0}
			a3 = arr{-3, 0}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 2:
			a0 = arr{0, 0}
			a1 = arr{-1, 1}
			a2 = arr{-2, 2}
			a3 = arr{-3, 3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 3:
			a0 = arr{0, 0}
			a1 = arr{0, 1}
			a2 = arr{0, 2}
			a3 = arr{0, 3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 4:
			a0 = arr{0, 0}
			a1 = arr{1, 1}
			a2 = arr{2, 2}
			a3 = arr{3, 3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 5:
			a0 = arr{0, 0}
			a1 = arr{1, 0}
			a2 = arr{2, 0}
			a3 = arr{3, 0}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 6:
			a0 = arr{0, 0}
			a1 = arr{1, -1}
			a2 = arr{2, -2}
			a3 = arr{3, -3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 7:
			a0 = arr{0, 0}
			a1 = arr{0, -1}
			a2 = arr{0, -2}
			a3 = arr{0, -3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)

		case 8:
			a0 = arr{0, 0}
			a1 = arr{-1, -1}
			a2 = arr{-2, -2}
			a3 = arr{-3, -3}
			sp.a00, sp.a01, sp.a02, sp.a03, sp.inBounds = stage002B_getSpokeRC(r, c, a0, a1, a2, a3)
		}

		sps = append(sps, sp)
	}
	return sps

}

func stage002B_getSpokeRC(r int, c int, a0, a1, a2, a3 arr) (arr, arr, arr, arr, bool) {
	aa0 := arr{r + a0.row, c + a0.col}
	aa1 := arr{r + a1.row, c + a1.col}
	aa2 := arr{r + a2.row, c + a2.col}
	aa3 := arr{r + a3.row, c + a3.col}

	inBounds := true

	aRowZero := aa0.row >= 0 && aa1.row >= 0 && aa2.row >= 0 && aa3.row >= 0
	aColZero := aa0.col >= 0 && aa1.col >= 0 && aa2.col >= 0 && aa3.col >= 0

	if aRowZero && aColZero {
		inBounds = true
	} else {
		inBounds = false
		return aa0, aa1, aa2, aa3, inBounds
	}

	// test for UpperRanges
	urc0 := GetUpperInBounds(aa0) //returns bool - true is inBounds
	urc1 := GetUpperInBounds(aa1)
	urc2 := GetUpperInBounds(aa2)
	urc3 := GetUpperInBounds(aa3)

	if urc0 && urc1 && urc2 && urc3 {
		inBounds = true
	} else {
		inBounds = false
	}

	return aa0, aa1, aa2, aa3, inBounds
}

func GetUpperInBounds(aa arr) bool {
	mr := len(MTX)

	mc := len(MTX[mr-1])
	ar := aa.row < mr
	ac := aa.col < mc

	if ar && ac {
		return true
	} else {
		return false
	}

}
