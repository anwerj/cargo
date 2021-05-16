package main

import (
	"github.com/anwerj/MyGod/god"
	"github.com/anwerj/MyGod/hocon"
	"github.com/anwerj/MyGod/initial"
	"github.com/anwerj/MyGod/install"
	"github.com/anwerj/MyGod/jsonnet"
	"github.com/anwerj/MyGod/prime_factors"
	"github.com/anwerj/MyGod/strings"
)

func registry() map[string] func(h god.Hand) {
	flist := map[string]func(h god.Hand) {
		"god" : god.Main,
		"hocon" : hocon.Main,
		"initial" : initial.Main,
		"install" : install.Main,
		"jsonnet" : jsonnet.Main,
		"prime_factors" : prime_factors.Main,
		"strings" : strings.Main,
	}

	return flist
}
