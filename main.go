package main

import (
	"time"
	"math/rand"
)

func randIndex(s string) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(s))
}

func getRandomCharacter(s string) string {
	randomIndex := randIndex(s)
	randomChar := string(s[randomIndex])
	return randomChar
}

func genRandStr(s string, lenStr int) string {
	outputStr := ""
	
	for i := 0; i < lenStr; i++ {
		outputStr = outputStr + getRandomCharacter(s)
	}

	return outputStr
}
