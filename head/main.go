package head

import (
	"fmt"

	"github.com/anwerj/cargo/head/input"
	"github.com/anwerj/cargo/head/log"
	"github.com/anwerj/cargo/head/output"
	"github.com/anwerj/cargo/head/system"
	"github.com/anwerj/cargo/head/utils"
)

// Aid holds everything needed in project, for all dependencies just talk to hand
type Aid struct {
	O output.Output
	I input.Input
	U utils.Utils
	S system.System
	L log.Log
}

// Main func for Hand, Inly Greets
func Main(a Aid) {
	fmt.Print("Yes my friend! You have reached to the ")
	a.O.WoutLB().Blue("\"Head\"")
	fmt.Print(" of the cargo.\n")
}

// Call will return in God's Hand
func Call() Aid {
	return Aid{
		O: output.New(), I: input.New(), U: utils.New(), S: system.New(), L: log.New()}
}

func LogInfo(a ...interface{}) {
	l := log.New()
	l.SetSkip(2).Info(a...)
}
