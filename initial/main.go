package initial

import (
	"fmt"

	"github.com/anwerj/MyGod/god"
)

// Main of Initial
func Main(h god.Hand) {
	fmt.Printf("hello, world\n")
	h.L.Info("Hello", []string{"World"})
}
