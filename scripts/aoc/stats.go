package aoc

import (
	"fmt"
	"strings"
)

func GetStats(year int, cookie string) {
	fmt.Printf("fetching for year %d\n", year)

	// make the request
	url := fmt.Sprintf("https://adventofcode.com/%d/leaderboard/self", year)
	body := GetWithAOCCookie(url, cookie)

	for _, l := range strings.Split(string(body), "\n") {
		fmt.Println(l)

	}
}
