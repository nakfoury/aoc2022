package main_test

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"testing"

	. "github.com/nakfoury/aoc2022/solutions"
)

func BenchmarkDay(b *testing.B) {
	for d, f := range DAY_TO_FUNC {
		file, _ := os.Open(fmt.Sprintf("input/day%d.txt", d))

		b.Run(fmt.Sprintf("Day %d Part 1", d), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				inp := bufio.NewScanner(file)
				f(false, inp)
				file.Seek(0, io.SeekStart)
			}
		})

		b.Run(fmt.Sprintf("Day %d Part 2", d), func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				inp := bufio.NewScanner(file)
				f(true, inp)
				file.Seek(0, io.SeekStart)
			}
		})
	}
}
