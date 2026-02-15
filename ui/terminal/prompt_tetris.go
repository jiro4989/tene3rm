package terminal

import (
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
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyEsc:
				t.StopGame()
			case termbox.KeySpace:
				t.MinoMoveBottom()
			}

			switch ev.Ch {
			case 'h':
				t.MinoMoveLeft()
			case 'j':
				t.MinoMoveDown()
			case 'l':
				t.MinoMoveRight()
			}
		}

		time.Sleep(50 * time.Millisecond)
	}
}
