package main

import (
	"fmt"
	"log"

	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	pairs := ParsePlan(input)

	Part1(pairs)
	Part2(pairs)
}

func Part1(pairs []Pair) {
	internal.PrintPartHeading("Part 1")

	tally := 0
	for _, pair := range pairs {
		if pair[0].FullyContains(pair[1]) || pair[1].FullyContains(pair[0]) {
			tally++
		}
	}
	fmt.Println("Number of assignment pairs where one assignment fully contains the other one:", tally)
}

func Part2(pairs []Pair) {
	fmt.Println()
	internal.PrintPartHeading("Part 2")

	tally := 0
	for _, pair := range pairs {
		if pair[0].Overlaps(pair[1]) {
			tally++
		}
	}
	fmt.Println("Number of assignment pairs where one assignment overlaps the other one:", tally)
}
