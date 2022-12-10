package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

const InputFile = "day10/input.txt"

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
	var signalStrengths int
	cpu := NewCPU([]int{20, 60, 100, 140, 180, 220}, func(cpu *CPU) {
		fmt.Printf("Cycle: %d, Register: %d \n", cpu.CurrentCycle, cpu.Register)
		signalStrengths += cpu.CurrentCycle * cpu.Register
	})

	for scanner.Scan() {
		line := scanner.Text()
		cpu.exec(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of Signal Strengths: %d \n", signalStrengths)
}

type CPU struct {
	Register     int
	CurrentCycle int
	WatchCycles  []int
	WatchFn      func(*CPU)
}

func NewCPU(watchCycles []int, watchFn func(*CPU)) *CPU {
	return &CPU{1, 1, watchCycles, watchFn}
}

func (cpu *CPU) cycle() {
	if slices.Contains(cpu.WatchCycles, cpu.CurrentCycle) {
		cpu.WatchFn(cpu)
	}
	cpu.CurrentCycle++
}

func (cpu *CPU) exec(line string) {
	fields := strings.Fields(line)
	switch fields[0] {
	case "addx":
		val, err := strconv.Atoi(fields[1])
		if err != nil {
			panic(err)
		}
		cpu.addx(val)
	case "noop":
		cpu.noop()
	}
	// fmt.Printf("exec %s, cycle %d, register %d \n", line, cpu.CurrentCycle, cpu.Register)
}

func (cpu *CPU) addx(val int) {
	cpu.cycle()
	cpu.cycle()
	cpu.Register += val
}

func (cpu *CPU) noop() {
	cpu.cycle()
}
