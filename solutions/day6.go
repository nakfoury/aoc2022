package solutions

import (
	"bufio"
)

func isMarker(m []byte, length int) bool {
	chars := make(map[byte]bool, length)
	for i := 0; i < len(m); i++ {
		if chars[m[i]] {
			return false
		} else {
			chars[m[i]] = true
		}
	}
	return true
}

func Day6a(inp *bufio.Scanner) int {
	var result int = 4

	inp.Scan()
	buffer := inp.Text()
	for result < len(buffer) {
		if isMarker([]byte(buffer[result-4:result]), 4) {
			break
		}
		result++
	}

	return result
}

func Day6b(inp *bufio.Scanner) int {
	var result int = 14

	inp.Scan()
	buffer := inp.Text()
	for result < len(buffer) {
		if isMarker([]byte(buffer[result-14:result]), 14) {
			break
		}
		result++
	}

	return result
}
