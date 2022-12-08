package main

import (
	"bufio"
	"fmt"
	"os"
)

const InputFile = "day08/input.txt"

func main() {
	forest := readInput()
	part1(forest)
}

func part1(forest [][]int) {
	visible := make([][]bool, len(forest))
	for i := range visible {
		visible[i] = make([]bool, len(forest[0]))
	}
	lastRow := len(forest) - 1
	lastCol := len(forest[0]) - 1

	// forward check
	colMax := make([]int, len(forest[0]))
	for i := range colMax {
		colMax[i] = -1
	}
	for row := 0; row <= lastRow; row++ {
		rowMax := -1
		for col := 0; col <= lastCol; col++ {
			height := forest[row][col]

			if height > rowMax {
				visible[row][col] = true
				rowMax = height
			}
			if height > colMax[col] {
				visible[row][col] = true
				colMax[col] = height
			}
		}
	}

	for i := range colMax {
		colMax[i] = -1
	}
	for row := lastRow; row >= 0; row-- {
		rowMax := -1
		for col := lastCol; col >= 0; col-- {
			height := forest[row][col]

			if height > rowMax {
				visible[row][col] = true
				rowMax = height
			}
			if height > colMax[col] {
				visible[row][col] = true
				colMax[col] = height
			}
		}
	}

	visibleCount := 0
	for row := range visible {
		for col := range visible[0] {
			if visible[row][col] {
				fmt.Print("X")
				visibleCount++
			} else {
				fmt.Print("o")
			}
		}
		fmt.Println()
	}

	fmt.Printf("(Part 1) %d trees are visible from outside the grid. \n", visibleCount)
}

func readInput() [][]int {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	lines := make([][]int, 0)

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		lineInt := make([]int, 0, len(line))
		for i := range line {
			if err != nil {
				panic("Could not parse byte to int")
			}
			lineInt = append(lineInt, int(line[i]-'0'))
		}
		lines = append(lines, lineInt)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	for row := range lines {
		fmt.Printf("%d\n", lines[row])
	}

	return lines
}
