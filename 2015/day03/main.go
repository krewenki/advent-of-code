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

type Santa struct {
	x, y int
}

func (s *Santa) move(dir string) {
	if dir == ">" {
		s.x++
	}
	if dir == "<" {
		s.x--
	}
	if dir == "^" {
		s.y++
	}
	if dir == "v" {
		s.y--
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
	houses := make(map[string]int)
	santa := &Santa{0, 0}
	houses["0,0"]++
	for _, d := range strings.Split(input, "") {
		santa.move(d)
		houses[fmt.Sprintf("%d,%d", santa.x, santa.y)]++
	}
	return len(houses)
}

func part2(input string) int {
	houses := make(map[string]int)
	santa := &Santa{0, 0}
	robot := &Santa{0, 0}
	houses["0,0"]++
	for ind, d := range strings.Split(input, "") {
		if ind%2 == 0 {
			santa.move(d)
			houses[fmt.Sprintf("%d,%d", santa.x, santa.y)]++
		} else {
			robot.move(d)
			houses[fmt.Sprintf("%d,%d", robot.x, robot.y)]++
		}
	}
	return len(houses)
}
