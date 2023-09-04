package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"
)

var input string

func init() {
	if _, err := os.Stat("input.txt"); err != nil {
		fileBytes, err := os.ReadFile("input.txt")
		if err != nil {
			fmt.Print(err)
		}
		input = string(fileBytes)
		input = strings.TrimRight(input, "\n")
	}
}

type Box struct {
	l, w, h int
}

func (b Box) surfaceArea() int {
	return (b.l*b.w)*2 + (b.w*b.h)*2 + (b.l*b.h)*2
}

func (b Box) smallestSide() int {
	return min(b.w*b.l, b.w*b.h, b.h*b.l)
}

func (b Box) shortestPermiter() int {
	x := (b.w + b.h) * 2
	y := (b.w + b.l) * 2
	z := (b.h + b.l) * 2
	return min(x, y, z)
}

func (b Box) volume() int {
	return b.w * b.h * b.l
}

func min(n ...int) int {
	min := n[0]
	for _, num := range n {
		if min > num {
			min = num
		}
	}
	return min
}

func main() {
	var part int
	flag.IntVar(&part, "part", 1, "part 1 or 2")
	flag.Parse()
	fmt.Println("Running part", part)

	if part == 1 {
		ans := part1(input)
		fmt.Println("Output:", ans)
	} else {
		ans := part2(input)
		fmt.Println("Output:", ans)
	}
}

func part1(input string) int {
	var total int
	for _, line := range strings.Split(input, "\n") {
		var l, w, h int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			panic(err)
		}
		b := Box{l, w, h}
		total += b.surfaceArea()
		total += b.smallestSide()
	}
	return total
}

func part2(input string) int {
	var total int
	for _, line := range strings.Split(input, "\n") {
		var l, w, h int
		_, err := fmt.Sscanf(line, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			panic(err)
		}
		b := Box{l, w, h}
		total += b.shortestPermiter()
		total += b.volume()
	}
	return total
}
