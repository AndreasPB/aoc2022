// Day 1: Calorie Counting

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	var elves []int
	var currentCalories int
	for scanner.Scan() {
		if scanner.Text() != "" {
			calories, err := strconv.Atoi(scanner.Text())
			if err != nil {
				log.Fatalln(err)
			}

			currentCalories += int(calories)
		} else {
			// Reset and append
			elves = append(elves, currentCalories)
			currentCalories = 0
		}
	}

	sort.Ints(elves)

	// Highest calories - Part 1
	fmt.Println("Highest calories", elves[len(elves)-1])

	// Top 3 combined - Part 2
	var topTreeCombined int
	for _, calories := range elves[len(elves)-3:] {
		topTreeCombined += calories
	}

	fmt.Println("Top 3 combined:", topTreeCombined)
}
