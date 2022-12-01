package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	const InputFile = "day01/input.txt"

	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	elves := make([]int, 0)
	var elf int

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			elves = append(elves, elf)
			elf = 0
			continue
		}
		cal, err := strconv.Atoi(text)
		if err != nil {
			panic(fmt.Sprintf("Could not parse int on line: %v \n", text))
		}
		elf += cal
	}

	sort.Sort(sort.Reverse(sort.IntSlice(elves)))
	fmt.Printf("The elf carring most Calories carries %d cal. \n", elves[0])
	fmt.Printf("The top three elfes carry %d cal. \n", elves[0]+elves[1]+elves[2])
}
