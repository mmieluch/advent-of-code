package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fname := os.Args[1]
	input := loadInput(fname)
	elves := parseInput(input)

	topCarrier := elves.TopCarrier()

	fmt.Println(topCarrier.GetTotal())
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

func parseInput(input []byte) Elves {
	input = bytes.TrimSpace(input)
	segs := bytes.Split(input, []byte("\n\n"))

	var ee Elves

	for _, seg := range segs {
		elf := Elf{}

		cals := bytes.Split(seg, []byte("\n"))

		for _, cal := range cals {
			ical, _ := strconv.Atoi(string(cal))
			elf.AddFoodItem(uint(ical))
		}

		ee = append(ee, elf)
	}

	return ee
}
