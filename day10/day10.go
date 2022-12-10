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
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Part 1
	var signalStrengths int
	watchCycles := []int{20, 60, 100, 140, 180, 220}
	watchFn := func(cpu *CPU) {
		if slices.Contains(watchCycles, cpu.Time) {
			// fmt.Printf("Cycle: %d, Register: %d \n", cpu.Time, cpu.Register)
			signalStrengths += cpu.Time * cpu.Register
		}
	}

	// Part 2
	crt := make(CRT, 40*6)

	cpu := NewCPU([]func(*CPU){watchFn, crt.Draw})

	for scanner.Scan() {
		line := scanner.Text()
		cpu.exec(line)
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Printf("(Part 1) Sum of Signal Strengths: %d \n", signalStrengths)
	fmt.Println("(Part 2) CRT Image:")
	crt.Print()
}

type CPU struct {
	Register int
	Time     int
	WatchFn  []func(*CPU)
}

func NewCPU(watchFn []func(*CPU)) *CPU {
	return &CPU{1, 1, watchFn}
}

func (cpu *CPU) cycle() {
	// fmt.Printf("Cycle %d \n", cpu.Time)
	for _, fn := range cpu.WatchFn {
		fn(cpu)
	}
	cpu.Time++
}

func (cpu *CPU) exec(line string) {
	// fmt.Printf("Exec %s \n", line)
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
	// fmt.Printf("Finish %s (Register value %d) \n", line, cpu.Register)
}

func (cpu *CPU) addx(val int) {
	cpu.cycle()
	cpu.cycle()
	cpu.Register += val
}

func (cpu *CPU) noop() {
	cpu.cycle()
}

type CRT []byte

func (crt CRT) Draw(cpu *CPU) { // Direct Register Access!!11
	if col := (cpu.Time - 1) % 40; col >= cpu.Register-1 && col <= cpu.Register+1 {
		crt[cpu.Time-1] = '#'
	} else {
		crt[cpu.Time-1] = ' '
	}
	// fmt.Printf("Drawing Pos %d: %q \n", cpu.Time%40, crt[cpu.Time-1])
}

func (crt CRT) Print() {
	for i := 0; i < len(crt); i += 40 {
		fmt.Println(string(crt[i : i+40]))
	}
}
