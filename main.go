package main

import (
	"log"
	"math/rand"
	"sort"
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
func solve(list []int) bool {
	if is_chitoitsu(list) {
		return true
	}
	matrix := [][]int{{}, {}, {}, {}}
	for _, value := range list {
		group(matrix, value)
	}
	log.Println(matrix)

	group_scan(matrix)
	return false
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

func group_scan(m [][]int) bool {
	if !validate_mod3(m) {
		return false
	}
	if !validate_33332(m) {
		return false
	}
	return true
}

func validate_mod3(m [][]int) bool {
	//スートのサイズを3で割った時
	//あまりが2であるすーとグループが1つであること
	c := 0
	for _, a := range m {
		switch len(a) % 3 {
		case 0:
			// noop
		case 1:
			log.Println("immi")
			return false
		case 2:
			c++
		}
	}
	return c == 1
}

func validate_33332(m [][]int) bool {
	for _, a := range m {
		if !validate_suit_group(a) {
			return false
		}
	}
	return true
}

func validate_suit_group(a []int) bool {
	//TODO: implement
	for _, v := range a {
		// 4で割ると本来のインデックスに
		v = v / 4
	}

	//ソート
	sort.Ints(a)
	if len(a)%3 == 2 {
		//ペアを探す
		pair_numbers := pairable_numbers(a)
		//ペア候補がなかったらぬける
		if len(pair_numbers) == 0 {
			return false
		}
	} else if len(a)%3 == 0 {
		return true
	}
	// 来ないはず
	return false
}

func pairable_numbers(sorted []int) []int {
	//カウンタ
	retval := []int{}
	a := 999 // 2つ前
	b := 999 // 1つ前
	for _, v := range sorted {
		if b == v && a != v {
			retval = append(retval, v)
		} else {
			b = v
		}
	}
	return retval
}

func is_chitoitsu(list []int) bool {
	//カウンタ
	c := map[int]int{}

	//コピー
	l := list

	for _, v := range l {
		// 4で割ると本来のインデックスに
		v = v / 4
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
