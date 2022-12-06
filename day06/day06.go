package main

import (
	"bytes"
	"fmt"
	"os"
)

const InputFile = "day06/input.txt"

func main() {
	data, err := os.ReadFile(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not read file %v \n", err))
	}

	var i int
	for i = 4; i < len(data); i++ {
		if data[i] != data[i-1] &&
			data[i-1] != data[i-2] &&
			data[i-2] != data[i-3] &&
			data[i-3] != data[i] &&
			data[i] != data[i-2] &&
			data[i-1] != data[i-3] {
			break
		}
	}

	fmt.Printf("(Part 1) Start-of-package marker occurs after %d chars \n", i+1)

	needed := 14
	unique := 1

	for i = 1; i < len(data); i++ {
		begin := i - unique
		unique++
		if idx := bytes.LastIndexByte(data[begin:i], data[i]); idx != -1 {
			unique -= idx + 1
			if unique < 1 {
				unique = 1
			}
		}
		fmt.Printf("%4d %c, %c (%d) \n", i, data[begin:i], data[i], unique)
		if unique == needed {
			break
		}

	}

	fmt.Printf("(Part 2) Start-of-Message marker occurs after %d chars \n", i+1)

}
