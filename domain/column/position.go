package column

import "github.com/jiro4989/tene3rm/domain"

type Position struct {
	x, y domain.RangeInt
}

func NewPosition(x, y domain.RangeInt) Position {
	return Position{
		x: x,
		y: y,
	}
}

func (p Position) MoveLeft() Position {
	p.x = p.x.SafePlus(-1)
	return p
}

func (p Position) MoveRight() Position {
	p.x = p.x.SafePlus(1)
	return p
}

func (p Position) MoveDown() Position {
	p.y = p.y.SafePlus(1)
	return p
}

func (p Position) MoveUp() Position {
	p.y = p.y.SafePlus(-1)
	return p
}

func (p Position) X() int {
	return p.x.Value()
}

func (p Position) Y() int {
	return p.y.Value()
}
