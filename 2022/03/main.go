package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/mmieluch/advent-of-code/2022/03/efficiency"
	"github.com/mmieluch/advent-of-code/2022/03/storage"
	"github.com/mmieluch/advent-of-code/2022/internal"
)

func main() {
	input, err := internal.GetInput()
	if err != nil {
		log.Fatal(err)
	}

	lines := parseManifest(input)
	rr, err := pack(lines)
	if err != nil {
		log.Fatal(fmt.Errorf("packing rucksacks failed: %w", err))
	}

	Part1(rr)
}

func Part1(rr []storage.Rucksack) {
	internal.PrintPartHeading("Part 1")

	total := uint(0)
	for _, r := range rr {
		// Extract the intersection
		is := r.CommonItems()
		// Determine the priority of each common item
		for _, item := range is {
			p, err := efficiency.Priority(item)
			if err != nil {
				log.Fatal("Couldn't determine priority for the item:", string(item))
			}
			log.Printf("Priority for item %s is %d", string(item), p)
			total += uint(p)
		}
	}

	log.Println("Total of priorities of common items:", total)
}

// pack takes a slice of rune slices representing individual rucksack items, and
// returns a slice of Rucksacks created using the items.
func pack(items [][]rune) ([]storage.Rucksack, error) {
	rr := make([]storage.Rucksack, len(items))

	for i, item := range items {
		rs, err := storage.NewRucksack(item)
		if err != nil {
			return []storage.Rucksack{}, err
		}
		rr[i] = rs
	}

	return rr, nil
}

// parseManifest breaks up the input string into individual lines, then splits
// each line into individual rules. Finally, it returns a slice of rune slices.
func parseManifest(input string) [][]rune {
	lines := strings.Split(input, "\n")
	var out [][]rune

	for _, line := range lines {
		out = append(out, []rune(line))
	}

	return out
}
