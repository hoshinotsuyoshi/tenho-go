package main

import (
    "log"
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
    log.Print("all finished.")
}
