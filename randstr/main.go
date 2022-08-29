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
	usage := `randomstring [Argument]
                  
                  randomstring:
                    Without arguments return a string of 32 characters between uppercases, lowercases and numbers.
                  
                  randomstring version | randomstring -v:
                    Return the version of this example.

                  randomstring help | randomstring -h:
                    Return this help.`

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
