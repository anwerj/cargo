package initial

import (
	"fmt"

	"github.com/anwerj/cargo/head"
)

// Main of Initial
func Main(a head.Aid) {
	fmt.Printf("hello, world\n")
	a.L.Info("Hello", []string{"World"})
}
