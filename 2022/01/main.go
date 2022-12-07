package main

import (
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"

	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	elves := parseInput(input)

	// Part 1
	runPart1(elves)
	// Part 2
	runPart2(elves)
}

func runPart1(ee Elves) {
	topCarrier := ee.TopCarrier()
	fmt.Println("The elf carrying the most calorific snacks is carrying the total of", topCarrier.GetTotal(), "calories.")
}

func runPart2(ee Elves) {
	totals := make([]int, len(ee))

	// Get the totals for each elf. We don't need to know which specific elves
	// are the top 3 carriers, only how much each of them carries on them.
	for idx, elf := range ee {
		totals[idx] = int(elf.GetTotal())
	}

	sort.Sort(sort.Reverse(sort.IntSlice(totals)))
	top := totals[:3]
	topTotal := 0

	for _, val := range top {
		topTotal += val
	}

	fmt.Println("The three elves carrying the most calorific snacks carry", topTotal, "calories in total.")
}

func loadInput(fname string) []byte {
	f, err := os.ReadFile(fname)
	if err != nil {
		log.Fatal(err)
	}
	return f
}

type Elves []Elf

func (ee Elves) TopCarrier() *Elf {
	maxKey := 0
	max := uint(0)

	for key, e := range ee {
		etotal := e.GetTotal()
		if etotal > max {
			max = etotal
			maxKey = key
		}
	}

	return &ee[maxKey]
}

type Elf struct {
	foodItems []uint
}

func (e *Elf) AddFoodItem(cals uint) {
	e.foodItems = append(e.foodItems, cals)
}

func (e *Elf) GetTotal() uint {
	total := uint(0)

	for _, calories := range e.foodItems {
		total += calories
	}

	return total
}

func parseInput(input string) Elves {
	segs := strings.Split(input, "\n\n")

	var ee Elves

	for _, seg := range segs {
		elf := Elf{}

		cals := strings.Split(seg, "\n")

		for _, cal := range cals {
			ical, _ := strconv.Atoi(string(cal))
			elf.AddFoodItem(uint(ical))
		}

		ee = append(ee, elf)
	}

	return ee
}
