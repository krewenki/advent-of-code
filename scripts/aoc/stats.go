package aoc

import (
	"fmt"
	"regexp"
	"strings"
)

func GetStats(year int, cookie string) (int, map[string]int) {
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
