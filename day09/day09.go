package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day09/input.txt"

func main() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rope2 := NewRope(2)
	rope10 := NewRope(10)

	iteration := 0
	for scanner.Scan() {
		iteration++
		line := scanner.Text()
		fields := strings.Fields(line)
		dir := fields[0]
		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		rope2.MoveHead(dir, steps)
		rope10.MoveHead(dir, steps)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Numer of visited positions with 2 knots: %d \n", rope2.CountPositions())
	fmt.Printf("(Part 2) Numer of visited positions with 10 knots: %d \n", rope10.CountPositions())
}

type Position struct {
	X, Y int
}

func (pos *Position) Move(direction string) {
	// fmt.Printf("Move %p %q \n", pos, direction)
	switch direction {
	case "U": // up
		pos.Y++
	case "R": // right
		pos.X++
	case "D": // down
		pos.Y--
	case "L": // left
		pos.X--
	case "UR": // up right
		pos.Y++
		pos.X++
	case "DR": // down right
		pos.Y--
		pos.X++
	case "DL": // down left
		pos.Y--
		pos.X--
	case "UL": // up left
		pos.Y++
		pos.X--
	case "":
	default:
		panic("Could not parse direction string")
	}
}

type Rope struct {
	Knots  []Position
	PosMap map[Position]bool
}

func NewRope(knots int) *Rope {
	r := &Rope{make([]Position, knots), make(map[Position]bool)}
	r.MapTailPosition()
	return r
}

func (r *Rope) MoveHead(direction string, steps int) {
	// fmt.Printf("== %s %d ==\n", direction, steps)
	for i := 0; i < steps; i++ {
		r.Knots[0].Move(direction)
		for knot := 1; knot < len(r.Knots); knot++ {
			r.Follow(knot)
		}
		r.MapTailPosition()
	}
	// r.Print(80, 50)
	// fmt.Println()
}

func (r *Rope) Follow(knot int) {
	xDelta := r.Knots[knot-1].X - r.Knots[knot].X
	yDelta := r.Knots[knot-1].Y - r.Knots[knot].Y
	if Abs(xDelta) <= 1 && Abs(yDelta) <= 1 {
		return
	}
	r.Knots[knot].Move(FollowDirection(xDelta, yDelta))
}

func (r *Rope) MapTailPosition() {
	r.PosMap[r.Knots[len(r.Knots)-1]] = true
}

func (r *Rope) CountPositions() int {
	return len(r.PosMap)
}

func (r *Rope) Print(width, height int) {
	var xMax, yMax int
	for i := len(r.Knots) - 1; i >= 0; i-- {
		if max := Abs(r.Knots[i].X); max > xMax {
			xMax = max
		}
		if max := r.Knots[i].Y; max > yMax {
			yMax = max
		}
	}
	if 2*yMax > height {
		height = 2 * yMax
	}
	if 2*xMax > width {
		width = 2 * xMax
	}
	matrix := make([][]int, height)
	for i := range matrix {
		matrix[i] = make([]int, width)
	}
	for i := len(r.Knots) - 1; i >= 0; i-- {
		knot := &r.Knots[i]
		if i == 0 {
			i = -1
		}
		matrix[knot.Y+height/2][knot.X+width/2] = i
	}

	for y := len(matrix) - 1; y >= 0; y-- {
		for x := 0; x < len(matrix[y]); x++ {
			knot := strconv.Itoa(matrix[y][x])
			switch knot {
			case "-1":
				knot = "H"
			case "0":
				knot = "."
			case strconv.Itoa(len(r.Knots) - 1):
				knot = "T"
			}
			if x == width/2 && y == height/2 && knot == "." {
				knot = "s"
			}
			fmt.Print(knot)
		}
		fmt.Println()
	}
}

func FollowDirection(xDelta, yDelta int) string {
	var dirX, dirY string
	switch {
	case yDelta > 0:
		dirY = "U"
	case yDelta < 0:
		dirY = "D"
	}
	switch {
	case xDelta > 0:
		dirX = "R"
	case xDelta < 0:
		dirX = "L"
	}
	return string(dirY + dirX)
}

// Helper

func Abs(n int) int {
	y := n >> (strconv.IntSize - 1)
	return (n ^ y) - y
}
