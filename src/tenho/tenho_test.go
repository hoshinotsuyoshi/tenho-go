package tenho_test

import (
	"fmt"
	"github.com/hoshinotsuyoshi/tenho-go/src/tenho"
)

func ExampleHand_HaiString() {
	hand := tenho.Hand{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(hand.HaiString())
	// Output:
	// 🀟 🀡 🀙 🀋 🀌 🀜 🀊 🀘 🀟 🀠 🀗 🀕 🀈 🀛
}

/*
mapのキーの順序をちゃんとするのが面倒
func ExampleHand_GroupSuit() {
	hand := tenho.Hand{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(hand.GroupSuit())
	// Output:
	// map[0:[] 1:[4 5 3 1] 2:[8 7 5] 3:[6 8 0 3 6 7 2]]
}
*/

func ExampleSuitsGroupedHand_Solve_false() {
	hand := tenho.SuitsGroupedHand{
		tenho.Jihai: {},
		tenho.Manzu: {4, 5, 3, 1},
		tenho.Sozu:  {8, 7, 5},
		tenho.Pinzu: {6, 8, 0, 3, 6, 7, 2},
	}
	fmt.Println(hand.Solve())
	// Output:
	// false
}

func ExampleSuitsGroupedHand_Solve_true() {
	hand := tenho.SuitsGroupedHand{
		tenho.Jihai: {2, 2},
		tenho.Manzu: {3, 4, 5},
		tenho.Sozu:  {7, 7, 7},
		tenho.Pinzu: {3, 2, 1, 9, 8, 7},
	}
	fmt.Println(hand.Solve())
	// Output:
	// true
}

func ExampleShuffledHand() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.ShuffledHand(seed))
	// Output:
	// [14 18 9 30 30 12 29 23 10 5 28 3 11 0]
}

func ExampleHand_Solve_false() {
	list := tenho.Hand{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(list.Solve())
	// Output:
	// false
}

func ExampleHand_Solve_true() {
	list := tenho.Hand{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 16, 16, 17, 17}
	fmt.Println(list.Solve())
	// Output:
	// true
}

func ExampleHand_Solve_chitoitsu() {
	list := tenho.Hand{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	fmt.Println(list.Solve())
	// Output:
	// true
}
