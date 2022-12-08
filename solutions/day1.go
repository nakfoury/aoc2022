package solutions

import (
	"bufio"
	"sort"
	"strconv"
)

func init() {
	DAY_TO_FUNC[1] = Day1
}

func Day1(part2 bool, inp *bufio.Scanner) any {
	elves := make(map[int]int)
	curElf := 0
	elves[curElf] = 0

	for inp.Scan() {
		cal, err := strconv.Atoi(inp.Text())
		if err != nil || cal == 0 {
			curElf += 1
			elves[curElf] = 0
			continue
		} else {
			elves[curElf] += cal
		}
	}

	maxElf := 0
	var elfSlice []int

	for _, v := range elves {
		if v > maxElf {
			maxElf = v
		}
		elfSlice = append(elfSlice, v)
	}

	if !part2 {
		return maxElf
	} else {
		sort.Slice(elfSlice, func(i int, j int) bool {
			return elfSlice[i] > elfSlice[j]
		})
		return elfSlice[0] + elfSlice[1] + elfSlice[2]
	}

}
