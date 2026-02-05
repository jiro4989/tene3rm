package terminal

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/jiro4989/tene3rm/domain/model"
	"github.com/jiro4989/tene3rm/domain/service"
	"github.com/manifoldco/promptui"
	"github.com/nsf/termbox-go"
)

// promptWithMath は単純な算数入力を求めるプロンプトを表示する。
func promptWithMath(path string) (bool, error) {
	validate := func(input string) error {
		_, err := model.NewNumberWithUserInputText(model.NewUserInputText(input))
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
		Label:    fmt.Sprintf("%s: remove file '%s'? (%d %s %d = ?)", appname, path, a.Value(), op, b.Value()),
		Validate: validate,
	}
	result, err := p.Run()
	if err != nil {
		return false, err
	}

	resultNum, err := model.NewNumberWithUserInputText(model.NewUserInputText(result))
	if err != nil {
		return false, err
	}

	return want.Equal(resultNum), nil
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

	var (
		a10 int = a / 10
		a1  int = a % 10
		b10 int = b / 10
		b1  int = b % 10
	)

	vals[0][2] = fmt.Sprintf("%d", a10)
	vals[0][3] = fmt.Sprintf("%d", a1)
	vals[1][1] = "x"
	vals[1][2] = fmt.Sprintf("%d", b10)
	vals[1][3] = fmt.Sprintf("%d", b1)

	if err := termbox.Init(); err != nil {
		return false, err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	pos := model.NewPosition(3, 2, 0, 2, 3, 4)
	for {
		drawBackground(pos.X(), pos.Y())
		time.Sleep(50 * time.Millisecond)
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case termbox.KeyCtrlC, termbox.KeyCtrlD, termbox.KeyEnter:
				goto fin
			}

			switch ev.Ch {
			case 'h':
				pos = pos.MoveLeft()
			case 'j':
				pos = pos.MoveDown()
			case 'k':
				pos = pos.MoveUp()
			case 'l':
				pos = pos.MoveRight()
			case '0':
				setNum(pos, "0")
			case '1':
				setNum(pos, "1")
			case '2':
				setNum(pos, "2")
			case '3':
				setNum(pos, "3")
			case '4':
				setNum(pos, "4")
			case '5':
				setNum(pos, "5")
			case '6':
				setNum(pos, "6")
			case '7':
				setNum(pos, "7")
			case '8':
				setNum(pos, "8")
			case '9':
				setNum(pos, "9")
			}
		}
	}
fin:

	n1, err := model.NewNumberWithUserInputText(model.NewUserInputText(strings.Join(vals[2], "")))
	if err != nil {
		return false, err
	}
	n2, err := model.NewNumberWithUserInputText(model.NewUserInputText(strings.Join(vals[3], "")))
	if err != nil {
		return false, err
	}
	n3, err := model.NewNumberWithUserInputText(model.NewUserInputText(strings.Join(vals[4], "")))
	if err != nil {
		return false, err
	}

	r1 := n1.Equal(model.NewNumber(a * b1))
	r2 := n2.Equal(model.NewNumber(a * b10))
	r3 := n3.Equal(model.NewNumber(a * b))

	return r1 && r2 && r3, nil
}

func setNum(pos model.Position, s string) {
	vals[pos.Y()][pos.X()] = s
}

func drawBackground(x, y int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	const leftPad = 1
	const topPad = 1

	drawLine(strsToLine(vals[0]), leftPad, 1)
	drawLine(strsToLine(vals[1]), leftPad, 2)
	drawLine(horizontalLine, leftPad, 3)
	drawLine(strsToLine(vals[2]), leftPad, 4)
	drawLine(strsToLine(vals[3]), leftPad, 5)
	drawLine(horizontalLine, leftPad, 6)
	drawLine(strsToLine(vals[4]), leftPad, 7)

	drawLine("h: move left, j: move down, k: move up, l: move right", 1, 9)
	drawLine("ENTER: confirm", 1, 10)

	var y2 int
	if 2 <= y && y < 4 {
		y2 = y + 1 + topPad
	} else {
		y2 = y + 2 + topPad
	}
	termbox.SetCell(x*2+leftPad, y2, []rune(vals[y][x])[0], termbox.ColorWhite, termbox.ColorBlack)
	termbox.Flush()
}

const fg = termbox.ColorDefault
const bg = termbox.ColorDefault

func drawLine(text string, x, y int) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}

func strsToLine(vals []string) string {
	return strings.Join(vals, " ")
}

const horizontalLine = "-------"
