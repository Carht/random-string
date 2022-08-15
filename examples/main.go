package main

import (
	"os"
	"fmt"
	"log"
	"strconv"
	"github.com/carht/randomstring"
)

func main() {
	abc := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if len(os.Args) > 1 {
		outStrLen, err := strconv.Atoi(os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
		
		outStr := randomstring.GenRandStr(abc, outStrLen)
		fmt.Println(outStr)
		
	} else {
		outStr := randomstring.GenRandStr(abc, 32)
		fmt.Println(outStr)
	}	
}
