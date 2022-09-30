package main

import (
	"os"

	"github.com/anwerj/cargo/head"
)

func main() {
	flist, a := registry(), head.Call()

	ch := ""
	if len(os.Args) > 1 {
		ch = os.Args[1]
		if ch[:1] == "/" {
			//ch = ch[]
		}
	}

	fn, ok := flist[ch]
	if ok {
		fn(a)
	} else {
		var cmds []string
		for fn := range flist {
			cmds = append(cmds, fn)
		}
		f := a.I.MultiChoice(cmds, "Where you want to go")()

		flist[f](a)
	}
}
