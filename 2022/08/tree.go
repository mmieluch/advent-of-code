package main

type tree struct {
	height uint
	left   []uint
	right  []uint
	above  []uint
	below  []uint
}

func (t *tree) IsHidden() bool {
	if !isSurrounded(t) {
		return false
	}

	if t.height == 0 {
		return true
	}

	return isHiddenOnSide(t.height, t.left) &&
		isHiddenOnSide(t.height, t.right) &&
		isHiddenOnSide(t.height, t.above) &&
		isHiddenOnSide(t.height, t.below)
}

func isHiddenOnSide(height uint, others []uint) bool {
	for _, val := range others {
		if val >= height {
			return true
		}
	}
	return false
}

func isSurrounded(t *tree) bool {
	return len(t.left) > 0 && len(t.right) > 0 && len(t.above) > 0 && len(t.below) > 0
}

func (t *tree) IsVisible() bool {
	return !t.IsHidden()
}

func (t *tree) ScenicScore() uint {
	if !isSurrounded(t) {
		return 0
	}

	l := viewingDistance(t.height, reverseNeighbours(t.left))
	r := viewingDistance(t.height, t.right)
	a := viewingDistance(t.height, reverseNeighbours(t.above))
	b := viewingDistance(t.height, t.below)

	return l * r * a * b
}

func reverseNeighbours(tt []uint) []uint {
	var out []uint

	for i := len(tt) - 1; i > -1; i-- {
		out = append(out, tt[i])
	}

	return out
}

func viewingDistance(h uint, tt []uint) uint {
	if len(tt) == 0 {
		return uint(0)
	}

	var total uint

	for _, t := range tt {
		// if t > 0 {
		total++
		// }

		if t >= h {
			break
		}
	}

	return total
}

func newTreeFromMatrix(m matrix, x, y int) tree {
	return tree{
		height: m.Pos(x, y),
		left:   m.LeftPos(x, y),
		right:  m.RightPos(x, y),
		above:  m.AbovePos(x, y),
		below:  m.BelowPos(x, y),
	}
}
