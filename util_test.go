package randomstring

import "testing"

func TestIfSomeCharacterIsMemberOfAString(t *testing.T) {
	result := isMember("h", "some test hi!")
	expected := true

	if expected != true {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}

func TestIfSomeCharacterNotExistInAString(t *testing.T) {
	result := isMember("h", "an example of testing.")
	expected := false

	if expected != false {
		t.Errorf("The expected value is %t, but the answer was %t", expected, result)
	}
}
