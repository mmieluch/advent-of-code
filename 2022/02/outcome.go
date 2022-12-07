package main

import "errors"

type Outcome string

const OutcomeWin Outcome = "WIN"

const OutcomeLoss Outcome = "LOSS"

const OutcomeDraw Outcome = "DRAW"

func NewOutcomeFromAlias(alias string) (Outcome, error) {
	switch alias {
	case "X":
		return OutcomeLoss, nil
	case "Y":
		return OutcomeDraw, nil
	case "Z":
		return OutcomeWin, nil
	default:
		return "", errors.New("unknown outcome alias: " + alias)
	}
}
