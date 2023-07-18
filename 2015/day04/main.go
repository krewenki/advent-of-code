package main

import (
	"crypto/md5"
	_ "embed"
	"encoding/hex"
	"flag"
	"fmt"
	"os"
	"strings"
)

var input string

func init() {
	// if input.txt exists
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
