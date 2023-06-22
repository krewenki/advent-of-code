package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"strings"
)

//go:embed input.txt
var input string

func init() {
	// do this in init (not main) so test file has same input
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

func getMD5(text string) string {
	hash := md5.Sum([]byte(text))
	return hex.EncodeToString(hash[:])
}

func part1(input string) int {
	for i := 0; i <= 10000000000; i++ {
		md5 := getMD5(input + fmt.Sprint(i))
		firstFive := md5[:5]
		if firstFive == "00000" {
			return i
		}
	}
	return 0
}

func part2(input string) int {
	for i := 0; i <= 10000000000; i++ {
		md5 := getMD5(input + fmt.Sprint(i))
		firstSix := md5[:6]
		if firstSix == "000000" {
			return i
		}
	}
	return 0
}
