package main

import (
	"fmt"

	"github.com/nakfoury/aoc2022/input"
	. "github.com/nakfoury/aoc2022/solutions"
)

func main() {
	for i, f := range DAY_TO_FUNC {
		fmt.Println("Day ", i, "Part 1: ", f(false, input.GetInput(i)))
		fmt.Println("Day ", i, "Part 2: ", f(true, input.GetInput(i)))
	}
}
