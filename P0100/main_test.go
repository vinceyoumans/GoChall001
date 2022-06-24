package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

type tt struct {
	input    [2]int
	expected LWCombo001
}

//func TestGetResultGame(Scissors, Rock)
func TestGetResultGame(t *testing.T) {
	ts := []tt{
		//0
		{
			input:    [2]int{Scissors, Scissors},
			expected: LWCombo001{3, 3, 0, "It's a Tie"},
		},
		//1
		{
			input:    [2]int{1, 3},
			expected: LWCombo001{1, 3, 1, "Rock crushes Scissors"},
		},
		//2
		{
			input:    [2]int{Rock, Scissors},
			expected: LWCombo001{1, 3, 1, "Rock crushes Scissors"},
		},
		//3
		{
			input:    [2]int{5, 4},
			expected: LWCombo001{X1: 5, X2: 4, Winner: 2, Description: "Lizard poisons Spock"},
		},
		{
			input:    [2]int{Spock, Scissors},
			expected: LWCombo001{X1: 5, X2: 3, Winner: 1, Description: "Spock smashes Scissors"},
		},
		{
			input:    [2]int{Scissors, Spock},
			expected: LWCombo001{X1: 3, X2: 5, Winner: 2, Description: "Spock smashes Scissors"},
		},
	}

	for x, xx := range ts {
		fmt.Println("===   doing ", x)
		fmt.Println(" ===  ", xx)
		RestTest := GetResultGame(xx.input[0], xx.input[1])
		assert.Equal(t, xx.expected, RestTest)
	}
}
