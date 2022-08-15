package randomstring

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

	for range input {
		result := GetRandomCharacter(input)
		char := isMember(result, input)

		if char != true {
			t.Errorf("The expected value is %t, but the answer was %t", true, char)
		}
	}
}

func TestGenerateRandomStringOfSomeLength(t *testing.T) {
	abcLower := "abcdefghijklmnopqrstuvwxyz"

	for i := range abcLower {
		result := GenRandStr(abcLower, i)

		if len(result) != i {
			t.Errorf("The expected value is %d, but the answer was %d", i, len(result))
		}
	}
}

func TestVerifyIfTheInputLengthIsZero(t *testing.T) {
	hi := "Hello Golang"
	outputLenStr := 0

	result := GenRandStr(hi, outputLenStr)
	expected := ""

	if result != expected {
		t.Errorf("The expected value is %s, but the answer was %s", expected, result)
	}
}

func TestVerifyIfTheInputLenIsNegative(t *testing.T) {
	hi := "hello Golang"
	outputLenStr := []int{-1, -2, -3, -5, -10, -50, -100, -1000, -1234, -9999}

	for i := range outputLenStr {
		result := GenRandStr(hi, outputLenStr[i])
		expected := ""

		if result != expected {
			t.Errorf("The expected value is %s, but the answer was %s", expected, result)
		}
	}
}

func TestVerifyEmptyStringAndZeroAndNegativeLenOutputStr(t *testing.T) {
	hi := ""
	outputLenStr := []int{0, -1, -2, -3, -4, -5, -7, -11, -20, -50, -32, -64, -256}

	for i := range outputLenStr {
		result := GenRandStr(hi, outputLenStr[i])
		expected := ""

		if result != expected {
			t.Errorf("The expected value is %s, but the answer was %s", expected, result)
		}
	}
}

func TestVerifyEmptyStringAndPositiveLenOutputStr(t *testing.T) {
	hi := ""
	outputLenStr := []int{1, 2, 3, 4, 5, 7, 11, 20, 50, 32, 64, 256, 9999}

	for i := range outputLenStr {
		result := GenRandStr(hi, outputLenStr[i])
		expected := ""

		if result != expected {
			t.Errorf("The expected value is %s, but the answer was %s", expected, result)
		}
	}
}
