package hocon

import (
	"encoding/json"
	"fmt"

	"github.com/anwerj/cargo/head"
)

type config struct {
	Steps []interface{} `json:"steps"`
}

// Main will ask for file or option
func Main(a head.Aid) {
	o := a.I.ExpectStr(a.I.A(), a.I.Str("Enter the file"))
	if o == "install" {
		Install(a)
		return
	}
	Generate(a, o)
}

// Install Registry
func Install(a head.Aid) {
	d := a.I.ExpectStr(a.I.A(), a.I.MultiChoice([]string{
		"hocon/configs"},
		"Choose config directory"))
	dirs := a.S.ReadFiles(d, func(s string) bool { return s != "base.properties" })

	for _, p := range dirs {
		if "Yes" == a.I.MultiChoice([]string{"Yes", "No"}, p)() {
			Generate(a, p)
		}
	}
}

// Generate on Path
func Generate(a head.Aid, p string) {
	var conf config

	f := a.S.AbsPath(p)

	if a.S.PathExists(f) == false {
		a.O.Error("Path does not exists", f)
		return
	}
	c := a.S.Exec("pyhocon", "-f", "json", "-i", f)

	err := json.Unmarshal([]byte(c), &conf)
	if err != nil {
		a.O.Error("Could not unmarshal", err, c)
	}

	b, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		a.O.Error("Could not marshal", err, c)
	}

	fmt.Printf("%s\n", b)
}
