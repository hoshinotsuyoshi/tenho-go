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

func ExampleHand_GroupSuit() {
	hand := tenho.Hand{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(hand.GroupSuit())
	// Output:
	// [[] [4 5 3 1] [8 7 5] [6 8 0 3 6 7 2]]
}

func ExampleSuitsGroupedHand_Solve_false() {
	hand := tenho.SuitsGroupedHand{{}, {4, 5, 3, 1}, {8, 7, 5}, {6, 8, 0, 3, 6, 7, 2}}
	fmt.Println(hand.Solve())
	// Output:
	// false
}

func ExampleSuitsGroupedHand_Solve_true() {
	hand := tenho.SuitsGroupedHand{{2, 2}, {3, 4, 5}, {7, 7, 7}, {3, 2, 1, 9, 8, 7}}
	fmt.Println(hand.Solve())
	// Output:
	// true
}

func ExampleShuffledHand() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.ShuffledHand(seed))
	// Output:
	// [31 33 25 11 12 28 10 24 31 32 23 21 8 27]
}

func ExampleSolve_false() {
	list := tenho.Hand{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(tenho.Solve(list))
	// Output:
	// false
}

func ExampleSolve_true() {
	list := tenho.Hand{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 16, 16, 17, 17}
	fmt.Println(tenho.Solve(list))
	// Output:
	// true
}

func ExampleSolve_chitoitsu() {
	list := tenho.Hand{0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}
	fmt.Println(tenho.Solve(list))
	// Output:
	// true
}
