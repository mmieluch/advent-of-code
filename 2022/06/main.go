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

	Part1(input)
	Part2(input)
}

func Part1(input string) {
	internal.PrintPartHeading("Part 1")

	idx, err := SOPMarkerIdx([]byte(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start-of-packet marker index:", idx)
}

func Part2(input string) {
	fmt.Println()
	internal.PrintPartHeading("Part 2")

	idx, err := MsgMarkerIdx([]byte(input))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Start-of-message marker index:", idx)
}
