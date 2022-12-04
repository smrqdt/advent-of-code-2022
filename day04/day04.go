package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day04/input.txt"

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

	contained := 0
	overlapCount := 0

	for scanner.Scan() {
		text := scanner.Text()

		bounds := parseLine(text)
		overlap := countOverlap(bounds)
		// could have saved some time if I read the task more carefully…
		// overlapCount += overlap
		if overlap > 0 {
			overlapCount++
		}

		if fullyContained(bounds) {
			fmt.Printf("Fully contained: %v %v (Overlap: %d)\n", text, bounds, overlap)
			contained++
		} else {
			fmt.Printf("  Not contained: %v %v (Overlap: %d)\n", text, bounds, overlap)
		}

	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Fully contained: %d\n", contained)
	fmt.Printf("(Part 2) Sum of overlap: %d\n", overlapCount)
}

func parseLine(line string) []int {
	boundsStr := strings.FieldsFunc(line, func(r rune) bool { return r == '-' || r == ',' })
	bounds := []int{}
	for _, b := range boundsStr {
		bInt, err := strconv.Atoi(b)
		if err != nil {
			panic(err)
		}
		bounds = append(bounds, bInt)
	}
	return bounds
}

func fullyContained(bounds []int) bool {
	aMin := bounds[0]
	aMax := bounds[1]
	bMin := bounds[2]
	bMax := bounds[3]
	return (aMin >= bMin && aMax <= bMax) || (bMin >= aMin && bMax <= aMax)
}

func countOverlap(bounds []int) int {
	aMin := bounds[0]
	aMax := bounds[1]
	bMin := bounds[2]
	bMax := bounds[3]

	if aMin > bMin {
		_, bMin = bMin, aMin
		aMax, bMax = bMax, aMax
	}

	if bMin <= aMax {
		return min(aMax, bMax) - bMin + 1
	}
	return 0
}

func min(a int, b int) int {
	if a > b {
		return b
	}
	return a
}

// func max(a int, b int) int {
// 	if a < b {
// 		return b
// 	}
// 	return a
// }
