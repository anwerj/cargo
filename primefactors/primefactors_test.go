package primefactors

import (
	"fmt"
	"testing"
)

func benmarkRecursive(b *testing.B) {
	var listy []int
	fmt.Println("Started with ", b.N)
	for n := 0; n < b.N; n++ {
		recursive(n, listy)
	}
}
