package main

import (
	"fmt"
	"log"

	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInputTrimmed()
	if err != nil {
		log.Fatal(err)
	}
	m := parse(input)

	var tt []tree

	for y, row := range m {
		for x, _ := range row {
			tt = append(tt, newTreeFromMatrix(m, x, y))
		}
	}

	Part1(tt)
	Part2(tt)
}

func Part1(tt []tree) {
	internal.PrintPartHeading("Part 1")

	total := 0
	for _, t := range tt {
		if t.IsVisible() {
			total++
		}
	}

	fmt.Println("Total number of trees visible from outside the grid:", total)
}

func Part2(tt []tree) {
	fmt.Println()
	internal.PrintPartHeading("Part 2")

	max := uint(0)

	for _, t := range tt {
		s := t.ScenicScore()
		if s > max {
			max = s
		}
	}

	fmt.Println("The highest scenic score of all the visible trees is:", max)
}
