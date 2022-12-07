package main

var scoreMatrix map[ShapeName]map[ShapeName]int = map[ShapeName]map[ShapeName]int{
	ShapeNameRock: {
		ShapeNameRock:     3,
		ShapeNamePaper:    6,
		ShapeNameScissors: 0,
	},
	ShapeNamePaper: {
		ShapeNameRock:     0,
		ShapeNamePaper:    3,
		ShapeNameScissors: 6,
	},
	ShapeNameScissors: {
		ShapeNameRock:     6,
		ShapeNamePaper:    0,
		ShapeNameScissors: 3,
	},
}

type Round struct {
	Ord int
	Op  Shape
	Me  Shape
}

func (r Round) Score() int {
	result := scoreMatrix[r.Op.Name][r.Me.Name]
	shapeBonus := int(r.Me.Value)

	return result + shapeBonus
}
