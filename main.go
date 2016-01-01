package main

import (
    "log"
    "math/rand"
    "time"
)

func main() {
    log.Print("started.")

    // チャネル
    sleep1_finished := make(chan bool)

    go func() {
        // 1秒かかるコマンド
        log.Print("sleep1 started.")
        time.Sleep(1 * time.Second)
        log.Print("sleep1 finished.")
        sleep1_finished <- true
    }()

    // 終わるまで待つ
    <- sleep1_finished

    dice()
}

func dice() {
    log.Printf("%v", random())
}

// http://qiita.com/cubicdaiya/items/819886c57e9d17e4b019
// 3までの数字をランダムに返す
func random()(int) {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(4)
}
