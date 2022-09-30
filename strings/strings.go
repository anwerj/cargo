package strings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anwerj/cargo/head"
)

var c2f = map[string]func(a head.Aid, s string){
	"len":   strlen,
	"slice": strslice,
}

func Main(a head.Aid) {
	a.O.Green("Welcome Stranger!")
	c := a.I.ExpectStr(a.I.A(), a.I.Str("Please enter the command/string"))

	f, ok := c2f[c]
	if ok {
		f = def
	} else {
		c = a.I.ExpectStr(a.I.A(), a.I.Str("Please enter the string"))
	}
	f(a, c)
}

func def(a head.Aid, s string) {
	fmt.Println("Handling with default function: " + s)
	l := len(s)

	fmt.Printf("[%d:%d] = %s\n", -1, l, s[:l])
	fmt.Printf("[%d:%d] = %s\n", l-1, l, s[l-1:l])
	fmt.Printf("[%d:%d] = %s\n", 1, l, s[1:l])
	fmt.Printf("[%d:%d] = %s\n", 1, l-1, s[1:l-1])
}

func strlen(a head.Aid, s string) {
	a.O.Green("Length of the string '%s' is %d.", s, len(s))
}

func strslice(a head.Aid, s string) {
	a.O.Red("Command is under construction, thanks for your patience")
	var ps string
	var pl, pr int
	var err error
	for {
		ps = a.I.Str("Enter the left,right points: 0, " + strconv.Itoa(len(s)))()
		pl, pr = 0, len(s)
		pps := strings.Split(ps, ",")
		if len(pps) > 0 {
			pl, err = strconv.Atoi(pps[0])
			if err != nil {
				pl = 0
			}
		}
		if len(pps) > 1 {
			pr, err = strconv.Atoi(pps[1])
			if err != nil {
				pr = len(s)
			}
		}

		a.O.Green("Sliced string for [%d:%d] is '%s'", pl, pr, s[pl:pr])
		fmt.Println()
	}
}
