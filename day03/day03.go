package main

import (
	"bufio"
	"fmt"
	"os"
)

const InputFile = "day03/input.txt"

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

	scanner := bufio.NewScanner(f)
	prioSum := 0
	for scanner.Scan() {
		text := scanner.Text()
		duplicates := findDuplicateItems(text)
		priorities := 0
		for _, item := range duplicates {
			prio, err := getPriority(item)
			if err != nil {
				panic(err)
			}
			priorities += prio
		}
		// fmt.Printf("%v: %q %v (%d)\n", text, duplicates, duplicates, priorities)
		prioSum += priorities
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of Priorities of items in both compartments: %d \n", prioSum)
}

func part2() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	prioSum := 0
	groupSize := 3
	for scanner.Scan() {
		rucksacks := make([]string, groupSize)
		rucksacks[0] = scanner.Text()
		for i := range rucksacks[1:] {
			if !scanner.Scan() {
				fmt.Fprintln(os.Stderr, "Could not fill last group")
				return
			}
			rucksacks[1:][i] = scanner.Text()
		}
		badge, err := findBadgeItem(rucksacks...)
		if err != nil {
			fmt.Printf("Rucksacks %v\n", rucksacks)
			panic(err)
		}
		prio, err := getPriority(badge)
		if err != nil {
			panic(err)
		}
		// fmt.Printf("Rucksacks %v contain badge %q (%d)\n", rucksacks, badge, prio)
		prioSum += prio
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 2) Sum of Priorities of badges: %d \n", prioSum)
}

func findDuplicateItems(content string) (duplicates []rune) {
	set := make(map[rune]bool)
	half := len(content) / 2
	for _, item := range content[:half] {
		set[item] = false
	}
	for _, item := range content[half:] {
		if duplicate, exists := set[item]; exists && !duplicate {
			duplicates = append(duplicates, item)
			set[item] = true
		}
	}
	return duplicates
}

func findBadgeItem(rucksacks ...string) (badge rune, err error) {
	itemSet := make(map[rune]int)
	for _, rucksack := range rucksacks {
		// make sure each item type is only counted once per rucksack
		rucksackSet := make(map[rune]bool)
		for _, item := range rucksack {
			if _, contains := rucksackSet[item]; !contains {
				itemSet[item]++
				rucksackSet[item] = true
				return item, nil
				// if itemSet[item] == len(rucksacks) {
				// 	if badge != 0 {
				// 		return badge, fmt.Errorf("Multiple Badges found! (%q, %q, %d)", badge, item, itemSet[item])
				// 	}
				// 	badge = item
				// }
			}
		}
	}
	if badge != 0 {
		return
	}
	return 0, fmt.Errorf("No item found %d time", len(rucksacks))
}

func getPriority(item rune) (int, error) {
	if item >= 'a' && item <= 'z' {
		return int(item - 'a' + 1), nil
	}
	if item >= 'A' && item <= 'Z' {
		return int(item - 'A' + 27), nil
	}
	return 0, fmt.Errorf("Could not get priority of '%q'", item)
}
