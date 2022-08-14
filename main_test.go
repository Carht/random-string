package main

import "testing"

func TestGetARandomIndexFromAString(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	for i := 0; i < 256; i++ {
		result := randIndex(input)

		if result > len(input) || result < 0 {
			t.Errorf("Error, The result is output of the posible values, the result was %d", result)
		}
	}
}

func TestGetARandomElementFromAString(t *testing.T) {
	input := "abcdefghijklmnopqrstuvwxyz"
	result := getRandomCharacter(input)

	for range input {
		char := isMember(result, input)

		if char != true {
			t.Errorf("The expected value is %t, but the answer was %t", true, char)
		}
	}
}

