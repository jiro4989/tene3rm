package terminal

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/jiro4989/tene3rm/domain/model"
	"github.com/jiro4989/tene3rm/domain/service"
	"github.com/manifoldco/promptui"
	"github.com/nsf/termbox-go"
)

// promptWithMath は単純な算数入力を求めるプロンプトを表示する。
func promptWithMath(path string) (bool, error) {
	validate := func(input string) error {
		_, err := model.NewUserInputText(input).ToInt()
		if err != nil {
			return err
		}
		return nil
	}

	svc := service.NewMathService()
	seed := time.Now().Unix()
	r := rand.New(rand.NewSource(seed))
	want, a, b, op := svc.SimpleOperations(r, r, r)

	p := promptui.Prompt{
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a, op, b),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	resultNum, err := model.NewUserInputText(result).ToInt()
	if err != nil {
		return false, err
	}

	return want == resultNum, nil
}

var vals = [5][]string{
	{" ", " ", " ", " "},
	{" ", " ", " ", " "},
	{" ", " ", " ", " "},
	{" ", " ", " ", " "},
	{" ", " ", " ", " "},
}

// promptWithMath2 は筆算での計算結果を求めるプロンプトを表示する。
func promptWithMath2(path string) (bool, error) {
	a := rand.Intn(90) + 10
	b := rand.Intn(90) + 10

	cag, err := model.NewColumnAdditionGame(a, b)
	if err != nil {
		return false, err
	}

	if err := termbox.Init(); err != nil {
		return false, err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	for {
		drawScreen(cag)
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
			case '0':
				cag = cag.SetString("0")
			case '1':
				cag = cag.SetString("1")
			case '2':
				cag = cag.SetString("2")
			case '3':
				cag = cag.SetString("3")
			case '4':
				cag = cag.SetString("4")
			case '5':
				cag = cag.SetString("5")
			case '6':
				cag = cag.SetString("6")
			case '7':
				cag = cag.SetString("7")
			case '8':
				cag = cag.SetString("8")
			case '9':
				cag = cag.SetString("9")
			}
		}
	}
fin:

	return cag.Evaluate()
}

func setNum(pos model.Position, s string) {
	vals[pos.Y()][pos.X()] = s
}

func drawScreen(cag model.ColumnAdditionGame) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const leftPad = 1
	const topPad = 1
	var lastLineY int
	for y, line := range cag.ResultStringLines() {
		drawLine(line, leftPad, y+topPad)
		lastLineY = y
	}
	drawLine("h: move left, j: move down, k: move up, l: move right", leftPad, lastLineY+1)
	drawLine("ENTER: confirm", leftPad, lastLineY+2)

	x, y := cag.PositionXY()
	var y2 int
	if 2 <= y && y < 4 {
		y2 = y + 1 + topPad
	} else {
		y2 = y + 2 + topPad
	}
	termbox.SetCell(x*2+leftPad, y2, cag.CurrentPositionCellValueRune(), termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}

func drawLine(text string, x, y int) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, termbox.ColorDefault, termbox.ColorDefault)
	}
}
