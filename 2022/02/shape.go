package main

import (
	"errors"
	"fmt"
)

type ShapeName string

const ShapeNameRock ShapeName = "Rock"

const ShapeNamePaper ShapeName = "Paper"

const ShapeNameScissors ShapeName = "Scissors"

type ShapeValue int

const ShapeValRock = 1

const ShapeValPaper = 2

const ShapeValScissors = 3

type Shape struct {
	Name  ShapeName
	Value ShapeValue
}

var (
	Rock = Shape{
		Name:  ShapeNameRock,
		Value: ShapeValRock,
	}
	Paper = Shape{
		Name:  ShapeNamePaper,
		Value: ShapeValPaper,
	}
	Scissors = Shape{
		Name:  ShapeNameScissors,
		Value: ShapeValScissors,
	}
)

func (s Shape) LosesTo() (Shape, error) {
	switch s.Name {
	case ShapeNameRock:
		return Paper, nil
	case ShapeNamePaper:
		return Scissors, nil
	case ShapeNameScissors:
		return Rock, nil
	default:
		return Shape{}, errors.New("shape name not handled: " + string(s.Name))
	}
}

func (s Shape) WinsOver() (Shape, error) {
	switch s.Name {
	case ShapeNameRock:
		return Scissors, nil
	case ShapeNamePaper:
		return Rock, nil
	case ShapeNameScissors:
		return Paper, nil
	default:
		return Shape{}, errors.New("shape name not handled: " + string(s.Name))
	}
}

func NewShapeFromAlias(alias string) (Shape, error) {
	switch alias {
	case "A", "X":
		return Rock, nil
	case "B", "Y":
		return Paper, nil
	case "C", "Z":
		return Scissors, nil
	}

	return Shape{}, errors.New(fmt.Sprintf("unknown alias for opponent's shape: %s", alias))
}

func NewShapeForOutcome(ops Shape, outcomeAlias string) (Shape, error) {
	// First, we need to determine the outcome based on the alias.
	outcome, err := NewOutcomeFromAlias(outcomeAlias)
	if err != nil {
		return Shape{}, err
	}

	switch outcome {
	case OutcomeLoss:
		return ops.WinsOver()
	case OutcomeDraw:
		return ops, nil
	case OutcomeWin:
		return ops.LosesTo()
	default:
		return Shape{}, errors.New("unhandled outcome: " + string(outcome))
	}
}
