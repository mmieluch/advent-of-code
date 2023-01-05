package main

import (
	"strconv"
	"strings"
)

type row []uint

type matrix []row

func (m matrix) Row(y int) row {
	return m[y]
}

func (m matrix) Col(x int) []uint {
	var out []uint

	for _, row := range m {
		out = append(out, row[x])
	}

	return out
}

func (m matrix) Pos(x, y int) uint {
	return m[y][x]
}

func (m matrix) LeftPos(x, y int) []uint {
	row := m.Row(y)

	if x == 0 {
		return []uint{}
	}

	out := make([]uint, x)
	for i, n := range row[:x] {
		out[i] = n
	}

	return out
}

func (m matrix) RightPos(x, y int) []uint {
	row := m.Row(y)
	if x >= len(row)-1 {
		return []uint{}
	}

	out := make([]uint, len(row)-x-1)
	for i, n := range row[x+1:] {
		out[i] = n
	}

	return out
}

func (m matrix) AbovePos(x, y int) []uint {
	if y == 0 {
		return []uint{}
	}

	col := m.Col(x)
	out := make([]uint, y)

	for i, n := range col[:y] {
		out[i] = n
	}

	return out
}

func (m matrix) BelowPos(x, y int) []uint {
	col := m.Col(x)
	if y == len(col)-1 {
		return []uint{}
	}

	out := make([]uint, len(col)-y-1)

	for i, n := range col[y+1:] {
		out[i] = n
	}

	return out
}

func parse(input string) matrix {
	var m matrix

	for _, line := range strings.Split(input, "\n") {
		var r row

		chars := strings.Split(line, "")
		for _, ch := range chars {
			num, _ := strconv.Atoi(ch)
			r = append(r, uint(num))
		}
		m = append(m, r)
	}

	return m
}
