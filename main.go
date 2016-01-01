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

    // 136枚の中から14枚返したい
    //dice()
    sub()
}

func dice() {
    log.Printf("%v", rand_hai())
}

func rand_hai()(string){
    a := 126976 // ton https://codepoints.net/U+1F000
    return string(a + dice34())
}

func rand_hai_14()(string){
    a := 126976 // ton https://codepoints.net/U+1F000
    return string(a + dice34())
}

// http://qiita.com/cubicdaiya/items/819886c57e9d17e4b019
// 0-33の内からランダムに1つ返す
func dice34()(int) {
    rand.Seed(time.Now().UnixNano())
    return rand.Intn(34)
}


// http://d.hatena.ne.jp/hake/20150930/p1
func shuffle(list []int){
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i)          // 0　.. i-1 の乱数発生
		list[i - 1], list[j] = list[j], list[i - 1]
	}
}

func sub() {
	rand.Seed(time.Now().UnixNano())

	// データ要素数指定、および初期データ作成
	size := 136
	list := make([]int, size, size)
	for i := 0; i < size; i++ { list[i] = i }

	// シャッフル
	shuffle(list)

	for j := 0; j < 14; j++ {
		log.Println(list[j])
	}
}
