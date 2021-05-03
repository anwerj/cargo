package strings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/anwerj/MyGod/god"
)

var c2f = map[string]func(h god.Hand, s string){
	"len":   strlen,
	"slice": strslice,
}

func Main(h god.Hand) {
	h.O.Green("Welcome Stranger!")
	c := h.I.ExpectStr(h.I.A(), h.I.Str("Please enter the command/string"))

	f, ok := c2f[c]
	if ok == false {
		f = def
	} else {
		c = h.I.ExpectStr(h.I.A(), h.I.Str("Please enter the string"))
	}
	f(h, c)
}

func def(h god.Hand, s string) {
	fmt.Println("Handling with default function: " + s)
	l := len(s)

	fmt.Printf("[%d:%d] = %s\n", -1, l, s[:l])
	fmt.Printf("[%d:%d] = %s\n", l-1, l, s[l-1:l])
	fmt.Printf("[%d:%d] = %s\n", 1, l, s[1:l])
	fmt.Printf("[%d:%d] = %s\n", 1, l-1, s[1:l-1])
}

func strlen(h god.Hand, s string) {
	h.O.Green("Length of the string '%s' is %d.", s, len(s))
}

func strslice(h god.Hand, s string) {
	h.O.Red("Command is under construction, thanks for your patience")
	var ps string
	var pl, pr int
	var err error
	for true {
		ps = h.I.Str("Enter the left,right points: 0, " + strconv.Itoa(len(s)))()
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

		h.O.Green("Sliced string for [%d:%d] is '%s'", pl, pr, s[pl:pr])
		fmt.Println()
	}
}
