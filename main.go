package main

import (
	"fmt"

	"github.com/nakfoury/aoc2022/input"
	. "github.com/nakfoury/aoc2022/solutions"
)

func main() {
	fmt.Println("day1 part1:", Day1(false))
	fmt.Println("day1 part2:", Day1(true))
	fmt.Println("day2 part1:", Day2a(input.GetInput(2)))
	fmt.Println("day2 part2:", Day2b(input.GetInput(2)))
	fmt.Println("day3 part1:", Day3a(input.GetInput(3)))
	fmt.Println("day3 part2:", Day3b(input.GetInput(3)))
	fmt.Println("day4 part1:", Day4a(input.GetInput(4)))
	fmt.Println("day4 part2:", Day4b(input.GetInput(4)))
	fmt.Println("day5 part1:", Day5a(input.GetInput(5)))
	fmt.Println("day5 part2:", Day5b(input.GetInput(5)))
	fmt.Println("day6 part1:", Day6a(input.GetInput(6)))
	fmt.Println("day6 part2:", Day6b(input.GetInput(6)))
}
