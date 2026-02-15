package tetris

// 配列のコピーコストが高くなると嫌なので、
// ポインターレシーバーにして構造体の複製が発生しないようにする。

type Tetris struct {
	board Board
	mino  Mino
	score Score
}

func NewTetris() Tetris {
	return Tetris{
		board: newBoard(),
	}
}

func (t *Tetris) MinoMove(f func() Mino) {
	mino := f()
	if t.board.canMove(mino.x, mino.y) {
		t.mino = mino
	}
}

func (t *Tetris) MinoMoveLeft() {
	t.MinoMove(t.mino.MoveLeft)
}

func (t *Tetris) MinoMoveRight() {
	t.MinoMove(t.mino.MoveRight)
}

func (t *Tetris) MinoMoveDown() {
	t.MinoMove(t.mino.MoveDown)
}

func (t *Tetris) MinoCanMoveDown() bool {
	x, y := t.mino.x, t.mino.y
	return t.board.canMove(x, y+1)
}

func (t *Tetris) ScorePlus() {
	t.score = t.score.Plus()
}

func (t *Tetris) DeleteRows() {
	for i := 0; i < 25; i++ {
		row := t.board.value[i]
		if row.IsFulfilled() {
			t.board.clearRow(i)
		}
	}
}

type Cell int

func (c Cell) IsNotEmpty() bool {
	return c != cellEmpty
}

const (
	cellEmpty Cell = 0
	cellWall
	cellMino
)

type Row struct {
	value []Cell
}

func newRow(value []Cell) Row {
	return Row{
		value: value,
	}
}

func (r Row) IsFulfilled() bool {
	for _, cell := range r.value {
		if cell.IsNotEmpty() {
			continue
		}
		return false
	}
	return true
}

func newEmptyRow() Row {
	value := []Cell{
		cellWall, cellWall, cellWall, cellEmpty, cellEmpty, cellEmpty, cellEmpty, cellEmpty, cellEmpty, cellEmpty, cellEmpty, cellWall, cellWall, cellWall,
	}
	return newRow(value)
}

func newBottomRow() Row {
	value := []Cell{
		cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall, cellWall,
	}
	return newRow(value)
}

type Board struct {
	value []Row
}

func newBoard() Board {
	value := []Row{}
	for i := 0; i < 25; i++ {
		value = append(value, newEmptyRow())
	}
	for i := 0; i < 3; i++ {
		value = append(value, newBottomRow())
	}
	return Board{
		value: value,
	}
}

func (b *Board) cell(x, y int) Cell {
	return b.value[y].value[x]
}

func (b *Board) canMove(x, y int) bool {
	return b.cell(x, y).IsNotEmpty()
}

func (b *Board) clearRow(y int) {
	b.value[y] = newEmptyRow()
}
