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
        // 0.2秒かかるコマンド
        log.Print("sleep1 started.")
        time.Sleep(200 * 1000 * 1000 * time.Nanosecond)
        log.Print("sleep1 finished.")
        sleep1_finished <- true
    }()

    // 終わるまで待つ
    <- sleep1_finished

    dice()
}

func dice() {
    log.Printf("%v", string(random()))
}

// http://qiita.com/cubicdaiya/items/819886c57e9d17e4b019
// 34枚の内からランダムに1枚返す
func random()(int) {
    a := 126976 // ton https://codepoints.net/U+1F000
    rand.Seed(time.Now().UnixNano())
    return a + rand.Intn(34)
}
