package head

import (
	"fmt"
	"strings"

	"github.com/fatih/color"
)

const lineSeparator = "\n*********************************************************\n"

// Output collection of dependencies
type Output struct {
	color *color.Color
	lb    bool
}

// NewOutput will return the OutPut
func NewOutput() Output {
	return Output{color: color.New(), lb: true}
}

// Error prints error output
func (O *Output) Error(format string, a ...interface{}) {
	f, d := color.RedString(format), fmt.Sprint(a...)
	fmt.Println(f, d)
}

// Green prints  green output
func (O *Output) Green(format string, a ...interface{}) {
	if O.lb {
		color.Green(format, a...)
		return
	}
	fmt.Print(color.GreenString(format))
}

// Red prints red output
func (O *Output) Red(format string, a ...interface{}) {
	if O.lb {
		color.Red(format, a...)
		return
	}
	fmt.Print(color.RedString(format))
}

// WithLB will ask output to line break
func (O *Output) WithLB() *Output {
	O.lb = true
	return O
}

// WoutLB will ask output to ignore line break
func (O *Output) WoutLB() *Output {
	O.lb = false
	return O
}

// Boxes will print map in boxes
func (O *Output) Boxes(s map[string]interface{}) {
	for k, v := range s {
		l := (80 - len(k)) / 2
		if l%2 == 0 {
			l--
		}
		format := strings.Repeat("*", l)
		O.Green("\n"+format+" %s "+format, k)
		fmt.Println(v)
	}
}

// SprintKV will simply print the key-value pair
func (O *Output) SprintKV(
	k string,
	v interface{}) (o string) {

	o += fmt.Sprint(k) + fmt.Sprintln(v)

	return o
}
