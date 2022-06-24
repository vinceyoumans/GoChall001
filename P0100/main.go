package main

import (
	"fmt"
	"strconv"
)

const (
	Rock     int = 1
	Paper    int = 2
	Scissors int = 3
	Lizard   int = 4
	Spock    int = 5
)

const (
	Tie     int = 0
	Player1 int = 1
	Player2 int = 2
)

type LWCombo001 struct {
	X1          int    // player 1
	X2          int    // player 2
	Winner      int    //
	Description string //  the string description
}

var MapLoseWinLookup3 map[string]LWCombo001 = map[string]LWCombo001{
	"3.2": LWCombo001{
		X1: Scissors,
		X2: Paper,
		//Winner:      Player1,
		Description: "Scissors cut Paper",
	},

	"2.1": LWCombo001{
		X1: Paper,
		X2: Rock,
		//Winner:      Player1,
		Description: "Paper covers Rock",
	},
	"1.4": LWCombo001{
		X1: Rock,
		X2: Lizard,
		//Winner:      Player1,
		Description: "Rock crushes Lizard",
	},
	"4.5": LWCombo001{
		X1: Lizard,
		X2: Spock,
		//Winner:      Player1,
		Description: "Lizard poisons Spock",
	},
	"5.3": LWCombo001{
		X1: Spock,
		X2: Scissors,
		//Winner:      Player1,
		Description: "Spock smashes Scissors",
	},
	"3.4": LWCombo001{
		X1: Scissors,
		X2: Lizard,
		//Winner:      Player1,
		Description: "Scissors decapitates Lizard",
	},
	"4.2": LWCombo001{
		X1: Lizard,
		X2: Paper,
		//Winner:      Player1,
		Description: "Lizard eats Paper",
	},
	"2.5": LWCombo001{
		X1: Paper,
		X2: Spock,
		//Winner:      Player1,
		Description: "Paper disproves Spock",
	},
	"5.1": LWCombo001{
		X1: Spock,
		X2: Rock,
		//Winner:      Player1,
		Description: "Spock vaporizes Rock",
	},
	"1.3": LWCombo001{
		X1: Rock,
		X2: Scissors,
		//Winner:      Player1,
		Description: "Rock crushes Scissors",
	},
}

func main() {
	fmt.Println("starting main")

	RestTest := GetResultGame(Rock, Scissors)
	fmt.Println(RestTest)
	fmt.Println(PrettyResult(RestTest))

	fmt.Println("===============")
	RestTest = GetResultGame(Scissors, Rock)
	fmt.Println(RestTest)
	fmt.Println(PrettyResult(RestTest))

	RestTest = GetResultGame(Lizard, Rock)
	fmt.Println(RestTest)
	fmt.Println(PrettyResult(RestTest))

	fmt.Println("===============")
	RestTest = GetResultGame(Rock, Lizard)
	fmt.Println(RestTest)
	fmt.Println(PrettyResult(RestTest))
}

func GetResultGame(xx1 int, xx2 int) LWCombo001 {
	// concept is that if 1v2 is a looser, then 2v1 should be the winner

	var resp LWCombo001
	// check to confirm not a tie
	if xx1 == xx2 {
		//fmt.Println("===  layer 1 wins")
		resp.X1 = xx1
		resp.X2 = xx2
		resp.Winner = Tie
		resp.Description = "It's a Tie"
		return resp
	}

	ind12 := GetRTSIndexString(xx1, xx2)
	ind21 := GetRTSIndexString(xx2, xx1)

	res12, ok12 := MapLoseWinLookup3[ind12]
	if ok12 {
		//fmt.Println("===  layer 1 wins")
		resp.X1 = xx1
		resp.X2 = xx2
		resp.Winner = Player1
		resp.Description = res12.Description
	}

	res21, ok21 := MapLoseWinLookup3[ind21]
	if ok21 {
		//fmt.Println("====  player 2 wins")
		resp.X1 = xx1
		resp.X2 = xx2
		resp.Winner = Player2
		resp.Description = res21.Description
	}
	return resp
}

func GetRTSIndexString(a, b int) string {
	//  concept is to get the Key of the map
	return strconv.Itoa(a) + "." + strconv.Itoa(b)
}

//  I added this to make the results more understandable.
func PrettyResult(test LWCombo001) string {
	var g string
	//  I added this for the PrettyPrintfunction
	MM := map[int]string{
		1: "Rock",
		2: "Paper",
		3: "Scissors",
		4: "Lizard",
		5: "Spock",
	}

	switch test.Winner {
	case Tie:
		g = "player 1 -" + MM[test.X1] + "(tie) "
		g += " ---- Player 2 -" + MM[test.X2] + "(tie) ======"
		g += "===" + test.Description
	case Player1:
		g = "player 1 -" + MM[test.X1] + "(winner) "
		g += " ---- Player 2 -" + MM[test.X2] + "(looser) ======"
		g += "===" + test.Description
	case Player2:
		g = "player 1 -" + MM[test.X1] + "(looser) "
		g += " ---- Player 2 -" + MM[test.X2] + "(winner) ======"
		g += "===" + test.Description
	}
	return g
}
