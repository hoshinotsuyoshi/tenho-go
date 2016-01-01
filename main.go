package main

import (
	"log"
	"math/rand"
	"time"
)

func main() {
	aa := []int{101, 100, 70, 70, 60, 60, 50, 50, 40, 40, 30, 30, 20, 20}
	if is_chitoitsu(aa) {
		log.Println("a")
	} else {
		log.Println("b")
	}
	return

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

	return list[:14]
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
	for _, value := range list {
		group(matrix, value)
	}
	log.Println(matrix)
}

// スート分類してくれる
func group(m [][]int, j int) {
	// 4で割ると本来のインデックスに
	i := j / 4
	switch {
	case i < 7:
		m[0] = append(m[0], i)
		log.Println(i)
	case i < 7+(9*1):
		m[1] = append(m[1], i-7)
		log.Println(i)
	case i < 7+(9*2):
		m[2] = append(m[2], i-7-(9*1))
		log.Println(i)
	case i < 7+(9*3):
		m[3] = append(m[3], i-7-(9*2))
		log.Println(i)
	}
}

func is_chitoitsu(list []int) bool {
	//カウンタ
	c := map[int]int{}

	//コピー
	l := list

	for _, v := range l {
		count, ok := c[v]
		if ok {
			if count == 1 {
				c[v] = 2
			} else {
				// c[v] == 2
				return false
			}
		} else {
			c[v] = 1
		}
		//8個チェック
		if len(c) >= 8 {
			return false
		}
	}
	return true
}
