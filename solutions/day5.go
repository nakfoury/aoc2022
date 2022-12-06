package solutions

import (
	"bufio"
	"regexp"
	"strconv"
)

func loadLine(line string, stacks [][]byte) [][]byte {
	for i := 1; i < len(line); i += 4 {
		if line[i] >= 48 && line[i] <= 57 {
			return stacks
		}
		if line[i] != 32 {
			stacks[(i-1)/4] = append(stacks[(i-1)/4], line[i])
		}
	}
	return stacks
}

func reverseStack(stack []byte) []byte {
	n := len(stack) - 1
	for i := 0; i < (n+1)/2; i++ {
		stack[i], stack[n-i] = stack[n-i], stack[i]
	}
	return stack
}

func reverseStacks(stacks [][]byte) [][]byte {
	for i := 0; i < len(stacks); i++ {
		stacks[i] = reverseStack(stacks[i])
	}
	return stacks
}

func moveCrate(n int, from int, to int, stacks [][]byte) [][]byte {
	tmp := make([]byte, n)

	for i := 0; i < n; i++ {
		top := len(stacks[from]) - 1
		tmp[i] = stacks[from][top]
		stacks[from] = stacks[from][:top]
	}

	stacks[to] = append(stacks[to], tmp...)

	return stacks
}

func moveCrate9001(n int, from int, to int, stacks [][]byte) [][]byte {
	top := len(stacks[from]) - n
	var tmp []byte = stacks[from][top:]
	stacks[from] = stacks[from][:top]

	stacks[to] = append(stacks[to], tmp...)

	return stacks
}

func Day5a(inp *bufio.Scanner) string {
	re := regexp.MustCompile(`move ([0-9]+) from ([0-9]) to ([0-9])`)

	var result []byte
	var stacks [][]byte

	loaded := false
	for inp.Scan() {
		line := inp.Text()

		if stacks == nil {
			stacks = make([][]byte, (len(line)+1)/4)
			result = make([]byte, (len(line)+1)/4)
		}

		if line == "" {
			loaded = true
			stacks = reverseStacks(stacks)
			continue
		}

		if !loaded {
			stacks = loadLine(line, stacks)
		} else {
			m := re.FindStringSubmatch(line)
			n, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[2])
			to, _ := strconv.Atoi(m[3])
			moveCrate(n, from-1, to-1, stacks)
		}
	}

	for i, stack := range stacks {
		result[i] = stack[len(stack)-1]
	}

	return string(result)
}

func Day5b(inp *bufio.Scanner) string {
	re := regexp.MustCompile(`move ([0-9]+) from ([0-9]) to ([0-9])`)

	var result []byte
	var stacks [][]byte

	loaded := false
	for inp.Scan() {
		line := inp.Text()

		if stacks == nil {
			stacks = make([][]byte, (len(line)+1)/4)
			result = make([]byte, (len(line)+1)/4)
		}

		if line == "" {
			loaded = true
			stacks = reverseStacks(stacks)
			continue
		}

		if !loaded {
			stacks = loadLine(line, stacks)
		} else {
			m := re.FindStringSubmatch(line)
			n, _ := strconv.Atoi(m[1])
			from, _ := strconv.Atoi(m[2])
			to, _ := strconv.Atoi(m[3])
			moveCrate9001(n, from-1, to-1, stacks)
		}
	}

	for i, stack := range stacks {
		result[i] = stack[len(stack)-1]
	}

	return string(result)
}
