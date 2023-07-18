package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
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

func atoi(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func part1(input string) int {
	r, _ := regexp.Compile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	grid := [1000][1000]int{}
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindStringSubmatch(line)
		if matches == nil {
			panic("no matches")
		}
		action := matches[1]
		startX, startY := atoi(matches[2]), atoi(matches[3])
		endX, endY := atoi(matches[4]), atoi(matches[5])

		if action == "turn on" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] = 1
				}
			}
		} else if action == "turn off" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] = 0
				}
			}
		} else if action == "toggle" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					if grid[x][y] == 0 {
						grid[x][y] = 1
					} else {
						grid[x][y] = 0
					}
				}
			}
		}
	}
	total := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			total += grid[x][y]
		}

	}

	return total
}

func part2(input string) int {
	r, _ := regexp.Compile(`(turn on|turn off|toggle) (\d+),(\d+) through (\d+),(\d+)`)
	grid := [1000][1000]int{}
	for _, line := range strings.Split(input, "\n") {
		matches := r.FindStringSubmatch(line)
		if matches == nil {
			panic("no matches")
		}
		action := matches[1]
		startX, startY := atoi(matches[2]), atoi(matches[3])
		endX, endY := atoi(matches[4]), atoi(matches[5])

		if action == "turn on" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] += 1
				}
			}
		} else if action == "turn off" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] -= 1
					if grid[x][y] < 0 {
						grid[x][y] = 0
					}
				}
			}
		} else if action == "toggle" {
			for x := startX; x <= endX; x++ {
				for y := startY; y <= endY; y++ {
					grid[x][y] += 2
				}
			}
		}
	}
	total := 0
	for x := 0; x < 1000; x++ {
		for y := 0; y < 1000; y++ {
			total += grid[x][y]
		}

	}

	return total
}
