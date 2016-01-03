package tenho_test

import (
	"fmt"
	"tenho"
)

func ExampleCards_HaiString() {
	cards := tenho.Cards{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(cards.HaiString())
	// Output:
	// 🀟 🀡 🀙 🀋 🀌 🀜 🀊 🀘 🀟 🀠 🀗 🀕 🀈 🀛
}

func ExampleShuffledCards() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.ShuffledCards(seed))
	// Output:
	// [31 33 25 11 12 28 10 24 31 32 23 21 8 27]
}

func ExampleSolve_false() {
	list := tenho.Cards{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(tenho.Solve(list))
	// Output:
	// false
}

func ExampleSolve_true() {
	list := tenho.Cards{7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 16, 16, 17, 17}
	fmt.Println(tenho.Solve(list))
	// Output:
	// true
}
