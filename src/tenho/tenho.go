package tenho

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type OptionStruct struct {
	NoKokushi      bool
	NoChitoitsu    bool
	NoNormal       bool
	OutputPerTrial int
}

var option OptionStruct

func Start(o OptionStruct) {
	option = o

	start := time.Now().UnixNano()
	var i float64
	i = 0
	var hand Hand
	var ok bool
	for {
		i++
		seed := time.Now().UnixNano()
		hand, ok = tryOnce(seed)
		if ok || int(i)%option.OutputPerTrial == 0 {
			end := time.Now().UnixNano()
			diff := float64(end-start) / 1000000000
			m := i / diff
			out := 0
			if m >= 0 {
				out = int(m)
			}
			hai := hand.HaiString()
			fmt.Printf("\r%v回試行  %v秒経過 %v回/秒 %v", i, diff, out, hai)
		}
		if ok {
			break
		}
	}
	fmt.Printf("\n")
}

func tryOnce(seed int64) (Hand, bool) {
	hand := ShuffledHand(seed)
	ok := hand.Solve()
	return hand, ok
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

	hand := make([]int, MahjongSetSize, MahjongSetSize)
	copy(hand, GetMahjongSet())
	hand2 := make([]int, 0, 0)
	var j int

	for k := MahjongSetSize; k > MahjongSetSize-HandSize; k-- {
		j = rand.Intn(k)
		hand2 = append(hand2, hand[j])
		hand = append(hand[:j], hand[j+1:]...)
	}

	sort.Ints(hand2)

	retval := *new(Hand)
	for i := 0; i < HandSize; i++ {
		retval[i] = hand2[i]
	}
	return retval
}

type Hand [HandSize]int

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

// 国士無双判定
func (hand Hand) solveKokushi() bool {
	//あがりパターン列挙
	agaris := [13][14]int{
		[14]int{0, 0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 4, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 5, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 6, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 7, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 15, 15, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 16, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 24, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 25, 33},
		[14]int{0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33, 33},
	}
	for _, v := range agaris {
		if v == hand {
			return true
		}
	}
	return false
}

// あがり判定する
func (hand Hand) Solve() bool {
	return (!(option.NoKokushi) && hand.solveKokushi()) || (!(option.NoChitoitsu) && hand.solveChitoitsu()) || (!(option.NoNormal) && hand.GroupSuit().Solve())
}

type SuitGroup struct {
	innerSuitGroup
	color int
}

type innerSuitGroup []int

func NewSuitGroup(color int) *SuitGroup {
	s := SuitGroup{innerSuitGroup{}, color}
	return &s
}

type SuitsGroupedHand map[int]SuitGroup

const (
	Jihai = iota
	Manzu
	Sozu
	Pinzu
)

// スート分類
func (hand Hand) GroupSuit() SuitsGroupedHand {
	m := SuitsGroupedHand{
		Jihai: *NewSuitGroup(Jihai),
		Manzu: *NewSuitGroup(Manzu),
		Sozu:  *NewSuitGroup(Sozu),
		Pinzu: *NewSuitGroup(Pinzu),
	}
	for _, i := range hand {
		quo := (i - 7 + 9) / 9
		var mod int
		if i-7 >= 0 {
			mod = (i - 7) % 9
		} else {
			mod = i
		}
		s := m[quo]
		s.append(mod)
		m[quo] = s
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
		switch len(a.list()) % 3 {
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

func (a innerSuitGroup) list() innerSuitGroup {
	return a
}

func (a *innerSuitGroup) SetSuitGroup(b innerSuitGroup) {
	*a = b
}

func (a *innerSuitGroup) append(w int) {
	*a = append(*a, w)
}

// 33332形を形成するスートグループがどうかを判定
func (a SuitGroup) valid_suit_group(i int) bool {
	// 対子が含まれているスートグループがただ1つある前提

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
			rest := NewSuitGroup(i)
			c := 2
			for _, w := range a.list() {
				// ペア候補以外は新スライスに入れる
				// ペア候補は３枚目以降は新スライスに入れる
				if w != v || c <= 0 {
					rest.SetSuitGroup(append(rest.list(), w))
				}
				if w == v {
					c--
				}
			}
			if rest.valid_3cards() {
				return true
			} else {
				continue
			}
		}
		return false
	} else if len(a.list())%3 == 0 {
		return a.valid_3cards()
	}
	// 到達しないはず
	panic("到達しないはず")
}

func (a SuitGroup) valid_3cards() bool {
	// 刻子や順子のみで構成されている場合true
	// a is sorted
	// a.size % 3 is0
	for {
		if a.remove_kotsu() {
			continue
		}
		break
	}

	// 字牌でなければ順子チェック
	if a.color != Jihai {
		for {
			if a.remove_shuntsu() {
				continue
			}
			break
		}
	}
	return len(a.list()) == 0
}

func (a *innerSuitGroup) remove_kotsu() bool {
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

func (a *innerSuitGroup) remove_shuntsu() bool {
	// 順子を除去できればtrue
	// a is sorted
	rest := innerSuitGroup{}
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

func (a innerSuitGroup) pairable_numbers() innerSuitGroup {
	// a is sorted
	counter := []int{}
	var x, y int // 2つ前と1つ前
	for _, v := range a {
		if y == v && x != v {
			counter = append(counter, v)
		} else {
			y = v
		}
	}
	return counter
}
