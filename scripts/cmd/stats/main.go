package main

import (
	"github.com/krewenki/advent-of-code/scripts/aoc"
)

func main() {
	_, _, cookie := aoc.ParseFlags()

	aoc.GetStats(cookie)
}
