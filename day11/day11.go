package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const InputFile = "day11/input.txt"

func main() {
	part1 := MonkeyBusiness(20, 3)
	fmt.Printf("(Part 1) Monkey Business: %d \n", part1)
	part2 := MonkeyBusiness(10000, 1)
	fmt.Printf("(Part 1) Monkey Business: %d \n", part2)
}

func MonkeyBusiness(rounds, worryDivisor int) int {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var section []string
	var monkeys []Monkey
	divisorProduct := 1
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			monkey, divisor := ParseMonkey(section)
			divisorProduct *= divisor
			monkeys = append(monkeys, monkey)
			section = []string{}
			continue
		}
		section = append(section, line)
	}
	if len(section) > 0 {
		monkey, divisor := ParseMonkey(section)
		divisorProduct *= divisor
		monkeys = append(monkeys, monkey)
	}

	for round := 1; round <= rounds; round++ {
		fmt.Printf("===== Round %2d =====\n", round)
		for i := range monkeys {
			monkeys[i].ThrowItems(&monkeys, worryDivisor, divisorProduct)
		}
		for i := range monkeys {
			fmt.Println(monkeys[i])
		}
		fmt.Println()
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].InspectionCount > monkeys[j].InspectionCount
	})

	return monkeys[0].InspectionCount * monkeys[1].InspectionCount
}

type Monkey struct {
	ID              int
	Items           []int
	Op              func(int) int
	Next            func(int) int
	InspectionCount int
}

func (m *Monkey) ThrowItems(monkeys *[]Monkey, worryDivisor int, modDivisor int) {
	for _, worry := range m.Items {
		worry = m.Op(worry)
		m.InspectionCount++
		worry /= worryDivisor
		worry %= modDivisor
		next := m.Next(worry)
		(*monkeys)[next].Catch(worry)
	}
	m.Items = nil
}

func (m *Monkey) Catch(item int) {
	m.Items = append(m.Items, item)
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %d (%d Items inspected)", m.ID, m.Items, m.InspectionCount)
}

// Parser

func ParseMonkey(lines []string) (Monkey, int) {
	id := parseID(lines[0])
	items := parseStartingNumbers(lines[1])
	opFn := parseOp(lines[2])
	nextFn, divisor := parseNext(lines[3:])

	return Monkey{id, items, opFn, nextFn, 0}, divisor
}

func parseID(line string) (id int) {
	// Monkey 0:
	re := regexp.MustCompile(`Monkey (\d)`)
	id, err := strconv.Atoi(re.FindStringSubmatch(line)[1])
	if err != nil {
		panic(err)
	}
	return
}

func parseStartingNumbers(line string) (items []int) {
	//   Starting items: 79, 98
	re := regexp.MustCompile(`Starting items: ((?:\d+(?:, )?)+)`)
	itemsMatchStr := re.FindStringSubmatch(line)[1]
	for _, itemStr := range strings.Split(itemsMatchStr, ", ") {
		item, err := strconv.Atoi(itemStr)
		if err != nil {
			panic(err)
		}
		items = append(items, item)
	}
	return
}

func parseOp(line string) (opFn func(int) int) {
	// Operation: new = old * 19
	re := regexp.MustCompile(`Operation: new = old (\+|\*) (\d+|old)`)
	opMatches := re.FindStringSubmatch(line)

	if opMatches[2] == "old" {
		switch opMatches[1] {
		case "+":
			return func(old int) int {
				return old + old
			}
		case "*":
			return func(old int) int {
				return old * old
			}
		}
	} else {
		number, err := strconv.Atoi(opMatches[2])
		if err != nil {
			panic(err)
		}
		switch opMatches[1] {
		case "+":
			return func(old int) int {
				return old + number
			}
		case "*":
			return func(old int) int {
				return old * number
			}
		}
	}
	panic("Could not parse Operation")
}

func parseNext(lines []string) (func(int) int, int) {
	// Test: divisible by 23
	re := regexp.MustCompile(`Test: divisible by (\d+)`)
	match := re.FindStringSubmatch(lines[0])
	divisor, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	//   If true: throw to monkey 2
	re = regexp.MustCompile(`If (true|false): throw to monkey (\d+)`)
	match = re.FindStringSubmatch(lines[1])
	if match[1] != "true" {
		panic("true match is not true monkey")
	}
	trueMonkey, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	//   If false: throw to monkey 3
	match = re.FindStringSubmatch(lines[2])
	if match[1] != "false" {
		panic("false match is not false monkey")
	}
	falseMonkey, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	return func(worry int) (monkey int) {
		if worry%divisor == 0 {
			return trueMonkey
		}
		return falseMonkey
	}, divisor

}
