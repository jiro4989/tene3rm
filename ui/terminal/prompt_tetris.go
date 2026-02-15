package terminal

import (
	"fmt"
	"time"

	"github.com/jiro4989/tene3rm/domain/tetris"
	"github.com/nsf/termbox-go"
)

// promptWithTetris はEasyModeのテトリスを表示する。
func promptWithTetris(_ string) (bool, error) {
	if err := termbox.Init(); err != nil {
		return false, err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	t := tetris.NewTetris()

	go startTetrisGameTimer(t)

	waitTetrisKeyInput(t)

	return false, nil
}

func startTetrisGameTimer(t *tetris.Tetris) {
	for t.Running() {
		drawTetrisScreen(t)

		if t.MinoCanMoveDown() {
			t.MinoMoveDown()
		} else {
			t.PutMino()
			if t.MinoIsOverlap() {
				t.StopGame()
			}
		}

		time.Sleep(1 * time.Second)
	}
}

func waitTetrisKeyInput(t *tetris.Tetris) {
	for t.Running() {
		minoIsMoved := false

		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				t.StopGame()
			case termbox.KeySpace:
				t.MinoMoveBottom()
				minoIsMoved = true
			}

			switch ev.Ch {
			case 'h':
				t.MinoMoveLeft()
				minoIsMoved = true
			case 'j':
				t.MinoMoveDown()
				minoIsMoved = true
			case 'l':
				t.MinoMoveRight()
				minoIsMoved = true
			}
		}

		if minoIsMoved {
			drawTetrisScreen(t)
		}

		time.Sleep(50 * time.Millisecond)
	}
}

func drawTetrisScreen(t *tetris.Tetris) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const leftPad = 1
	const topPad = 1

	cells := t.PreviewCells()
	for y, cell := range cells {
		drawTetrisBoardRow(cell, leftPad, y+topPad)
	}
	drawLine(fmt.Sprintf("SCORE %d", t.ScorePoint()), 30, 3)

	termbox.Flush()
}

var tetrisColorMap = map[tetris.Cell]termbox.Attribute{
	tetris.CellEmpty: termbox.ColorWhite,
	tetris.CellWall:  termbox.ColorGreen,
	tetris.CellMino:  termbox.ColorRed,
}

func drawTetrisBoardRow(row []tetris.Cell, x, y int) {
	for i, c := range row {
		bg := tetrisColorMap[c]
		termbox.SetCell(2*(x+i), y, ' ', termbox.ColorDefault, bg)
		termbox.SetCell(2*(x+i)+1, y, ' ', termbox.ColorDefault, bg)
	}
}
