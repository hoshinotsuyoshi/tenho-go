package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {

	/*
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
	*/

	// 136枚の中から14枚返したい
	list := shuffled_cards()

	// 出力
	// digit_output(list)
	solve(list)
	string_output(list)
}

// http://d.hatena.ne.jp/hake/20150930/p1
func shuffle(list []int) {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0　.. i-1 の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}
}

func shuffled_cards() []int {
	rand.Seed(time.Now().UnixNano())

	// データ要素数指定、および初期データ作成
	size := 136
	list := make([]int, size, size)
	for i := 0; i < size; i++ {
		list[i] = i
	}

	// シャッフル
	shuffle(list)

	return list
}

// 単純出力
func digit_output(list []int) {
	for j := 0; j < 14; j++ {
		log.Println(list[j])
	}
}

// 牌の出力
func string_output(list []int) {
	// http://qiita.com/ruiu/items/2bb83b29baeae2433a79
	// サイズ0、内部バッファの長さ69の[]byteの値を割り当てる
	b := make([]byte, 0, 70)

	// bに文字列を追加
	for j := 0; j < 14; j++ {
		// 126976 is 'ton'
		// https://codepoints.net/U+1F000
		b = append(b, string(list[j]/4+126976)...) // ...が必要
		b = append(b, string(32)...)               // ...が必要
	}
	log.Print(string(b))
}

// solve
// 字マンソーピンのリストをつくる
func solve(list []int) {
	matrix := [][]int{{}, {}, {}, {}}
	a := 0
	switch {
	case a == 0:
		log.Println("a")
	}
	log.Println(matrix)
}
