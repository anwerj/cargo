package primefactors

import (
	"fmt"

	"github.com/anwerj/cargo/head"
)

const lineBreaker = "\n*********************************************************\n"

// Main of Prime Factors
func Main(a head.Aid) {

	c := a.I.ExpectStr(2, a.I.MultiChoice([]string{"recursive", "iterative"}, "Enter the Mode"))
	t := a.I.ExpectInt(3, a.I.Int("Enter the integer"))

	fmt.Println("\nWe are now starting Prime Factorization in Mode", c, "for", t)

	var listy []int // Will always be empty as it is passed empty
	var frist []int
	var prist []int
	misty := head.NewSmap()
	misty.SetPad("%4s => ")

	for i := 2; i <= t; i++ {
		switch c {
		case "iterative":
			frist = iterative(i, listy)
		case "recursive":
			frist = recursive(i, listy)
		}
		if len(frist) == 1 {
			// i is a prime number
			prist = append(prist, i)
		} else {
			misty.Int(i, frist)
		}
	}

	a.O.Boxes(map[string]interface{}{
		"Count of Factors":  misty, //h.U.SortByKeys(misty, h.O.SprintKV, "%4s => "),
		"All prime numbers": prist,
	})
}

// recursive way of finding primes
func recursive(i int, listy []int) []int {

	for j := 2; j < (i/2 + 1); j++ {
		if i%j == 0 {
			listy = append(listy, j)
			return recursive(i/j, listy)
		}
	}
	listy = append(listy, i)
	return listy
}

// iterative way of finding primes
func iterative(n int, listy []int) []int {

	for n%2 == 0 {
		listy = append(listy, 2)
		n = n / 2
	}

	for i := 3; i*i < n; i += 2 {
		for n%i == 0 {
			listy = append(listy, i)
			n = n / i
		}
	}

	return listy
}
