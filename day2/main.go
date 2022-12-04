// Day 2: Rock Paper Scissors

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func contains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

func main() {
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)

	winningCombinations := []string{"A Y", "B Z", "C X"}
	drawingCombinations := []string{"A X", "B Y", "C Z"}

	points := 0
	for scanner.Scan() {
		strategy := scanner.Text()
		if err != nil {
			log.Fatalln(err)
		}

		switch yourHand := string(strategy[2]); yourHand {
		case "X":
			points += 1
		case "Y":
			points += 2
		case "Z":
			points += 3
		}

		if contains(winningCombinations, strategy) {
			points += 6
		} else if contains(drawingCombinations, strategy) {
			points += 3
		}

	}

	// Part 1 answer
	fmt.Println("Points:", points)
}
