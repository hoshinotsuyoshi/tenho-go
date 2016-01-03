package tenho_test

import (
	"fmt"
	"tenho"
)

func ExampleTry_once() {
	var seed int64
	seed = 1451836284287681922
	fmt.Println(tenho.Try_once(seed))
	// Output:
	// 🀟 🀡 🀙 🀋 🀌 🀜 🀊 🀘 🀟 🀠 🀗 🀕 🀈 🀛 false
}
