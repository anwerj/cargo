package main

import (
	"github.com/anwerj/cargo/head"
	"github.com/anwerj/cargo/hocon"
	"github.com/anwerj/cargo/initial"
	"github.com/anwerj/cargo/install"
	"github.com/anwerj/cargo/jsonnet"
	"github.com/anwerj/cargo/primefactors"
	"github.com/anwerj/cargo/regex"
	"github.com/anwerj/cargo/strings"
)

func registry() map[string] func(a head.Aid) {
	flist := map[string]func(a head.Aid) {
		"head" : head.Main,
		"hocon" : hocon.Main,
		"initial" : initial.Main,
		"install" : install.Main,
		"jsonnet" : jsonnet.Main,
		"primefactors" : primefactors.Main,
		"regex" : regex.Main,
		"strings" : strings.Main,
	}

	return flist
}
