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

func isVowel(s string) bool {
	return s == "a" || s == "e" || s == "i" || s == "o" || s == "u"
}

func isNiceV1(s string) bool {
	var numberOfVowels int
	var hasRepeatingCharacters bool
	var currentLetter, lastLetter, pair string

	// does string contain a double letter
	for i := 0; i < len(s); i++ {
		currentLetter = string(s[i])

		if isVowel(string(s[i])) {
			numberOfVowels++
		}
		if i > 0 {
			pair = lastLetter + currentLetter

			// pair can't be ab, cd, pq, or xy
			if pair == "ab" || pair == "cd" || pair == "pq" || pair == "xy" {
				return false
			}

			if s[i] == s[i-1] {
				hasRepeatingCharacters = true
			}
		}
		lastLetter = currentLetter
	}

	return numberOfVowels > 2 && hasRepeatingCharacters

}

func isNiceV2(s string) bool {
	var hasRepeatingCharacters, hasRepeatsSeparatedByOne bool
	var current, trailingOne, trailingTwo, pair string
	var distMap = make(map[string][]int)

	// does string contain a double letter
	for i := 1; i < len(s); i++ {
		current = string(s[i])
		trailingOne = string(s[i-1])
		if i > 1 {
			trailingTwo = string(s[i-2])
		} else {
			trailingTwo = " "
		}
		pair = trailingOne + current
		distMap[pair] = append(distMap[pair], i-1)

		if current == trailingTwo {
			hasRepeatsSeparatedByOne = true
		}
	}
	for _, v := range distMap {
		if len(v) > 1 {
			var min, max int
			min = 1000
			max = 0
			for _, i := range v {
				if i < min {
					min = i
				}
				if i > max {
					max = i
				}
			}
			if max-min > 1 {
				hasRepeatingCharacters = true
			}
		}
	}
	return hasRepeatingCharacters && hasRepeatsSeparatedByOne

}

func part1(input string) int {
	var count int
	for _, line := range strings.Split(input, "\n") {
		if isNiceV1(line) {
			count++
		}
	}
	return count
}

func part2(input string) int {
	var count int
	for _, line := range strings.Split(input, "\n") {
		if isNiceV2(line) {
			count++
		}
	}
	return count
}
