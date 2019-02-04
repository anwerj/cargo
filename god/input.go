package god

import (
	"fmt"
	"os"
	"strconv"

	"github.com/abiosoft/ishell"
)

// Input is collection of input dependencies
type Input struct {
	i *ishell.Shell
	a int
}

// NewInput creates a fresh Input with Dependencies
func NewInput() Input {
	return Input{i: ishell.New(), a: 1}
}

// A will increase and return Arg position
func (I *Input) A() int {
	I.a++
	return I.a
}

// Args will read cmd
func (I *Input) Args(i int) string {
	if len(os.Args) > i {
		return os.Args[i]
	}

	return ""
}

// MultiChoice will return as function to str
func (I *Input) MultiChoice(cs []string, q string) (
	out func() (s string)) {

	return func() (s string) {
		s = cs[I.i.MultiChoice(cs, q)]
		return
	}
}

// Int will ask to enter integer
func (I *Input) Int(m string) (
	out func() (i int)) {

	return func() (i int) {

		d := ""
		fmt.Println(m)
		fmt.Scanf("%s", &d)

		if o, e := strconv.Atoi(d); e != nil {
			return I.Int(e.Error())()
		} else {
			i = o
		}
		return
	}
}

// Str will ask to enter the string
func (I *Input) Str(m string) (
	out func() (s string)) {

	return func() (s string) {
		fmt.Println(m)
		if _, e := fmt.Scan(&s); e != nil {
			return I.Str(e.Error())()
		}
		return
	}
}

// ExpectInt will first check int in arg, then fire fallback
func (I *Input) ExpectInt(p int, c func() int) (out int) {
	s := I.Args(p)

	if s != "" {
		o, e := strconv.Atoi(s)
		if e != nil {
			out = c()
		} else {
			out = o
		}
	} else {
		out = c()
	}

	return
}

// ExpectStr will first check str in arg, then fire fallback
func (I *Input) ExpectStr(p int, c func() string) (out string) {
	s := I.Args(p)

	if s != "" {
		out = s
	} else {
		out = c()
	}

	return
}
