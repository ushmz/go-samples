package main

import (
	"fmt"
	"time"
)

func main() {
	//合計実行数
	const totalExecuteNum = 6
	//同時実行数
	const maxConcurrencyNum = 3

	sig := make(chan string, maxConcurrencyNum)
	res := make(chan string, totalExecuteNum)
	// `main() が終了するまで処理が先延ばしにされる．`
	defer close(sig)
	defer close(res)

	fmt.Printf("start concurrency execute %s \n", time.Now())
	for i := 0; i < totalExecuteNum; i++ {
		go wait3Sec(sig, res, fmt.Sprintf("no%d", i))
	}
	for {
		//全部が終わるまで待つ（`res` に送信された件数を確認）
		if len(res) >= totalExecuteNum {
			break
		}
	}

	fmt.Printf("end concurrency execute %s \n", time.Now())
}

func wait3Sec(sig chan string, res chan string, name string) {
	// `sig` に値を送信
	sig <- fmt.Sprintf("sig %s", name)
	time.Sleep(3 * time.Second)
	fmt.Printf("%s:end wait 3sec \n", name)
	// `res` に値を送信
	res <- fmt.Sprintf("sig %s", name)
	// `sig` の値を受け取り
	<-sig
}
