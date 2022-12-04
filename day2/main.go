package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

var scoreMapping = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
	"A": 1,
	"B": 2,
	"C": 3,
}

func main() {
	data, err := os.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	totalScore := 0

	scoreLines := strings.Split(string(data), "\n")
	for _, scoreLine := range scoreLines {
		totalScore += roundScore(scoreLine)
	}

	fmt.Println("your total score:", totalScore)
}

func ptsForShape(shape string) int {
	return scoreMapping[shape]
}

// Returns 0, 3 or 6, expects a scoreline string in the format `A X` as input
func ptsForOutcome(scoreLine string) int {
	outCome := string(scoreLine[2])
	switch outCome {
	case "X":
		return 0
	case "Y":
		return 3
	default:
		return 6
	}
}

// X means you need to lose,
// Y means you need to end the round in a draw,
// and Z means you need to win."
func yourShapeFromOutcome(scoreLine string) string {
	switch scoreLine {
	case "A X":
		return "Z"
	case "B X":
		return "X"
	case "C X":
		return "Y"

	case "A Y":
		return "X"
	case "B Y":
		return "Y"
	case "C Y":
		return "Z"

	case "A Z":
		return "Y"
	case "B Z":
		return "Z"
	case "C Z":
		return "X"
	default:
		return "X"
	}
}

// Expect a string in the format `A X` as input, returns the final score
func roundScore(scoreLine string) int {
	yourShape := yourShapeFromOutcome(scoreLine)
	return ptsForShape(yourShape) + ptsForOutcome(scoreLine)
}
