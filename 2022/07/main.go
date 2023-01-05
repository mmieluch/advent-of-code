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

	ops, err := parseInput(input)
	if err != nil {
		log.Fatalf("couldn't parse input: %s\n", err)
	}
	root, err := parseOperations(ops)
	if err != nil {
		log.Fatal(fmt.Errorf("couldn't parse operations: %w", err))
	}

	Part1(root)
	Part2(root)
}

func Part1(root Dir) {
	internal.PrintPartHeading("Part 1")

	total := uint(0)
	found := smallDirs(root.Children(), 100000)
	for _, f := range found {
		total += f.Size()
	}

	fmt.Println("The total size of all directories which individual sizes exceed 100000 is:", total)
}

func Part2(root Dir) {
	fmt.Println()
	internal.PrintPartHeading("Part 2")

	spaceTotal := uint(70000000)
	spaceRequired := uint(30000000)
	spaceUsed := root.Size()
	spaceUnused := spaceTotal - spaceUsed

	spaceMissing := spaceRequired - spaceUnused

	fmt.Println("Space total:", spaceTotal)
	fmt.Println("Space required:", spaceRequired)
	fmt.Println("Space used:", spaceUsed)
	fmt.Println("Space unused:", spaceUnused)
	fmt.Println("Still need to free additional:", spaceMissing)

	// The puzzle idiotically assumes that the root as a candidate for removal
	// as well, and its only protection is the fact that there are smaller
	// directories, whose removals would suffice to satisfy the space requirements.
	// Honestly, this puzzle is incredibly dumb. I'm not even going to entertain
	// the idea of removing the entire filesystem in order to install a software
	// update.
	dirs := dirSizes(root.Children())
	candidates := filterDirs(dirs, spaceMissing)

	var sName string
	sSize := spaceTotal

	for name, size := range candidates {
		if size < sSize {
			sName = name
			sSize = size
		}
	}

	fmt.Printf(
		"Smallest directory to remove to satisfy upgrade space requirements: %s [%d]\n",
		sName,
		sSize,
	)
}

// filterDirs takes a map of directory nodes and returns a similar map of only
// the nodes whose size matches or exceeds the given limit.
func filterDirs(dirs map[string]uint, limit uint) map[string]uint {
	out := make(map[string]uint)

	for name, size := range dirs {
		if size >= limit {
			out[name] = size
		}
	}

	return out
}

// dirSizes takes a slice of Node structs, and returns a map where keys are
// directory names, and the values are the respective total sizes of the dirs.
// All files are discarded from the output map.
func dirSizes(nodes []Node) map[string]uint {
	out := make(map[string]uint)

	for _, n := range nodes {
		if !n.IsDir() {
			continue
		}

		out[n.Name()] = n.Size()
		children := dirSizes(n.(Parent).Children())
		for name, size := range children {
			out[name] = size
		}
	}

	return out
}

// smallDirs takes a slice of dirs (marked as a slice of Node interfaces) and
// returns only the subset with total sizes smaller than the given limit.
func smallDirs(nodes []Node, limit uint) []Node {
	var found []Node

	for _, n := range nodes {
		if n.IsDir() {
			if n.Size() <= limit {
				found = append(found, n)
			}
			found = append(found, smallDirs(n.(Parent).Children(), limit)...)
		}
	}

	return found
}
