package model

type Position struct {
	x, xMin, xMax int
	y, yMin, yMax int
}

func NewPosition(x, y, xMin, yMin, xMax, yMax int) Position {
	return Position{
		x:    x,
		xMin: xMin,
		xMax: xMax,
		y:    y,
		yMin: yMin,
		yMax: yMax,
	}
}

func (p Position) MoveLeft() Position {
	p.x--
	if p.x < p.xMin {
		p.x = p.xMin
	}
	return p
}

func (p Position) MoveRight() Position {
	p.x++
	if p.xMax < p.x {
		p.x = p.xMax
	}
	return p
}

func (p Position) MoveDown() Position {
	p.y++
	if p.yMax < p.y {
		p.y = p.yMax
	}
	return p
}

func (p Position) MoveUp() Position {
	p.y--
	if p.y < p.yMin {
		p.y = p.yMin
	}
	return p
}

func (p Position) X() int {
	return p.x
}

func (p Position) Y() int {
	return p.y
}
