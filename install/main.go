package install

import (
	"fmt"
	"io/ioutil"

	"github.com/anwerj/MyGod/god"
)

// Main of install
func Main(h god.Hand) {
	dirs := h.S.ReadDir(".")

	con := fileContent(dirs)

	err := ioutil.WriteFile("registry.go", con, 0644)

	if err == nil {
		fmt.Println("Installation completed with ", dirs)
	} else {
		fmt.Println(err, con)
	}
}

// fileContent gives file content
func fileContent(dirs []string) []byte {
	text := "package main\n" +
		"\n" +
		"import (\n"

	for _, dir := range dirs {
		text += "\t\"github.com/anwerj/MyGod/" + dir + "\"\n"
	}

	text += ")\n\n" +
		"func registry() map[string] func(h god.Hand) {\n" +
		"\tflist := map[string]func(h god.Hand) {\n"

	for _, dir := range dirs {
		text += "\t\t\"" + dir + "\" : " + dir + ".Main,\n"
	}

	text += "\t}\n\n" +
		"\treturn flist\n" +
		"}\n"

	return []byte(text)
}
