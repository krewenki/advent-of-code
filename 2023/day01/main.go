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
	var sum int

	justDigits := regexp.MustCompile(`\d`)
	for _, line := range strings.Split(input, "\n") {
		nums := justDigits.FindAllString(line, -1)
		num, _ := strconv.Atoi(nums[0] + nums[len(nums)-1])
		fmt.Println(num, nums)
		sum += num
	}

	return sum
}

func part2(input string) int {
	var sum int

	digits := regexp.MustCompile(`1|2|3|4|5|6|7|8|9|one|two|three|four|five|six|seven|eight|nine`)
	for _, line := range strings.Split(input, "\n") {
		nums := digits.FindAllString(line, -1)
		first := nums[0]
		last := nums[len(nums)-1]
		num, _ := strconv.Atoi(words_to_digits(first) + words_to_digits(last))
		sum += num
	}

	return sum
}

func words_to_digits(input string) string {
	var output string

	output = strings.ReplaceAll(input, "one", "1")
	output = strings.ReplaceAll(output, "two", "2")
	output = strings.ReplaceAll(output, "three", "3")
	output = strings.ReplaceAll(output, "four", "4")
	output = strings.ReplaceAll(output, "five", "5")
	output = strings.ReplaceAll(output, "six", "6")
	output = strings.ReplaceAll(output, "seven", "7")
	output = strings.ReplaceAll(output, "eight", "8")
	output = strings.ReplaceAll(output, "nine", "9")
	output = strings.ReplaceAll(output, "zero", "0")

	return output
}
