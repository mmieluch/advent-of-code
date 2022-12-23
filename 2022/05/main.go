package main

import (
	"fmt"
	"log"

	"github.com/mmieluch/advent-of-code/2022/05/operations"
	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	stacks := operations.ParseStacks(input)
	instructions, err := operations.ParseInstructions(input)
	if err != nil {
		log.Fatal(err)
	}

	Part1(stacks, instructions)
	Part2(stacks, instructions)
}

func Part1(ss operations.Stacks, ii []operations.Instruction) {
	internal.PrintPartHeading("Part 1")

	ss = operations.CloneStacks(ss)
	updated, err := operations.ReorderSequentially(ss, ii)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", updated.GetTopItems())
}

func Part2(ss operations.Stacks, ii []operations.Instruction) {
	fmt.Println()
	internal.PrintPartHeading("Part 2")

	ss = operations.CloneStacks(ss)
	updated, err := operations.ReorderGrouped(ss, ii)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Result:", updated.GetTopItems())
}
