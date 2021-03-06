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
	group := hand.GroupSuit()
	fmt.Println(group[0])
	fmt.Println(group[1])
	fmt.Println(group[2])
	fmt.Println(group[3])
	// Output:
	// {[] 0}
	// {[4 5 3 1] 1}
	// {[8 7 5] 2}
	// {[6 8 0 3 6 7 2] 3}
}

/*
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
*/

func ExampleShuffledHand() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.ShuffledHand(seed))
	// Output:
	// [0 3 5 9 10 11 12 14 18 23 28 29 30 30]
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

func ExampleHand_Solve_kokushimuso() {
	list := tenho.Hand{0, 0, 1, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33}
	fmt.Println(list.Solve())
	// Output:
	// true
}

func ExampleHand_Solve_kokushimuso2() {
	list := tenho.Hand{0, 0, 0, 2, 3, 4, 5, 6, 7, 15, 16, 24, 25, 33}
	fmt.Println(list.Solve())
	// Output:
	// false
}
