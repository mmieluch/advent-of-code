package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	Part1(input)
	Part2(input)
}

// Part1 of the puzzle parses the input into a strict strategy, where the
// expected round outcome is pre-determined, as the parses treats the second
// input column as actual shapes played by us, so there's no adjustments allowed
// to the result of the games.
func Part1(input string) {
	internal.PrintPartHeading("Part 1")

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
	internal.PrintPartHeading("Part 2")

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

func printStrategy(s Strategy) {
	for i, r := range s {
		fmt.Printf("Round %d: %d\n", i+1, r.Score())
	}
	fmt.Printf("Strategy total score: %d\n", s.Total())
}
