package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

const InputFile = "day11/demo.txt"

var zero = big.NewInt(0)

func main() {
	part1 := monkeyBusiness(big.NewInt(3), 20)
	fmt.Printf("(Part 1) Monkey Business: %d \n", part1)
	fmt.Println()
	part2 := monkeyBusiness(nil, 10000)
	fmt.Printf("(Part 2) Monkey Business: %d \n", part2)

}

func monkeyBusiness(worryDivisor *big.Int, rounds int) int {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var section []string
	var monkeys []Monkey
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			monkey := ParseMonkey(section)
			fmt.Println(monkey)
			monkeys = append(monkeys, monkey)
			section = []string{}
			continue
		}
		section = append(section, line)
	}
	if len(section) > 0 {
		monkeys = append(monkeys, ParseMonkey(section))
	}

	for round := 1; round <= rounds; round++ {
		fmt.Printf("===== Round %2d =====\n", round)
		for i := range monkeys {
			monkeys[i].ThrowItems(&monkeys, worryDivisor)
		}
		for i := range monkeys {
			fmt.Println(monkeys[i])
		}
	}
	fmt.Println()

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
	Items           []*big.Int
	Op              func(*big.Int) *big.Int
	Next            func(*big.Int) int
	InspectionCount int
}

func (m *Monkey) ThrowItems(monkeys *[]Monkey, worryDivisor *big.Int) {
	for _, item := range m.Items {
		m.Op(item)
		m.InspectionCount++
		if worryDivisor != nil {
			item.Div(item, worryDivisor)
		}
		next := m.Next(item)
		(*monkeys)[next].Catch(item)
	}
	m.Items = nil
}

func (m *Monkey) Catch(item *big.Int) {
	m.Items = append(m.Items, item)
}

func (m Monkey) String() string {
	return fmt.Sprintf("Monkey %d: %d Items inspected", m.ID, m.InspectionCount)
}

// Parser

func ParseMonkey(lines []string) Monkey {
	id := parseID(lines[0])
	items := parseStartingNumbers(lines[1])
	opFn := parseOp(lines[2])
	nextFn := parseNext(lines[3:])

	return Monkey{id, items, opFn, nextFn, 0}
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

func parseStartingNumbers(line string) (items []*big.Int) {
	//   Starting items: 79, 98
	re := regexp.MustCompile(`Starting items: ((?:\d+(?:, )?)+)`)
	itemsMatchStr := re.FindStringSubmatch(line)[1]
	for _, itemStr := range strings.Split(itemsMatchStr, ", ") {
		item, err := strconv.ParseInt(itemStr, 10, 64)
		if err != nil {
			panic(err)
		}
		items = append(items, big.NewInt(item))
	}
	return
}

func parseOp(line string) (opFn func(*big.Int) *big.Int) {
	// Operation: new = old * 19
	re := regexp.MustCompile(`Operation: new = old (\+|\*) (\d+|old)`)
	opMatches := re.FindStringSubmatch(line)

	if opMatches[2] == "old" {
		switch opMatches[1] {
		case "+":
			return func(old *big.Int) *big.Int {
				return old.Add(old, old)
			}
		case "*":
			return func(old *big.Int) *big.Int {
				return old.Mul(old, old)
			}
		}
	} else {
		numberUint, err := strconv.ParseInt(opMatches[2], 10, 64)
		number := big.NewInt(numberUint)
		if err != nil {
			panic(err)
		}
		switch opMatches[1] {
		case "+":
			return func(old *big.Int) *big.Int {
				return old.Add(old, number)
			}
		case "*":
			return func(old *big.Int) *big.Int {
				return old.Mul(old, number)
			}
		}
	}
	panic("Could not parse Operation")
}

func parseNext(lines []string) func(*big.Int) int {
	// Test: divisible by 23
	re := regexp.MustCompile(`Test: divisible by (\d+)`)
	match := re.FindStringSubmatch(lines[0])
	divisorUint, err := strconv.ParseInt(match[1], 10, 64)
	divisor := big.NewInt(divisorUint)
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
	return func(worry *big.Int) (monkey int) {
		if mod := big.NewInt(0).Mod(worry, divisor); mod.Cmp(zero) == 0 {
			return trueMonkey
		}
		return falseMonkey
	}

}
