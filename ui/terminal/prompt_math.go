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
		_, err := model.NewNumberWithText(model.NewText(input))
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

	resultNum, err := model.NewNumberWithText(model.NewText(result))
	if err != nil {
		return false, err
	}

	return want.Equal(resultNum), nil
}

// promptWithMath2 は筆算での計算結果を求めるプロンプトを表示する。
func promptWithMath2(path string) (bool, error) {
	a := rand.Intn(90) + 10
	b := rand.Intn(90) + 10

	if err := termbox.Init(); err != nil {
		return false, err
	}
	defer termbox.Close()
	termbox.SetInputMode(termbox.InputEsc)
	termbox.Flush()

	drawBackground(a, b)

	time.Sleep(3 * time.Second)

	return false, nil
}

func drawBackground(a, b int) {
	termbox.Clear(termbox.ColorDefault, termbox.ColorDefault)

	vals := [5][]string{
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
		{" ", " ", " ", " "},
	}

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

	drawLine(strsToLine(vals[0]), 1, 1)
	drawLine(strsToLine(vals[1]), 1, 2)
	drawLine(horizontalLine, 1, 3)
	drawLine(strsToLine(vals[2]), 1, 4)
	drawLine(strsToLine(vals[3]), 1, 5)
	drawLine(horizontalLine, 1, 6)
	drawLine(strsToLine(vals[4]), 1, 7)

	drawLine("h: move left, j: move down, k: move up, l: move right", 1, 9)
	drawLine("ENTER: confirm", 1, 9)

	termbox.Flush()
}

const fg = termbox.ColorDefault
const bg = termbox.ColorDefault

//	  2 5
//	x 5 5
//
// -------
//
//	1 2 5
//
// 1 2 5
// -------
// 1 3 7 5
func drawLine(text string, x, y int) {
	for i, c := range text {
		termbox.SetCell(x+i, y, c, fg, bg)
	}
}

func strsToLine(vals []string) string {
	return strings.Join(vals, " ")
}

const horizontalLine = "-------"
