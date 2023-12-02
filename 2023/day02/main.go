package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var input string

func init() {
	fileBytes, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Print(err)
	}
	input = string(fileBytes)
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
	var maxRed, maxGreen, maxBlue int
	var possible bool
	var sum, gameId int

	// from puzzle description
	maxRed = 12
	maxGreen = 13
	maxBlue = 14

	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		fmt.Sscanf(parts[0], "Game %d", &gameId)

		hands := strings.Split(parts[1], ";")
		possible = true
		for _, hand := range hands {
			cubes := strings.Split(hand, ",")
			for _, cube := range cubes {
				counts := strings.Fields(cube)
				num, _ := strconv.Atoi(counts[0])
				color := counts[1]
				switch color {
				case "red":
					possible = possible && (num <= maxRed)
				case "green":
					possible = possible && (num <= maxGreen)
				case "blue":
					possible = possible && (num <= maxBlue)
				}
			}
		}
		if possible {
			sum += gameId
		}
	}
	return sum
}

func part2(input string) int {
	var minRed, minGreen, minBlue int
	var sum, gameId int
	for _, line := range strings.Split(input, "\n") {
		parts := strings.Split(line, ":")
		fmt.Sscanf(parts[0], "Game %d", &gameId)

		hands := strings.Split(parts[1], ";")

		minRed = 0
		minGreen = 0
		minBlue = 0

		for _, hand := range hands {
			cubes := strings.Split(hand, ",")
			for _, cube := range cubes {
				counts := strings.Fields(cube)
				num, _ := strconv.Atoi(counts[0])
				color := counts[1]
				switch color {
				case "red":
					if num > minRed {
						minRed = num
					}
				case "green":
					if num > minGreen {
						minGreen = num
					}
				case "blue":
					if num > minBlue {
						minBlue = num
					}
				}
			}
		}
		sum += minRed * minGreen * minBlue
	}
	return sum
}
