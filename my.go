package main

import (
	"my/god"
	"os"
)

func main() {
	flist, h := registry(), god.Call()

	ch := ""
	if len(os.Args) > 1 {
		ch = os.Args[1]
		if ch[:1] == "/" {

		}
	}

	fn, ok := flist[ch]
	if ok {
		fn(h)
	} else {
		var cmds []string
		for fn := range flist {
			cmds = append(cmds, fn)
		}
		f := h.I.MultiChoice(cmds, "Where you want to go")()

		flist[f](h)
	}
}
