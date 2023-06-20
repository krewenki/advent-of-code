package main

import (
	_ "embed"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
	input = strings.TrimRight(input, "\n")
	if len(input) == 0 {
		panic("empty input.txt file")
	}
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
	var sum int
	for _, c := range input {
		if c == '(' {
			sum++
		} else if c == ')' {
			sum--
		}
	}
	return sum
}

func part2(input string) int {
	var sum int
	for ind, c := range input {
		if c == '(' {
			sum++
		} else if c == ')' {
			sum--
			if sum < 0 {
				return ind + 1
			}
		}
	}
	return 0
}
