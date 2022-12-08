package solutions

import (
	"bufio"
)

func init() {
	DAY_TO_FUNC[2] = Day2
}

var RPS = map[byte]int{
	'A': 1,
	'B': 2,
	'C': 3,
	'X': 1,
	'Y': 2,
	'Z': 3,
}

var STRAT = map[byte]int{
	'X': 0,
	'Y': 3,
	'Z': 6,
}

func Day2a(inp *bufio.Scanner) int {
	var totalScore = 0
	for inp.Scan() {
		line := inp.Text()

		op := RPS[line[0]]
		me := RPS[line[2]]

		var outcome int
		if op == me {
			outcome = 3
		} else if me == op+1 || (me == 1 && op == 3) {
			outcome = 6
		} else {
			outcome = 0
		}

		totalScore += me
		totalScore += outcome
	}
	return totalScore
}

func Day2b(inp *bufio.Scanner) int {
	var totalScore = 0

	for inp.Scan() {
		line := inp.Text()

		op := RPS[line[0]]
		strat := STRAT[line[2]]
		me := (op+RPS[line[2]])%3 + 1

		totalScore += me
		totalScore += strat
	}

	return totalScore
}

func Day2(part2 bool, inp *bufio.Scanner) any {
	if !part2 {
		return Day2a(inp)
	} else {
		return Day2a(inp)
	}
}
