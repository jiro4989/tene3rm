package tetris

type Mino struct {
	value int
	x, y  int
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
