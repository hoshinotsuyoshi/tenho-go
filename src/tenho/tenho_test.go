package tenho_test

import (
	"fmt"
	"tenho"
)

func ExampleTryOnce() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.TryOnce(seed))
	// Output:
	// 🀟 🀡 🀙 🀋 🀌 🀜 🀊 🀘 🀟 🀠 🀗 🀕 🀈 🀛 false
}

func ExampleShuffledCards() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.ShuffledCards(seed))
	// Output:
	// [31 33 25 11 12 28 10 24 31 32 23 21 8 27]
}

func ExampleSolve_false() {
	list := []int{31, 33, 25, 11, 12, 28, 10, 24, 31, 32, 23, 21, 8, 27}
	fmt.Println(tenho.Solve(list))
	// Output:
	// false
}
