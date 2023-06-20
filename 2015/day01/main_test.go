package main

import (
	"testing"
)

func Test_part1(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{
			name:  "return zero for empty input",
			input: "",
			want:  0,
		},
		{
			name:  "positive one",
			input: "(",
			want:  1,
		},
		{
			name:  "negative one",
			input: ")",
			want:  -1,
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
			name:  "empty",
			input: "",
			want:  0,
		},
		{
			name:  "enter the basement on first move",
			input: ")",
			want:  1,
		},
		{
			name:  "enter the basement on 3rd move",
			input: "())",
			want:  3,
		},
		{
			name:  "Never enter the basement",
			input: "(((((((((",
			want:  0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := part2(tt.input); got != tt.want {
				t.Errorf("part2() = %v, want %v", got, tt.want)
			}
		})
	}
}
