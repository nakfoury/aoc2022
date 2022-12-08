package solutions

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	DAY_TO_FUNC[4] = Day4
}

func parseAssignment(ass string) (int, int) {
	re := regexp.MustCompile("([0-9]+)-([0-9]+)")
	m := re.FindStringSubmatch(ass)
	x, err := strconv.Atoi(m[1])
	if err != nil {
		panic(err)
	}
	y, err := strconv.Atoi(m[2])
	if err != nil {
		panic(err)
	}
	return x, y
}

func Day4a(inp *bufio.Scanner) int {
	var result int = 0

	for inp.Scan() {
		assignments := strings.Split(inp.Text(), ",")
		x1, y1 := parseAssignment(assignments[0])
		x2, y2 := parseAssignment(assignments[1])

		if (x1 >= x2 && y1 <= y2) || (x1 <= x2 && y1 >= y2) {
			result += 1
		}

	}

	return result
}

func Day4b(inp *bufio.Scanner) int {
	result := 0

	for inp.Scan() {
		assignments := strings.Split(inp.Text(), ",")
		x1, y1 := parseAssignment(assignments[0])
		x2, y2 := parseAssignment(assignments[1])

		if (x1 >= x2 && x1 <= y2) || (y1 >= x2 && y1 <= y2) || (x2 >= x1 && x2 <= y1) || (y2 >= x1 && y2 <= y1) {
			result += 1
		}

	}
	return result
}

func Day4(part2 bool, inp *bufio.Scanner) any {
	if !part2 {
		return Day4a(inp)
	} else {
		return Day4b(inp)
	}
}
