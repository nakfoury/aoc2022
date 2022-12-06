package solutions

import (
	"bufio"
)

func byteToPriority(elt byte) int {
	if elt >= 65 && elt <= 90 {
		return int(elt - 38)
	} else if elt >= 97 && elt <= 122 {
		return int(elt - 96)
	} else {
		return 0
	}
}

func isInBag(badge byte, bag string) bool {
	for i := range bag {
		if badge == bag[i] {
			return true
		}
	}
	return false
}

func Day3a(inp *bufio.Scanner) int {
	var result int = 0

	for inp.Scan() {
		rucksack := inp.Text()
		size := len(rucksack) / 2
		var found bool = false

		for i := 0; i < size; i++ {
			for j := size; j < size*2; j++ {
				if rucksack[i] == rucksack[j] {
					result += byteToPriority(rucksack[i])
					found = true
				}
				if found {
					break
				}
			}
			if found {
				break
			}
		}

	}

	return result
}

func Day3b(inp *bufio.Scanner) int {
	result := 0

	more := true
	group := make([]string, 3)

	for more {
		for i := 0; i < 3; i++ {
			more = inp.Scan()
			group[i] = inp.Text()
		}
		for i := range group[0] {
			badge := group[0][i]
			if isInBag(badge, group[1]) && isInBag(badge, group[2]) {
				result += byteToPriority(badge)
				break
			}

		}
	}

	return result
}
