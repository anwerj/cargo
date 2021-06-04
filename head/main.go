package head

import (
	"fmt"
)

const welcomeMessage = "Yes my child! You have reached to the "

// Aid holds everything needed in project, for all dependencies just talk to hand
type Aid struct {
	O Output
	I Input
	U Utils
	S System
	L Logger
}

// Main func for Hand, Inly Greets
func Main(a Aid) {
	fmt.Print(welcomeMessage)
	a.O.Red("\"God\"")
}

// Call will return in God's Hand
func Call() Aid {
	return Aid{
		O: NewOutput(), I: NewInput(), U: NewUtils(), S: NewSystem(), L: NewLogger()}
}

func LogInfo(a ...interface{}) {
	l := NewLogger()
	l.SetSkip(2).Info(a...)
}
