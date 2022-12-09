package solutions

import (
	"bufio"
	"strconv"
)

func init() {
	DAY_TO_FUNC[9] = Day9
}

type Knot struct {
	X, Y int
}

func (p *Knot) move(dir Knot) {
	p.X += dir.X
	p.Y += dir.Y
}

func (p *Knot) follow(q Knot) {
	if p.X > q.X {
		p.X -= 1
	} else if p.X < q.X {
		p.X += 1
	}

	if p.Y > q.Y {
		p.Y -= 1
	} else if p.Y < q.Y {
		p.Y += 1
	}

}

func (p *Knot) isFarFrom(q Knot) bool {
	return p.X-q.X > 1 || q.X-p.X > 1 || p.Y-q.Y > 1 || q.Y-p.Y > 1
}

type Rope struct {
	Knots [10]Knot
}

func (r *Rope) moveAll(dir Knot) {
	r.Knots[0].move(dir)
	for i := 1; i < len(r.Knots); i++ {
		if r.Knots[i].isFarFrom(r.Knots[i-1]) {
			r.Knots[i].follow(r.Knots[i-1])
		} else {
			break
		}
	}
}

func Day9(part2 bool, inp *bufio.Scanner) any {
	result := 0

	DIR := map[byte]Knot{
		'R': {1, 0},
		'L': {-1, 0},
		'U': {0, 1},
		'D': {0, -1},
	}

	visitedTail := map[Knot]bool{
		{0, 0}: true,
	}

	head := Knot{0, 0}
	tail := Knot{0, 0}

	rope := Rope{}
	rope.Knots[0] = head
	rope.Knots[len(rope.Knots)-1] = tail

	for inp.Scan() {
		line := inp.Text()
		dir := DIR[line[0]]
		count, _ := strconv.Atoi(line[2:])

		if !part2 {
			for i := 0; i < count; i++ {
				head.move(dir)
				if head.isFarFrom(tail) {
					tail.follow(head)
					visitedTail[tail] = true
				}
			}
		} else {
			for i := 0; i < count; i++ {
				rope.moveAll(dir)
				visitedTail[rope.Knots[len(rope.Knots)-1]] = true
			}
		}
	}

	for range visitedTail {
		result++
	}

	return result
}
