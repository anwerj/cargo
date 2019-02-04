package god

import (
	"fmt"
)

const welcomeMessage = "Yes my child! You have reached to the "

// Hand holds everything needed in project, for all dependencies just talk to hand
type Hand struct {
	O Output
	I Input
	U Utils
	S System
	L Logger
}

// Main func for Hand, Inly Greets
func Main(h Hand) {
	fmt.Print(welcomeMessage)
	h.O.Red("\"God\"")
}

// Call will return in God's Hand
func Call() Hand {
	return Hand{
		O: NewOutput(), I: NewInput(), U: NewUtils(), S: NewSystem(), L:NewLogger()}
}

func LogInfo(a... interface{}) {
	l := NewLogger()
	l.SetSkip(2).Info(a...)
}
