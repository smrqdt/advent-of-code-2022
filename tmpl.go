package main

import (
	"bufio"
	"fmt"
	"os"
)

const InputFile = "day00/input.txt"

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
	for scanner.Scan() {
		text := scanner.Text()
		fmt.Printf("%v \n", text)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) \n")
}
