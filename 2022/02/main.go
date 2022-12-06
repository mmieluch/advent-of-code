package main

import (
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	fname := os.Args[1]
	input := loadInput(fname)
	strategy, err := parseInput(input)
	if err != nil {
		log.Fatal("parsing strategy failed:", err)
	}

	for i, r := range strategy {
		fmt.Printf("Round %d: %d\n", i+1, r.Score())
	}
	fmt.Printf("Strategy total score: %d\n", strategy.Total())
}

type ShapeName string

const ShapeRock ShapeName = "Rock"

const ShapePaper ShapeName = "Paper"

const ShapeScissors ShapeName = "Scissors"

type ShapeValue int

const ShapeValRock = 1

const ShapeValPaper = 2

const ShapeValScissors = 3

type Shape struct {
	Name  ShapeName
	Value ShapeValue
}

type Round struct {
	Ord int
	Op  Shape
	Me  Shape
}

var scoreMatrix map[ShapeName]map[ShapeName]int = map[ShapeName]map[ShapeName]int{
	ShapeRock: {
		ShapeRock:     3,
		ShapePaper:    6,
		ShapeScissors: 0,
	},
	ShapePaper: {
		ShapeRock:     0,
		ShapePaper:    3,
		ShapeScissors: 6,
	},
	ShapeScissors: {
		ShapeRock:     6,
		ShapePaper:    0,
		ShapeScissors: 3,
	},
}

func (r Round) Score() int {
	result := scoreMatrix[r.Op.Name][r.Me.Name]
	shapeBonus := int(r.Me.Value)

	return result + shapeBonus
}

type Strategy []Round

func (s Strategy) Total() int {
	total := 0

	for _, r := range s {
		total += r.Score()
	}

	return total
}

func loadInput(filename string) string {
	body, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	return strings.TrimSpace(string(body))
}

func parseInput(input string) (Strategy, error) {
	lines := strings.Split(input, "\n")
	s := make(Strategy, len(lines))
	re := regexp.MustCompile(`(?P<op>[ABC]) (?P<me>[XYZ])`)

	for idx, line := range lines {
		if line == "" {
			continue
		}

		ss := re.FindStringSubmatch(line)

		if len(ss) != 3 {
			return Strategy{}, errors.New(fmt.Sprintf("invalid strategy line: %s", line))
		}

		// Opponent's shape
		ops, err := NewShape(ss[1])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", ss[1], err))
		}
		// My shape
		mys, err := NewShape(ss[2])
		if err != nil {
			return Strategy{}, errors.New(fmt.Sprintf("invalid shape alias: %s; %s", ss[2], err))
		}

		s[idx] = Round{
			Ord: idx + 1,
			Op:  ops,
			Me:  mys,
		}
	}

	return s, nil
}

func NewShape(alias string) (Shape, error) {
	switch alias {
	case "A", "X":
		return Shape{
			Name:  ShapeRock,
			Value: ShapeValRock,
		}, nil
	case "B", "Y":
		return Shape{
			Name:  ShapePaper,
			Value: ShapeValPaper,
		}, nil
	case "C", "Z":
		return Shape{
			Name:  ShapeScissors,
			Value: ShapeValScissors,
		}, nil
	}

	return Shape{}, errors.New(fmt.Sprintf("unknown alias for opponent's shape: %s", alias))
}
