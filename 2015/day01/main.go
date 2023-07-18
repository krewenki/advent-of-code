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
