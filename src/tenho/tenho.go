package tenho

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Start() {
	start := time.Now()
	// 処理
	var i float64
	i = 0
	for {
		i++
		end := time.Now()
		diff := float64(end.Sub(start) / 1000 / 1000 / 1000)
		m := int(i / diff)
		out := 0
		if m >= 0 {
			out = m
		}
		fmt.Printf("\r")
		fmt.Printf("\r%v回試行  %v秒経過 %v回/秒", i, diff, out)
		seed := time.Now().UnixNano()
		var hai string
		var ok bool
		hai, ok = tryOnce(seed)
		fmt.Printf("%v", hai)
		if ok {
			break
		}
	}
	fmt.Printf("\n")
}

type Cards []int

func tryOnce(seed int64) (string, bool) {
	list := ShuffledCards(seed)
	hai := HaiString(list)
	ok := Solve(list)
	return hai, ok
}

// http://d.hatena.ne.jp/hake/20150930/p1
func shuffle(list Cards) {
	for i := len(list); i > 1; i-- {
		j := rand.Intn(i) // 0 .. i-1 の乱数発生
		list[i-1], list[j] = list[j], list[i-1]
	}
}

func ShuffledCards(seed int64) Cards {
	rand.Seed(seed)

	// データ要素数指定、および初期データ作成
	size := 136
	list := make(Cards, size, size)
	for i := 0; i < size; i++ {
		list[i] = i / 4
	}

	// シャッフル
	shuffle(list)

	return list[:14]
}

// 牌文字への変換(スペース区切り)
func HaiString(list Cards) string {
	// http://qiita.com/ruiu/items/2bb83b29baeae2433a79
	// サイズ0、内部バッファの長さ69の[]byteの値を割り当てる
	b := make([]byte, 0, 70)

	// bに文字列を追加
	for j := 0; j < 14; j++ {
		// コードポイント上、普通の麻雀牌はU+1F000からの34個。
		// U+1F000 is 'MAHJONG TILE EAST WIND' ('東')
		// https://codepoints.net/U+1F000
		b = append(b, string(list[j]+0x1F000)...) // appendするには...が必要
		// 自分のMacではスペース区切りでないとうまく表示されないためスペースを挿入する
		// U+0020 is 'SPACE'
		// https://codepoints.net/U+0020
		b = append(b, string(0x20)...) // appendするには...が必要
	}
	return string(b)
}

// あがり判定する
func Solve(list Cards) bool {
	if is_chitoitsu(list) {
		return true
	}
	matrix := []Cards{{}, {}, {}, {}}
	for _, value := range list {
		group(matrix, value)
	}
	//fmt.Printf("%v", matrix)

	return group_scan(matrix)
}

// スート分類してくれる
func group(m []Cards, i int) {
	switch {
	case i < 7:
		m[0] = append(m[0], i)
	case i < 7+(9*1):
		m[1] = append(m[1], i-7)
	case i < 7+(9*2):
		m[2] = append(m[2], i-7-(9*1))
	case i < 7+(9*3):
		m[3] = append(m[3], i-7-(9*2))
	}
}

func group_scan(m []Cards) bool {
	if !valid_mod3(m) {
		return false
	}
	if !valid_33332(m) {
		return false
	}
	return true
}

func valid_mod3(m []Cards) bool {
	//スートのサイズを3で割った時
	//あまりが2であるすーとグループが1つであること
	c := 0
	for _, a := range m {
		switch len(a) % 3 {
		case 0:
			// noop
		case 1:
			return false
		case 2:
			c++
		}
	}
	return c == 1
}

func valid_33332(m []Cards) bool {
	for i := 0; i < 4; i++ {
		if !valid_suit_group(m[i], i) {
			return false
		}
	}
	return true
}

func valid_suit_group(a Cards, i int) bool {
	// 第二引数は字牌のとき0

	//ソート
	sort.Ints(a)
	if len(a)%3 == 2 {
		//ペアを探す
		pair_numbers := pairable_numbers(a)
		//ペア候補がなかったらぬける
		if len(pair_numbers) == 0 {
			return false
		}
		//ペア候補毎に繰り返し処理
		for _, v := range pair_numbers {
			//ペアとなる２枚を除去
			rest := Cards{}
			c := 2
			for _, w := range a {
				// ペア候補以外は新スライスに入れる
				// ペア候補は３枚目以降は新スライスに入れる
				if w != v || c <= 0 {
					rest = append(rest, w)
				}
				if w == v {
					c--
				}
			}
			if valid_3cards(rest, i) {
				return true
			} else {
				continue
			}
		}
		return false
	} else if len(a)%3 == 0 {
		return valid_3cards(a, i)
	}
	// 到達しないはず
	panic("到達しないはず")
}

func valid_3cards(a Cards, i int) bool {
	// 刻子や順子のみで構成されている場合true
	// a is sorted
	// a.size % 3 is0
	// 第二引数は字牌のとき0
	ok := false
	for {
		a, ok = remove_kotsu(a)
		if ok {
			continue
		}
		if i > 0 {
			a, ok = remove_shuntsu(a)
			if ok {
				continue
			}
		}
		return len(a) == 0
	}
}

func remove_kotsu(a Cards) (Cards, bool) {
	// 刻子を除去できればtrue
	// a is sorted
	retval := a
	if len(a) < 3 {
		return retval, false
	}
	if a[0] == a[1] && a[0] == a[2] {
		retval = a[3:]
		return retval, true
	}
	return retval, false
}

func remove_shuntsu(a Cards) (Cards, bool) {
	// 順子を除去できればtrue
	// a is sorted
	rest := Cards{}
	first := -1
	second := -1
	found := false
	for _, v := range a {
		if found {
			rest = append(rest, v)
			continue
		}
		if first == -1 {
			first = v
		} else if second == -1 && first+1 == v {
			second = v
		} else if second != -1 && first+2 == v {
			//flush
			first = -1
			second = -1
			found = true
		} else {
			rest = append(rest, v)
		}
	}
	return rest, found
}

func pairable_numbers(sorted Cards) Cards {
	//カウンタ
	retval := Cards{}
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

func is_chitoitsu(list Cards) bool {
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
