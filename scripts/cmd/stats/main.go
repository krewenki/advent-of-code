package main

import (
	"fmt"
	"time"

	"github.com/krewenki/advent-of-code/scripts/aoc"
)

func main() {
	_, _, cookie := aoc.ParseFlags()
	years := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022}
	allYearsTotal := 0
	stars := make(map[string]int)
	output := []string{}
	output = append(output, "| Year | Stars |\n")
	output = append(output, "| ---- | ----- |\n")
	for _, year := range years {
		total, _ := aoc.GetStats(year, cookie)

		stars[fmt.Sprintf("%d", year)] = total
		allYearsTotal += total

		// sleep for 1 second
		time.Sleep(time.Second)
	}

	for _, year := range years {
		output = append(output, fmt.Sprintf("| %d | %03d ⭐ |\n", year, stars[fmt.Sprintf("%d", year)]))
	}

	output = append(output, fmt.Sprintf("| Total | %03d ⭐ |\n", allYearsTotal))

}
