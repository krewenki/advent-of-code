package main

import (
	"testing"
)

var example = ``

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: "ugknbfddgicrmopn",
			want:  1,
		},
		{
			name:  "example2",
			input: "aaa",
			want:  1,
		},
		{
			name:  "example3",
			input: "jchzalrnumimnmhp",
			want:  0,
		},
		{
			name:  "example4",
			input: "haegwjzuvuyypxyu",
			want:  0,
		},
		{
			name:  "example5",
			input: "dvszwmarrgswjxmb",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part1(tt.input); got != tt.want {
				t.Errorf("part1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_part2(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "example",
			input: "qjhvhtzxzqqjkmpb",
			want:  1,
		},
		{
			name:  "example2",
			input: "xxyxx",
			want:  1,
		},
		{
			name:  "example3",
			input: "uurcxstgmygtbstg",
			want:  0,
		},
		{
			name:  "example4",
			input: "ieodomkazucvgmuy",
			want:  0,
		},
		// {
		// 	name:  "actual",
		// 	input: input,
		// 	want:  0,
		// },
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
