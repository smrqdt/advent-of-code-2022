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
	part1()
}

func part1() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	rope := NewRope()

	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		dir := fields[0]
		steps, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		rope.MoveHead(dir, steps)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Numer of visited Positions: %d \n", rope.CountPositions())
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
	Head   Position
	Tail   Position
	PosMap map[Position]bool
}

func NewRope() *Rope {
	r := &Rope{Position{}, Position{}, make(map[Position]bool)}
	r.MapTailPosition()
	return r
}

func (r *Rope) MoveHead(direction string, steps int) {
	for i := 0; i < steps; i++ {
		r.Head.Move(direction)
		r.TailFollow()
	}
}

func (r *Rope) TailFollow() {
	xDelta := r.Head.X - r.Tail.X
	yDelta := r.Head.Y - r.Tail.Y
	if Abs(xDelta) <= 1 && Abs(yDelta) <= 1 {
		return
	}
	r.Tail.Move(FollowDirection(xDelta, yDelta))
	r.MapTailPosition()
}

func (r *Rope) MapTailPosition() {
	r.PosMap[r.Tail] = true
}

func (r *Rope) CountPositions() int {
	return len(r.PosMap)
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
