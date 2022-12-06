package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

const InputFile = "day05/input.txt"

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
	header := make([]string, 0)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		header = append(header, text)
	}
	stacksOneForOne := parseStacks(header[:len(header)-1]) // part 1
	stacksAllAtOnce := make([][]byte, len(*stacksOneForOne))
	copy(stacksAllAtOnce, *stacksOneForOne) // part2

	printStacks(stacksOneForOne)

	for scanner.Scan() {
		n, from, to := parseMove(scanner.Text())
		moveOneForOne(stacksOneForOne, n, from, to)
		moveAll(&stacksAllAtOnce, n, from, to)
		printStacks(stacksOneForOne)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	topBoxesOneForOne := make([]byte, len(*stacksOneForOne))
	for i := range *stacksOneForOne {
		topBoxesOneForOne = append(topBoxesOneForOne, (*stacksOneForOne)[i][len((*stacksOneForOne)[i])-1])
	}

	topBoxesAllAtOne := make([]byte, len(stacksAllAtOnce))
	for i := range stacksAllAtOnce {
		topBoxesAllAtOne = append(topBoxesAllAtOne, stacksAllAtOnce[i][len(stacksAllAtOnce[i])-1])
	}

	fmt.Println("Stacks sorted one by one")
	printStacks(stacksOneForOne)

	fmt.Printf("(Part 1) Top Boxes after moving them one for one: %v \n", string(topBoxesOneForOne))

	fmt.Println("Stacks sorted all at once")
	printStacks(&stacksAllAtOnce)
	fmt.Printf("(Part 1) Top Boxes after moving them in stacks: %v \n", string(topBoxesAllAtOne))

}

func parseMove(line string) (n int, from int, to int) {
	re := regexp.MustCompile(`move (\d+) from (\d+) to (\d+)`)
	match := re.FindStringSubmatch(line)
	n, _ = strconv.Atoi(match[1])
	from, _ = strconv.Atoi(match[2])
	to, _ = strconv.Atoi(match[3])
	fmt.Println(line)
	return
}

func moveAll(stacks *[][]byte, n int, from int, to int) {
	from--
	to--
	index := len((*stacks)[from]) - n
	fmt.Printf("move %c from %d to %d \n", (*stacks)[from][index:], from, to)
	(*stacks)[to] = append((*stacks)[to], (*stacks)[from][index:]...)
	(*stacks)[from] = (*stacks)[from][:index]
}

func moveOneForOne(stacks *[][]byte, n int, from int, to int) {
	from--
	to--
	for i := 1; i <= n; i++ {
		index := len((*stacks)[from]) - i
		(*stacks)[to] = append((*stacks)[to], (*stacks)[from][index])
		fmt.Printf("move %c from %d to %d \n", (*stacks)[from][index], from, to)
	}
	// (*stacks)[to] = append((*stacks)[to], (*stacks)[from][index:]...)
	index := len((*stacks)[from]) - n
	(*stacks)[from] = (*stacks)[from][:index]
}

func parseStacks(lines []string) *[][]byte {
	numOfStacks := (len(lines[0]) + 1) / 4 // its always 9, but… https://xkcd.com/974/
	stacks := make([][]byte, numOfStacks)
	for i := range stacks {
		stacks[i] = make([]byte, 0, len(lines))
	}
	for row := len(lines) - 1; row >= 0; row-- {
		for col, stack := 1, 0; col < len(lines[row]); col, stack = col+4, stack+1 {
			if lines[row][col] != ' ' {
				stacks[stack] = append(stacks[stack], lines[row][col])
			}
		}
	}
	return &stacks
}

func printStacks(stacks *[][]byte) {
	highest := 0
	for i := range *stacks {
		if size := len((*stacks)[i]); size > highest {
			highest = size
		}
	}
	for layer := highest - 1; layer >= 0; layer-- {
		for stack := range *stacks {
			if len((*stacks)[stack]) > layer {
				fmt.Printf("[%c] ", (*stacks)[stack][layer])
			} else {
				fmt.Print("    ")
			}
		}
		fmt.Println()
	}
}
