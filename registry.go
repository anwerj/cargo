package main

import (
	"my/god"
	"my/hocon"
	"my/initial"
	"my/install"
	"my/jsonnet"
	"my/prime_factors"
	"my/strings"
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
