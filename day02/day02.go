package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	LOSE_POINTS = 0
	WIN_POINTS  = 6
	DRAW_POINTS = 3
)

type Figure struct {
	my     string
	their  string
	points int
}

var (
	Rock     = Figure{"X", "A", 1}
	Paper    = Figure{"Y", "B", 2}
	Scissors = Figure{"Z", "C", 3}
)

const InputFile = "day02/input.txt"

func main() {
	part1()
	part2()
}

func part1() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	var points int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, " ")
		other := &letters[0]
		own := &letters[1]
		switch *own {
		case Rock.my:
			points += Rock.points
			switch *other {
			case Rock.their:
				points += DRAW_POINTS
			case Scissors.their:
				points += WIN_POINTS
			}
		case Paper.my:
			points += Paper.points
			switch *other {
			case Paper.their:
				points += DRAW_POINTS
			case Rock.their:
				points += WIN_POINTS
			}
		case Scissors.my:
			points += Scissors.points
			switch *other {
			case Scissors.their:
				points += DRAW_POINTS
			case Paper.their:
				points += WIN_POINTS
			}
		}
	}
	fmt.Printf("(Part 1) Points: %d \n", points)
}

func part2() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	const (
		LOSE = "X"
		DRAW = "Y"
		WIN  = "Z"
	)

	var points int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		letters := strings.Split(line, " ")
		other := &letters[0]
		result := &letters[1]
		switch *other {
		case Rock.their:
			switch *result {
			case LOSE:
				points += Scissors.points
			case DRAW:
				points += Rock.points + DRAW_POINTS
			case WIN:
				points += Paper.points + WIN_POINTS
			}
		case Paper.their:
			switch *result {
			case LOSE:
				points += Rock.points
			case DRAW:
				points += Paper.points + DRAW_POINTS
			case WIN:
				points += Scissors.points + WIN_POINTS
			}
		case Scissors.their:
			switch *result {
			case LOSE:
				points += Paper.points
			case DRAW:
				points += Scissors.points + DRAW_POINTS
			case WIN:
				points += Rock.points + WIN_POINTS
			}
		}
	}
	fmt.Printf("(Part 2) Points: %d \n", points)
}
