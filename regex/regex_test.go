package regex

import (
	"regexp"
	"testing"
)

func BenchmarkContinuousCompile(b *testing.B) {
	for i := 0; i < b.N; i++ {
		compiled, err := regexp.Compile("^[a-zA-Z0-9]{14}$")
		if err != nil {
			panic(compiled)
		}
		compiled.Match([]byte("MatchMe"))
	}
}

func BenchmarkPreCompile(b *testing.B) {
	compiled, err := regexp.Compile("^[a-zA-Z0-9]{14}$")
	if err != nil {
		panic(compiled)
	}
	for i := 0; i < b.N; i++ {
		compiled.Match([]byte("MatchMe"))
	}
}
