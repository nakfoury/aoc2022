package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func day1(part2 bool) int {
	inp := getInput(1)

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

func getInput(day int) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("input/day%d.txt", day))
	if err != nil {
		log.Fatal(err)
	}

	return bufio.NewScanner(file)
}

func main() {
	fmt.Println("day1 part1:", day1(false))
	fmt.Println("day1 part2:", day1(true))
}
