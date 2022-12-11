// Day 3: Rucksack Reorganization

package main

import (
	"fmt"
	"os"
	"strings"
)

func charToInt(c rune) int {
	if int(c) > 90 {
		// lower case
		return int(c) - 96
	} else {
		// upper case
		return int(c) - 38
	}
}

func partOne(input []byte) int {
	rucksack := strings.Split(strings.TrimSpace(string(input)), "\n")
	total := 0
	for _, s := range rucksack {
		sideOne := s[len(s)/2:]
		sideTwo := s[:len(s)/2]
		for _, c := range sideOne {
			if strings.Contains(sideTwo, string(c)) {
				total += charToInt(c)
				break
			}
		}
	}
	return total
}
func main() {
	input, _ := os.ReadFile("input.txt")

	// Part One answer
	fmt.Println(partOne(input))

	// Part Two answer

}
