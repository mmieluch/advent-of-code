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

}

func Part1(ss operations.Stacks, ii []operations.Instruction) {
	internal.PrintPartHeading("Part 1")

	updated, err := operations.Reorder(ss, ii)
	if err != nil {
		log.Fatal(err)
	}

	topItems := ""
	for i := 1; i <= len(updated); i++ {
		s := ss[i]
		topItems += s.GetTopItem()
	}

	fmt.Println("Result:", topItems)
}
