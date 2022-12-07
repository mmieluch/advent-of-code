package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	fname := os.Args[1]
	input := loadInput(fname)

	Part1(input)
	Part2(input)
}

// Part1 of the puzzle parses the input into a strict strategy, where the
// expected round outcome is pre-determined, as the parses treats the second
// input column as actual shapes played by us, so there's no adjustments allowed
// to the result of the games.
func Part1(input string) {
	printPartHeading("Part 1")

	strategy, err := ParseStrictStrategy(input)
	if err != nil {
		log.Fatal("parsing strict strategy failed:", err)
	}
	printStrategy(strategy)
}

// Part2 of the puzzle focuses on constructing a strategy based on the desired
// output, rather than the logical one. This strategy will allow us to win the
// Rock, Paper, Scissors tournament without raising suspicions.
func Part2(input string) {
	fmt.Println("")
	printPartHeading("Part 2")

	strategy, err := ParseFlexibleStrategy(input)
	if err != nil {
		log.Fatal("parsing flexible strategy failed:", err)
	}
	printStrategy(strategy)
}

func loadInput(filename string) string {
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(body))
}

func printPartHeading(message string) {
	hBorder := strings.Repeat("#", len(message)+4)

	fmt.Println(hBorder)
	fmt.Printf("# %s #\n", message)
	fmt.Println(hBorder)
}

func printStrategy(s Strategy) {
	for i, r := range s {
		fmt.Printf("Round %d: %d\n", i+1, r.Score())
	}
	fmt.Printf("Strategy total score: %d\n", s.Total())
}
