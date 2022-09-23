package main

import (
	"fmt"
	"flag"
	"github.com/carht/randomstring"
)

func version() {
	version := "0.1.3"
	fmt.Println(version)
}

func main() {
	uABC := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	labc := "abcdefghijklmnopqrstuvwxyz"
	nums := "0123456789"

	retNumbers := flag.Int("n", 0, "Return a string of numbers.")
	retLowLetters := flag.Int("l", 0, "Return a string of lowercase letters.")
	retUpLetters := flag.Int("u", 0, "Return a string of uppercase letters.")
	retLowAndNums := flag.Int("nl", 0, "Return a string of lowercase letters with numbers.")
	retUpAndNums := flag.Int("nu", 0, "Return a string of uppercase letters with numbers.")
	retUpAndLow := flag.Int("lu", 0, "Return a string of uppercase with lowercase letters.")
	versn := flag.Bool("v", false, "Return the version.")

	flag.Parse()

	switch {
	case *versn:
		version()
	case *retNumbers != 0:
		fmt.Println(randomstring.GenRandStr(nums, *retNumbers))
	case *retLowLetters != 0:
		fmt.Println(randomstring.GenRandStr(labc, *retLowLetters))
	case *retUpLetters != 0:
		fmt.Println(randomstring.GenRandStr(uABC, *retUpLetters))
	case *retLowAndNums != 0:
		fmt.Println(randomstring.GenRandStr(nums+labc, *retLowAndNums))
	case *retUpAndNums != 0:
		fmt.Println(randomstring.GenRandStr(nums+uABC, *retUpAndNums))
	case *retUpAndLow != 0:
		fmt.Println(randomstring.GenRandStr(labc+uABC, *retUpAndLow))
	default:
		fmt.Println(randomstring.GenRandStr(labc+uABC, 32))
	}
}
