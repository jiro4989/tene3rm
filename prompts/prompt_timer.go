package prompts

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

// promptWithTimer は一定時間経過後に入力を促すプロンプトを表示する。
func promptWithTimer(_ string) (bool, error) {
	i := rand.Intn(4) + 4
	msg := fmt.Sprintf("%s < stop in %d seconds [enter]", face, i)
	fmt.Println(msg)

	now := time.Now()
	running := true

	// タイマーの端末上に表示
	go printTimer(now, &running)

	// タイマーを手動解除するまでの入力待ち
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	// タイマーをストップする
	running = false

	// 基準となる時間の前後 1 秒を有効時間とする
	aft := time.Now()
	st := aft.Add(-1 * time.Second).Sub(now)
	et := aft.Add(1 * time.Second).Sub(now)

	// 乱数を秒数に変換
	x := time.Duration(i) * time.Second

	return st <= x && x <= et, nil
}

func printTimer(t time.Time, running *bool) {
	// 行頭に移動するエスケープシーケンス
	const escMove = "\r"

	for *running {
		duration := time.Now().Sub(t).Seconds()
		fmt.Printf("%02.2f seconds", duration)
		time.Sleep(50 * time.Millisecond)
		fmt.Print(escMove)
	}
}
