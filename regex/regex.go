package regex

import "github.com/anwerj/cargo/head"

func Main(a head.Aid) {
	files(a)
}

func files(a head.Aid) (files []string) {
	// Random golang github repo files 
	files = []string{
		"/app/config/payments/upi/main.go",
		"/app/config/payments/upi/main.go",
	}
	return
}
