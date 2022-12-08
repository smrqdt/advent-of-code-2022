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
	part2(forest)
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
	printColHeader(len(visible[0]))

	for row := range visible {
		fmt.Printf("%2d  ", row)
		for col := range visible[0] {
			if visible[row][col] {
				fmt.Print("T")
				visibleCount++
			} else {
				fmt.Print(".")
			}
		}
		fmt.Printf("  %2d", row)
		fmt.Println()
	}
	printColHeader(len(visible[0]))

	fmt.Printf("(Part 1) %d trees are visible from outside the grid. \n", visibleCount)
}

func part2(forest [][]int) {
	lastRow := len(forest) - 1
	lastCol := len(forest[0]) - 1

	// scenicScores := make([][]int, len(forest))
	// for i := range scenicScores {
	// 	scenicScores[i] = make([]int, len(forest[0]))
	// }

	maxScenicScore := 0
	maxScoreRow, maxScoreCol := 0, 0
	for row := range forest {
		for col := range forest {
			height := forest[row][col]
			scenicScore := 1

			// check left
			distance := 0
			for leftCol := col - 1; leftCol >= 0; leftCol-- {
				distance++
				if forest[row][leftCol] >= height {
					break
				}
			}
			scenicScore *= distance

			// check top
			distance = 0
			for topRow := row - 1; topRow >= 0; topRow-- {
				distance++
				if forest[topRow][col] >= height {
					break
				}
			}
			scenicScore *= distance

			// check right
			distance = 0
			for rightCol := col + 1; rightCol <= lastCol; rightCol++ {
				distance++
				if forest[row][rightCol] >= height {
					break
				}
			}
			scenicScore *= distance

			// check bottom
			distance = 0
			for bottomRow := row + 1; bottomRow <= lastRow; bottomRow++ {
				distance++
				if forest[bottomRow][col] >= height {
					break
				}
			}
			scenicScore *= distance
			if maxScenicScore < scenicScore {
				maxScenicScore = scenicScore
				maxScoreRow = row
				maxScoreCol = col
			}
			// scenicScores[row][col] = scenicScore

		}
	}
	fmt.Printf("(Part 2) The best scenic score is %d at Row %d, Col %d. \n", maxScenicScore, maxScoreRow, maxScoreCol)
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

	// for row := range lines {
	// 	fmt.Printf("%d\n", lines[row])
	// }

	return lines
}

func printColHeader(cols int) {
	fmt.Print("    ")
	for col := 0; col < cols; col++ {
		fmt.Printf("%d", col/10)
	}
	fmt.Println()
	fmt.Print("    ")
	for col := 0; col < cols; col++ {
		fmt.Printf("%d", col%10)
	}
	fmt.Println()
}
