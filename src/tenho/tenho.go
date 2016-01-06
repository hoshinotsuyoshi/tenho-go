package tenho

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

func Start() {
	start := time.Now().UnixNano()
	// 処理
	var i float64
	i = 0
	for {
		i++
		end := time.Now().UnixNano()
		diff := float64(end-start) / 1000000000
		m := i / diff
		out := 0
		if m >= 0 {
			out = int(m)
		}
		seed := time.Now().UnixNano()
		var hai string
		var ok bool
		hai, ok = tryOnce(seed)
		if int(i)%10000 == 0 {
			fmt.Printf("\r%v回試行  %v秒経過 %v回/秒 %v", i, diff, out, hai)
		}
		if ok {
			fmt.Printf("\r%v回試行  %v秒経過 %v回/秒 %v", i, diff, out, hai)
			break
		}
		//fmt.Printf("\r%v回試行  %v秒経過 %v回/秒 %v", i, diff, out, hai)
		//if i > 100000 {
		//	break
		//}
	}
	fmt.Printf("\n")
}

func tryOnce(seed int64) (string, bool) {
	hand := ShuffledHand(seed)
	hai := hand.HaiString()
	ok := hand.Solve()
	return hai, ok
}

// http://d.hatena.ne.jp/hake/20150930/p1
func shuffle(hand Hand) {
	for i := len(hand); i > 1; i-- {
		j := rand.Intn(i) // 0 .. i-1 の乱数発生
		hand[i-1], hand[j] = hand[j], hand[i-1]
	}
}

const HandSize = 14
const MahjongSetSize = 136

var defaultSet []int

func GetMahjongSet() []int {
	if defaultSet == nil {
		size := MahjongSetSize
		defaultSet = make([]int, size, size)
		for i := 0; i < size; i++ {
			defaultSet[i] = i / 4
		}
	}
	return defaultSet
}

func ShuffledHand(seed int64) Hand {
	rand.Seed(seed)

	hand := make(Hand, MahjongSetSize, MahjongSetSize)
	copy(hand, GetMahjongSet())
	hand2 := make(Hand, 0, 0)
	var j int

	for k := MahjongSetSize; k > MahjongSetSize-HandSize; k-- {
		j = rand.Intn(k)
		hand2 = append(hand2, hand[j])
		hand = append(hand[:j], hand[j+1:]...)
	}

	return hand2
}

type Hand []int

// 牌文字への変換(スペース区切り)
func (hand Hand) HaiString() string {
	// http://qiita.com/ruiu/items/2bb83b29baeae2433a79
	// サイズ0、内部バッファの長さ69の[]byteの値を割り当てる
	b := make([]byte, 0, 70)

	// bに文字列を追加
	for j := 0; j < HandSize; j++ {
		// コードポイント上、普通の麻雀牌はU+1F000からの34個。
		// U+1F000 is 'MAHJONG TILE EAST WIND' ('東')
		// https://codepoints.net/U+1F000
		b = append(b, string(hand[j]+0x1F000)...) // appendするには...が必要
		// 自分のMacではスペース区切りでないとうまく表示されないためスペースを挿入する
		// U+0020 is 'SPACE'
		// https://codepoints.net/U+0020
		b = append(b, string(0x20)...) // appendするには...が必要
	}
	return string(b)
}

// 七対子判定
func (hand Hand) solveChitoitsu() bool {
	//カウンタ
	c := map[int]int{}

	for _, v := range hand {
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

// あがり判定する
func (hand Hand) Solve() bool {
	return hand.solveChitoitsu() || hand.GroupSuit().Solve()
}

type SuitGroup []int

type SuitsGroupedHand map[int]SuitGroup

const (
	Jihai = iota
	Manzu
	Sozu
	Pinzu
)

// スート分類
func (hand Hand) GroupSuit() SuitsGroupedHand {
	m := SuitsGroupedHand{}
	for _, i := range hand {
		quo := (i - 7 + 9) / 9
		var mod int
		if i-7 >= 0 {
			mod = (i - 7) % 9
		} else {
			mod = i
		}
		m[quo] = append(m[quo], mod)
	}
	return m
}

func (m SuitsGroupedHand) Solve() bool {
	return m.a_pair_existible() && m.valid_33332()
}

func (m SuitsGroupedHand) a_pair_existible() bool {
	//スートのサイズを3で割った時
	//あまりが2であるスートグループが1つであること
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

func (m SuitsGroupedHand) valid_33332() bool {
	for i := 0; i < 4; i++ {
		if !m[i].valid_suit_group(i) {
			return false
		}
	}
	return true
}

func (a SuitGroup) list() SuitGroup {
	return a
}

// 33332形を形成するスートグループがどうかを判定
func (a SuitGroup) valid_suit_group(i int) bool {
	// 対子が含まれているスートグループがただ1つある前提

	//ソート
	sort.Ints(a.list())
	if len(a.list())%3 == 2 {
		//ペアを探す
		pair_numbers := a.pairable_numbers()
		//ペア候補がなかったらぬける
		if len(pair_numbers) == 0 {
			return false
		}
		//ペア候補毎に繰り返し処理
		for _, v := range pair_numbers {
			//ペアとなる２枚を除去
			rest := SuitGroup{}
			c := 2
			for _, w := range a.list() {
				// ペア候補以外は新スライスに入れる
				// ペア候補は３枚目以降は新スライスに入れる
				if w != v || c <= 0 {
					rest = append(rest, w)
				}
				if w == v {
					c--
				}
			}
			if rest.valid_3cards(i) {
				return true
			} else {
				continue
			}
		}
		return false
	} else if len(a.list())%3 == 0 {
		return a.valid_3cards(i)
	}
	// 到達しないはず
	panic("到達しないはず")
}

func (a SuitGroup) valid_3cards(i int) bool {
	// 刻子や順子のみで構成されている場合true
	// a is sorted
	// a.size % 3 is0
	// 引数は字牌のとき0
	for {
		if a.remove_kotsu() {
			continue
		}
		if i > 0 {
			if a.remove_shuntsu() {
				continue
			}
		}
		return len(a) == 0
	}
}

func (a *SuitGroup) remove_kotsu() bool {
	// 刻子を除去できればtrue
	// a is sorted
	x := *a
	if len(x) < 3 {
		return false
	}
	if x[0] == x[1] && x[0] == x[2] {
		*a = x[3:]
		return true
	}
	return false
}

func (a *SuitGroup) remove_shuntsu() bool {
	// 順子を除去できればtrue
	// a is sorted
	rest := SuitGroup{}
	first := -1
	second := -1
	found := false
	for _, v := range *a {
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
	*a = rest
	return found
}

func (a SuitGroup) pairable_numbers() SuitGroup {
	// a is sorted
	counter := []int{}
	x := 999 // 2つ前
	y := 999 // 1つ前
	for _, v := range a {
		if y == v && x != v {
			counter = append(counter, v)
		} else {
			y = v
		}
	}
	return counter
}
