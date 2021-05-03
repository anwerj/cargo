package hocon

import (
	"encoding/json"
	"fmt"

	"github.com/anwerj/MyGod/god"
)

type config struct {
	Steps []interface{} `json:"steps"`
}

// Main will ask for file or option
func Main(h god.Hand) {
	o := h.I.ExpectStr(h.I.A(), h.I.Str("Enter the file"))
	if o == "install" {
		Install(h)
		return
	}
	Generate(h, o)
}

// Install Registry
func Install(h god.Hand) {
	d := h.I.ExpectStr(h.I.A(), h.I.MultiChoice([]string{
		"hocon/configs"},
		"Choose config directory"))
	dirs := h.S.ReadFiles(d, func(s string) bool { return s != "base.properties" })

	for _, p := range dirs {
		if "Yes" == h.I.MultiChoice([]string{"Yes", "No"}, p)() {
			Generate(h, p)
		}
	}
}

// Generate on Path
func Generate(h god.Hand, p string) {
	var conf config

	f := h.S.AbsPath(p)

	if h.S.PathExists(f) == false {
		h.O.Error("Path does not exists", f)
		return
	}
	c := h.S.Exec("pyhocon", "-f", "json", "-i", f)

	err := json.Unmarshal([]byte(c), &conf)
	if err != nil {
		h.O.Error("Could not unmarshal", err, c)
	}

	b, err := json.MarshalIndent(conf, "", "  ")
	if err != nil {
		h.O.Error("Could not marshal", err, c)
	}

	fmt.Printf("%s\n", b)
}
