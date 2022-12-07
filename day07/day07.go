package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const InputFile = "day07/input.txt"

var ROOTDIR *Directory = NewDirectory("", nil)

func main() {
	readInput()
	part1()
	part2()
}

func part1() {
	answer := 0

	walkDirs(ROOTDIR, func(dir *Directory) {
		if size := dir.TotalSize(); size <= 100000 {
			answer += size
		}
	})

	fmt.Printf("(Part 1) Total Size of all Directories not larger than 100000: %d\n", answer)
}

func part2() {
	diskSize := 70000000
	unused := diskSize - ROOTDIR.TotalSize()
	neededForUpdate := 30000000
	missing := neededForUpdate - unused

	fmt.Printf("            Unused Space: %d \n", unused)
	fmt.Printf("Missing Space for Update: %d \n", missing)

	answer := diskSize
	var answerDir *Directory

	walkDirs(ROOTDIR, func(dir *Directory) {
		if size := dir.TotalSize(); size > missing && size < answer {
			answer = size
			answerDir = dir
		}
	})

	fmt.Printf("(Part 2) Deleting the directory %s would free up %d \n", answerDir, answer)
}

func readInput() {
	f, err := os.Open(InputFile)
	if err != nil {
		panic(fmt.Sprintf("Could not open file %v \n", err))
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	cwd := ROOTDIR
	cmd := make([]string, 0)
	output := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		if args[0] == "$" {
			exec(&cwd, cmd, output)
			cmd = args[1:]
			output = make([]string, 0)
		} else {
			output = append(output, line)
		}
	}
	exec(&cwd, cmd, output)

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

// directory struct

type Directory struct {
	Name   string
	Parent *Directory
	Files  map[string]int
	Dirs   map[string]*Directory
}

func NewDirectory(name string, parent *Directory) *Directory {
	return &Directory{name, parent, make(map[string]int), make(map[string]*Directory)}
}

func (dir *Directory) String() string {
	return strings.Join(dir.Path(), "/")
}

func (dir *Directory) Path() []string {
	if dir.Parent != nil {
		return append(dir.Parent.Path(), dir.Name)
	}
	return []string{dir.Name}
}

func (dir *Directory) Tree() []string {
	tree := []string{}
	for name, dir := range dir.Dirs {
		tree = append(tree, fmt.Sprintf("├── %s", name))
		for _, line := range dir.Tree() {
			tree = append(tree, fmt.Sprintf("│   %s", line))
		}
	}
	for name := range dir.Files {
		tree = append(tree, fmt.Sprintf("├── %s", name))
	}
	return tree
}

func (dir *Directory) TreeString() string {
	return strings.Join(dir.Tree(), "\n")
}

func (dir *Directory) TotalSize() int {
	sum := 0
	for _, childDir := range dir.Dirs {
		sum += childDir.TotalSize()
	}
	for _, childFile := range dir.Files {
		sum += childFile
	}
	return sum
}

// commands and helpers

func exec(cwd **Directory, cmd []string, output []string) {
	if len(cmd) == 0 {
		return
	}
	switch cmd[0] {
	case "cd":
		newCwd := cd(*cwd, cmd[1])
		*cwd = newCwd
		fmt.Printf("cd %s: %s \n", cmd[1], newCwd)
	case "ls":
		fmt.Printf("ls: %v \n", output)
		ls(*cwd, output)
	}
}

func walkDirs(cwd *Directory, fn func(*Directory)) {
	fn(cwd)
	for i := range cwd.Dirs {
		walkDirs(cwd.Dirs[i], fn)
	}
}

func cd(cwd *Directory, path string) *Directory {
	if path == ".." {
		return cwd.Parent
	}
	if path == "/" {
		return ROOTDIR
	}
	dir, ok := cwd.Dirs[path]
	if ok {
		return dir
	}
	fmt.Println(strings.Join(ROOTDIR.Tree(), "\n"))
	panic(fmt.Sprintf("Could not cd from %s to %s", cwd, path))
}

func ls(cwd *Directory, output []string) {
	for _, line := range output {
		fields := strings.Fields(line)
		name := fields[1]
		if fields[0] == "dir" {
			cwd.Dirs[name] = NewDirectory(name, cwd)
			continue
		}
		size, err := strconv.Atoi(fields[0])
		if err != nil {
			panic(err)
		}
		cwd.Files[name] = size
	}
}
