package aoc

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func GetStats(cookie string) (int, map[string]int) {
	years := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022}
	allYearsTotal := 0
	stars := make(map[string]int)
	for _, year := range years {
		total, _ := GetStatsForYear(year, cookie)

		stars[fmt.Sprintf("%d", year)] = total
		allYearsTotal += total
	}

	table := generateTable(stars, allYearsTotal)

	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	// write to file
	filename := filepath.Join(path, "_stats.md")
	WriteToFile(filename, []byte(table))

	return allYearsTotal, stars
}

func GetStatsForYear(year int, cookie string) (int, map[string]int) {
	days := make(map[string]int)
	total := 0
	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d", year)
	body := GetWithAOCCookie(url, cookie)

	r := regexp.MustCompile(`<a aria-label="Day ([0-9]*)[, ]*([a-z ]*)`)

	for _, l := range strings.Split(string(body), "\n") {

		matches := r.FindStringSubmatch(l)
		if len(matches) == 3 {
			date := fmt.Sprintf("'%02s'", matches[1])
			if matches[2] == "two stars" {
				days[date] = 2
				total += 2
			} else if matches[2] == "one star" {
				days[date] = 1
				total += 1
			} else {
				days[date] = 0
			}

		} else {
			continue
		}

	}
	return total, days
}

func generateTable(stars map[string]int, total int) string {
	years := []int{2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022}
	output := []string{}
	output = append(output, "| Year | Stars |\n")
	output = append(output, "| ---- | ----- |\n")
	for _, year := range years {
		output = append(output, fmt.Sprintf("| %d | %03d ⭐ |\n", year, stars[fmt.Sprintf("%d", year)]))
	}

	output = append(output, fmt.Sprintf("| Total | %03d ⭐ |\n", total))

	return strings.Join(output, "")
}
