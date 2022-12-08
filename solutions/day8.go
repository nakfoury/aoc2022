package solutions

import (
	"bufio"
)

func init() {
	DAY_TO_FUNC[8] = Day8
}

func makeRow(line string) []int {
	result := make([]int, len(line))
	for i := 0; i < len(line); i++ {
		result[i] = int(line[i] - 48)
	}
	return result
}

func isVisible(row []int, reverse bool) bool {
	if !reverse {
		tree := row[len(row)-1]
		for i := 0; i < len(row)-1; i++ {
			if tree < row[i] {
				return false
			}
		}
	} else {
		tree := row[0]
		for i := len(row) - 1; i > 0; i-- {
			if tree < row[i] {
				return false
			}
		}
	}
	return true
}

func Day8(part2 bool, inp *bufio.Scanner) any {
	result := 0

	var forest [][]int
	for inp.Scan() {
		forest = append(forest, makeRow(inp.Text()))
	}

	height := len(forest)
	width := len(forest[0])

	visible := make([][]bool, height)
	for i := 0; i < height; i++ {
		visible[i] = make([]bool, width)
	}

	if !part2 {
		for i := 0; i < height; i++ {
			max := -1
			for j := 0; j < width; j++ {
				if forest[i][j] > max {
					visible[i][j] = true
					max = forest[i][j]
				}
			}

			max = -1
			for j := width - 1; j >= 0; j-- {
				if forest[i][j] > max {
					visible[i][j] = true
					max = forest[i][j]
				}
			}
		}

		for j := 0; j < width; j++ {
			max := -1
			for i := 0; i < height; i++ {
				if forest[i][j] > max {
					visible[i][j] = true
					max = forest[i][j]
				}
			}

			max = -1
			for i := height - 1; i >= 0; i-- {
				if forest[i][j] > max {
					visible[i][j] = true
					max = forest[i][j]
				}
			}
		}

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if visible[i][j] {
					result++
				}
			}
		}
	} else {
		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				score := getScenicScore(forest, i, j)
				if score > result {
					result = score
				}
			}
		}
	}

	return result
}

func getScenicScore(forest [][]int, i int, j int) int {
	result := 1

	tree := forest[i][j]

	subscore := 0
	for ii := i + 1; ii < len(forest); ii++ {
		subscore++
		if forest[ii][j] >= tree {
			break
		}
	}
	result *= subscore

	subscore = 0
	for ii := i - 1; ii >= 0; ii-- {
		subscore++
		if forest[ii][j] >= tree {
			break
		}
	}
	result *= subscore

	subscore = 0
	for jj := j + 1; jj < len(forest); jj++ {
		subscore++
		if forest[i][jj] >= tree {
			break
		}
	}
	result *= subscore

	subscore = 0
	for jj := j - 1; jj >= 0; jj-- {
		subscore++
		if forest[i][jj] >= tree {
			break
		}
	}
	result *= subscore

	return result
}
