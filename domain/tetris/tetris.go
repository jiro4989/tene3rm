package tetris

import "sync"

// 配列のコピーコストが高くなると嫌なので、
// ポインターレシーバーにして構造体の複製が発生しないようにする。

type Tetris struct {
	board   Board
	mino    Mino
	score   Score
	running bool
	mu      sync.Mutex
}

func NewTetris() *Tetris {
	return &Tetris{
		board:   newDefaultBoard(),
		mino:    newDefaultMino(),
		score:   newDefaultScore(),
		running: true,
	}
}

func (t *Tetris) PreviewCells() [][]Cell {
	cells := [][]Cell{}
	for _, row := range t.board.value {
		dst := make([]Cell, len(row.value))
		copy(dst, row.value)
		cells = append(cells, dst)
	}

	// Mino の現在位置に Cell を反映した上で返却する
	cells[t.mino.y][t.mino.x] = t.mino.value

	return cells
}

func (t *Tetris) StopGame() {
	t.mu.Lock()
	t.running = false
	t.mu.Unlock()
}

func (t *Tetris) Running() bool {
	return t.running
}

func (t *Tetris) PutMino() {
	t.mu.Lock()

	t.board.value[t.mino.y].value[t.mino.x] = t.mino.value
	t.deleteRows()
	t.mino = newDefaultMino()

	t.mu.Unlock()
}

func (t *Tetris) MinoMove(f func() Mino) {
	mino := f()
	if t.board.canMove(mino.x, mino.y) {
		t.mu.Lock()
		t.mino = mino
		t.mu.Unlock()
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

func (t *Tetris) MinoMoveBottom() {
	for i := 0; i < 25; i++ {
		t.MinoMove(t.mino.MoveDown)
	}
}

func (t *Tetris) MinoCanMoveDown() bool {
	x, y := t.mino.x, t.mino.y
	return t.board.canMove(x, y+1)
}

func (t *Tetris) MinoIsOverlap() bool {
	x, y := t.mino.x, t.mino.y
	return !t.board.canMove(x, y)
}

func (t *Tetris) scorePlus() {
	t.mu.Lock()
	t.score = t.score.Plus()
	t.mu.Unlock()
}

func (t *Tetris) ScorePoint() int {
	return t.score.value
}

func (t *Tetris) deleteRows() {
	for i := 0; i < 25; i++ {
		row := t.board.value[i]
		if row.IsFulfilled() {
			t.board.clearRow(i)
			t.scorePlus()
		}
	}
}

type Cell int

func (c Cell) IsNotEmpty() bool {
	return c != CellEmpty
}

func (c Cell) IsEmpty() bool {
	return !c.IsNotEmpty()
}

const (
	CellEmpty Cell = iota
	CellWall
	CellMino
)

type Row struct {
	value []Cell
}

func newRow(value []Cell) Row {
	return Row{
		value: value,
	}
}

func (r *Row) IsFulfilled() bool {
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
		CellWall, CellWall, CellWall, CellEmpty, CellEmpty, CellEmpty, CellEmpty, CellEmpty, CellEmpty, CellEmpty, CellEmpty, CellWall, CellWall, CellWall,
	}
	return newRow(value)
}

func newBottomRow() Row {
	value := []Cell{
		CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall, CellWall,
	}
	return newRow(value)
}

type Board struct {
	value []Row
}

func newDefaultBoard() Board {
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
	return b.cell(x, y).IsEmpty()
}

func (b *Board) clearRow(y int) {
	b.value[y] = newEmptyRow()
}
