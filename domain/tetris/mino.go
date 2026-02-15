package tetris

type Mino struct {
	value Cell
	x, y  int
}

func newDefaultMino() Mino {
	return Mino{
		value: cellMino,
		x:     7,
		y:     0,
	}
}

func (m Mino) MoveLeft() Mino {
	m.x--
	return m
}

func (m Mino) MoveRight() Mino {
	m.x++
	return m
}

func (m Mino) MoveDown() Mino {
	m.y++
	return m
}
