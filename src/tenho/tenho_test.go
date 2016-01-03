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
