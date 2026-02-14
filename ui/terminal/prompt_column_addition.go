package terminal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jiro4989/tene3rm/domain/column"
	"github.com/nsf/termbox-go"
)

// promptWithColumnAddition は筆算での計算結果を求めるプロンプトを表示する。
func promptWithColumnAddition(path string) (bool, error) {
	a := rand.Intn(90) + 10
	b := rand.Intn(90) + 10

	cag, err := column.NewColumnAdditionGame(a, b)
	if err != nil {
		return false, err
	}

	if err := termbox.Init(); err != nil {
		return false, err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	chMap := map[rune]string{
		'0': "0",
		'1': "1",
		'2': "2",
		'3': "3",
		'4': "4",
		'5': "5",
		'6': "6",
		'7': "7",
		'8': "8",
		'9': "9",
	}

	for {
		drawScreen(cag, path)
		time.Sleep(50 * time.Millisecond)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD, termbox.KeyEnter:
				goto fin
			}

			switch ev.Ch {
			case 'h':
				cag = cag.MoveLeft()
			case 'j':
				cag = cag.MoveDown()
			case 'k':
				cag = cag.MoveUp()
			case 'l':
				cag = cag.MoveRight()
			default:
				v, ok := chMap[ev.Ch]
				if ok {
					if cag.CurrentPositionCellIsEnterable() {
						cag = cag.SetString(v)
					}
				}
			}
		}
	}
fin:

	return cag.Evaluate()
}

func drawScreen(cag column.ColumnAdditionGame, path string) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const leftPad = 1
	const topPad = 1

	cursorLineY := topPad
	drawLine(fmt.Sprintf("%s < remove file '%s'?", face, path), leftPad, topPad)
	cursorLineY += 2

	startLineY := cursorLineY
	for y, line := range cag.ResultStringLines() {
		drawLine(line, leftPad, y+startLineY)
		cursorLineY++
	}

	cursorLineY++
	drawLine("h: move left, j: move down, k: move up, l: move right", leftPad, cursorLineY)
	cursorLineY++
	drawLine("ENTER: confirm", leftPad, cursorLineY)

	x, y := cag.PositionXY()
	var y2 int
	if 0 <= y && y < 2 {
		y2 = y + 5 + topPad
	} else {
		y2 = y + 6 + topPad
	}
	termbox.SetCell(x*2+leftPad, y2, cag.CurrentPositionCellValueRune(), termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}

func drawLine(text string, x, y int) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}
