package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"github.com/carht/randomstring"
)

func version() {
	version := "0.1.2"
	fmt.Println(version)
}

func usage() {
	usage := `randstr [Argument]
                 
        -h or help    : Return this help.
        -v or version : Return the version. 
        -n            : Return numbers.
        -l            : Return lowercase letters.
        -u            : Return uppercase letters.
        -nl or -ln    : Return lowercase letters with numbers.
        -un or -nu    : Return uppercase letters with numbers.
        -ul or -lu    : Return uppercase letters with lowercase letters.
        default       : Return a string of 32 characters between uppercases, lowercases and numbers.

        Note: All previous parameters can use a number for the output length.

        Some examples:
        
        $ randstr
        cDUKhgpkuIgMdqs3JibEjJoyBiBjgjZi
        
        $ randstr 35
        ghrzAtgElMIMxCRFMFCbUCnwiDuEkejKocQ

        $ randstr -n
        54770780734477138272862081928945
     
        $ randstr -n 33
        280309865114201905814323059174146

        $ randstr -nl 35
        hs9wsuhr6m2dyez2zdfyxduglsccygtyzaq
`

	fmt.Println(usage)
}

func selectInput(s string) {
	if len(os.Args) == 2 {
		randStr := randomstring.GenRandStr(s, 32)
		fmt.Println(randStr)
	} else {
		arg2 := os.Args[2]
		inputLenStr, err := strconv.Atoi(arg2)
		if err != nil {
			log.Fatal(err)
		}
		randStr := randomstring.GenRandStr(s, inputLenStr)
		fmt.Println(randStr)
	}
}


func main() {
	uABC := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	labc := "abcdefghijklmnopqrstuvwxyz"
	nums := "0123456789"
	

	if len(os.Args) > 1 {
		arg := os.Args[1]
	
		switch arg {
		case "version":
			version()
		case "-v":
			version()
		case "help":
			usage()
		case "-h":
			usage()
		case "-n":
			selectInput(nums)
		case "-l":
			selectInput(labc)
		case "-nl", "-ln":
			selectInput(nums+labc)
		case "-u":
			selectInput(uABC)
		case "-nu", "-un":
			selectInput(nums+uABC)
		case "-lu", "-ul":
			selectInput(labc+uABC)
		default:
			inputLenStr, err := strconv.Atoi(arg)
			if err != nil {
				log.Fatal(err)
			}

			randStr := randomstring.GenRandStr(labc+uABC, inputLenStr)
			fmt.Println(randStr)
		}
	} else {
		randStr := randomstring.GenRandStr(nums+labc+uABC, 32)
		fmt.Println(randStr)
	}
}
