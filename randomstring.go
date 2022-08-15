package randomstring

import (
	"math/rand"
	"time"
)

func randIndex(s string) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(len(s))
}

func GetRandomCharacter(s string) string {
	randomIndex := randIndex(s)
	randomChar := string(s[randomIndex])
	return randomChar
}

func GenRandStr(s string, lenStr int) string {

	if s == "" {
		return ""
	}

	outputStr := ""

	for i := 0; i < lenStr; i++ {
		outputStr = outputStr + GetRandomCharacter(s)
	}

	return outputStr
}
