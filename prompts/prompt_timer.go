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
	i := rand.Intn(6) + 4
	msg := fmt.Sprintf("%s < stop in %d seconds [enter]", face, i)
	fmt.Println(msg)

	now := time.Now()

	running := true
	go printTimer(now, &running)
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	running = false
	aft := time.Now()

	st := aft.Add(-1 * time.Second).Sub(now)
	x := aft.Sub(now)
	et := aft.Add(1 * time.Second).Sub(now)

	return st <= x && x <= et, nil
}

func printTimer(t time.Time, running *bool) {
	for *running {
		duration := time.Now().Sub(t).Seconds()
		fmt.Println()
		fmt.Printf("%02.2f seconds", duration)
		time.Sleep(50 * time.Millisecond)
		fmt.Print("\x1b[1M")
	}
}
