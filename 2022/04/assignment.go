package main

import (
	"strconv"
	"strings"
)

type Assignment []uint

func (a Assignment) ContainsSection(needle uint) bool {
	for _, val := range a {
		if val == needle {
			return true
		}
	}
	return false
}

func (a Assignment) FullyContains(target Assignment) bool {
	for _, val := range target {
		if !a.ContainsSection(val) {
			return false
		}
	}
	return true
}

func (a Assignment) Overlaps(target Assignment) bool {
	for _, val := range target {
		if a.ContainsSection(val) {
			return true
		}
	}
	return false
}

func NewAssignment(min, max uint) Assignment {
	distance := max - min + 1
	a := make(Assignment, distance)
	for i := uint(0); i < distance; i++ {
		a[i] = min + i
	}
	return a
}

type Pair []Assignment

func ParsePlan(input string) []Pair {
	var pairs []Pair

	for _, line := range strings.Split(input, "\n") {
		var pair Pair

		for _, r := range strings.Split(line, ",") {
			min, max := extractLimits(r)
			pair = append(pair, NewAssignment(min, max))
		}

		pairs = append(pairs, pair)
	}

	return pairs
}

func extractLimits(entry string) (min, max uint) {
	limits := strings.Split(entry, "-")

	_min, _ := strconv.Atoi(limits[0])
	_max, _ := strconv.Atoi(limits[1])

	return uint(_min), uint(_max)
}
