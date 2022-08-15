package main

import (
	"os"
	"fmt"
	"strconv"
	"github.com/carht/randomstring"
)

func main() {
	abc := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	if len(os.Args) > 1 {
		outStrLen, _ := strconv.Atoi(os.Args[1])
		
		outStr := randomstring.GenRandStr(abc, outStrLen)
		fmt.Println(outStr)
		
	} else {
		outStr := randomstring.GenRandStr(abc, 32)
		fmt.Println(outStr)
	}	
}
