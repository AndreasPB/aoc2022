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

func partOne(file *os.File) int {
	scanner := bufio.NewScanner(file)

	winningCombinations := []string{"A Y", "B Z", "C X"}
	drawingCombinations := []string{"A X", "B Y", "C Z"}

	points := 0
	for scanner.Scan() {
		strategy := scanner.Text()

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

	return points
}

func partTwo(file *os.File) int {
	scanner := bufio.NewScanner(file)

	const (
		Winning = "Z"
		Drawing = "Y"
		Losing  = "X"
	)

	const (
		Rock     = "A" // 1 point
		Paper    = "B" // 2 points
		Scissors = "C" // 3 points
	)

	points := 0
	for scanner.Scan() {
		strategy := scanner.Text()

		opponentHand := string(strategy[0])

		switch yourHand := string(strategy[2]); yourHand {
		case Winning:
			points += 6
			switch opponentHand {
			case Rock:
				points += 2
			case Paper:
				points += 3
			case Scissors:
				points += 1
			}
		case Drawing:
			points += 3
			switch opponentHand {
			case Rock:
				points += 1
			case Paper:
				points += 2
			case Scissors:
				points += 3
			}
		case Losing:
			points += 0
			switch opponentHand {
			case Rock:
				points += 3
			case Paper:
				points += 1
			case Scissors:
				points += 2
			}
		}
	}

	return points
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 1 answer
	fmt.Println("Points:", partOne(file))

	file, err = os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Part 2 answer
	fmt.Println("Points:", partTwo(file))
}
