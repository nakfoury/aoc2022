package solutions

import (
	"sort"
	"strconv"

	"github.com/nakfoury/aoc2022/input"
)

func Day1(part2 bool) int {
	inp := input.GetInput(1)

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
