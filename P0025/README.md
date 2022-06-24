# given 4 points on a graph, is it a rectangle?

	//1 Rock,
	//2 Paper,
	//3 Scissors,
	//4 Lizard,
	//5 Spock

	//T1: "Scissors cut Paper" - 3,2
	//T2: "Paper covers Rock"  - 2,1
	//T3: "Rock crushes Lizard" - 1,4
	//T4: "Lizard poisons Spock" - 4,5
	//T5": Spock smashes Scissors" - 5,3
	//T6: "Scissors decapitates Lizard" -3,4
	//T7: "Lizard eats Paper" - 4,2
	//T8: "Paper disproves Spock" - 2,5
	//T9: "Spock vaporizes Rock" - 5,1
	//T10: "Rock crushes Scissors" - 1,3

//=================
	-1 // rock
- 2, T2 - "paper covers Rock"
- 5, T9 - "Spock vaporizes Rock"

	-2 // paper
		- 3, T1: "Scissors cut Paper"
		- 4, T7: "Lizard eats Paper"


	-3 //Scisors
		- 5, T5": Spock smashes Scissors" - 5,3
		- 1, T10: "Rock crushes Scissors" - 1,3

	-4 //Lizard
		-1, T3: "Rock crushes Lizard" - 1,4
		-3, T6: "Scissors decapitates Lizard" -3,4


	-5 //Spock
		-4, T4: "Lizard poisons Spock" - 4,5
		-2, T8: "Paper disproves Spock" - 2,5


