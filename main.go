package main

import (
    "log"
    "time"
)

func main() {
    a := 1
    log.Print("started.")

    // チャネル
    sleep1_finished := make(chan bool)

    go func() {
        // 1秒かかるコマンド
        log.Print("sleep1 started.")
        time.Sleep(1 * time.Second)
        log.Print("sleep1 finished.")
        a++
        log.Printf("a : %v", a)
        sleep1_finished <- true
    }()

    // 終わるまで待つ
    <- sleep1_finished

    log.Print("all finished.")
}
