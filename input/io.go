package input

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func GetInput(day int) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("input/day%d.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

// func MakeTestInput(inp string) *bufio.Scanner {
// 	var s scanner.Scanner
// 	s.Init(strings.NewReader(inp))
// 	return s
// }
